package main

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Peersyst/xrpl-go/examples/helpers"
	"github.com/Peersyst/xrpl-go/pkg/crypto"
	rippleTime "github.com/Peersyst/xrpl-go/xrpl/time"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
)

func main() {
	// As of February 2025, Credential is only available on Devnet.
	client := helpers.GetDevnetRpcClient()

	// Configure wallets

	// Issuer
	fmt.Println("⏳ Setting up credential issuer wallet...")
	issuer, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Printf("❌ Error creating issuer wallet: %s\n", err)
		return
	}

	err = client.FundWallet(&issuer)
	if err != nil {
		fmt.Printf("❌ Error funding issuer wallet: %s\n", err)
		return
	}
	fmt.Printf("✅ Issuer wallet funded: %s\n", issuer.ClassicAddress)

	// -----------------------------------------------------

	// Subject (destination)
	fmt.Println("⏳ Setting up Subject wallet...")
	subjectWallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Printf("❌ Error creating subject wallet: %s\n", err)
		return
	}

	err = client.FundWallet(&subjectWallet)
	if err != nil {
		fmt.Printf("❌ Error funding subject wallet: %s\n", err)
		return
	}
	fmt.Printf("✅ Subject wallet funded: %s\n", subjectWallet.ClassicAddress)

	// -----------------------------------------------------

	// Creating the CredentialCreate transaction
	fmt.Println("⏳ Creating CredentialCreate transaction...")

	expiration, err := rippleTime.IsoTimeToRippleTime(time.Now().Add(time.Hour * 24).Format(time.RFC3339))
	credentialType := types.CredentialType("6D795F63726564656E7469616C")

	if err != nil {
		fmt.Printf("❌ Error converting expiration to ripple time: %s\n", err)
		return
	}

	txn := &transaction.CredentialCreate{
		BaseTx: transaction.BaseTx{
			Account: types.Address(issuer.ClassicAddress),
		},
		CredentialType: credentialType,
		Subject:        types.Address(subjectWallet.ClassicAddress),
		Expiration:     uint32(expiration),
		URI:            hex.EncodeToString([]byte("https://example.com")),
	}

	helpers.SubmitAndWait(client, txn, issuer)

	// -----------------------------------------------------

	// Creating the CredentialAccept transaction
	fmt.Println("⏳ Creating CredentialAccept transaction...")

	acceptTxn := &transaction.CredentialAccept{
		BaseTx: transaction.BaseTx{
			Account: types.Address(subjectWallet.ClassicAddress),
		},
		CredentialType: credentialType,
		Issuer:         types.Address(issuer.ClassicAddress),
	}

	helpers.SubmitAndWait(client, acceptTxn, subjectWallet)

	// -----------------------------------------------------

	// Creating the CredentialDelete transaction
	fmt.Println("⏳ Creating CredentialDelete transaction...")

	deleteTxn := &transaction.CredentialDelete{
		BaseTx: transaction.BaseTx{
			Account: types.Address(issuer.ClassicAddress),
		},
		CredentialType: credentialType,
		Issuer:         types.Address(issuer.ClassicAddress),
		Subject:        types.Address(subjectWallet.ClassicAddress),
	}

	helpers.SubmitAndWait(client, deleteTxn, issuer)
}
