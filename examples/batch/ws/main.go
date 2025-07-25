package main

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/crypto"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/transaction"
	txnTypes "github.com/Peersyst/xrpl-go/xrpl/transaction/types"
	"github.com/Peersyst/xrpl-go/xrpl/wallet"
	"github.com/Peersyst/xrpl-go/xrpl/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/websocket/types"
)

var (
	CreatePaymentTx = func(sender, receiver *wallet.Wallet, amount txnTypes.CurrencyAmount) *transaction.Payment {
		return &transaction.Payment{
			BaseTx: transaction.BaseTx{
				Account:         sender.GetAddress(),
				TransactionType: transaction.PaymentTx,
				Flags:           txnTypes.TfInnerBatchTxn,
			},
			Amount:      amount,
			Destination: receiver.GetAddress(),
		}
	}
)

func main() {
	// Connect to testnet
	client := websocket.NewClient(
		websocket.NewClientConfig().
			WithHost("wss://s.devnet.rippletest.net:51233").
			WithFaucetProvider(faucet.NewDevnetFaucetProvider()),
	)
	defer client.Disconnect()

	if err := client.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	// Create and fund wallets
	userWallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	user2Wallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}
	receiverWallet, err := wallet.New(crypto.ED25519())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("⏳ Funding wallets...")
	if err := client.FundWallet(&userWallet); err != nil {
		fmt.Println(err)
		return
	}
	if err := client.FundWallet(&user2Wallet); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("💸 Wallets funded")

	// Check initial balances
	userBalance, err := client.GetXrpBalance(userWallet.ClassicAddress)
	if err != nil {
		userBalance = "0"
	}
	user2Balance, err := client.GetXrpBalance(user2Wallet.ClassicAddress)
	if err != nil {
		user2Balance = "0"
	}

	receiverBalance, err := client.GetXrpBalance(receiverWallet.ClassicAddress)
	if err != nil {
		receiverBalance = "0"
	}

	fmt.Printf("💳 User initial balance: %s XRP\n", userBalance)
	fmt.Printf("💳 User2 initial balance: %s XRP\n", user2Balance)
	fmt.Printf("💳 Receiver initial balance: %s XRP\n", receiverBalance)
	fmt.Println()

	fmt.Printf("Batch transaction test\n")

	// Create test batch transaction
	batchTx := &transaction.Batch{
		BaseTx: transaction.BaseTx{
			Account:         txnTypes.Address(userWallet.ClassicAddress),
			TransactionType: transaction.BatchTx,
		},
		RawTransactions: []txnTypes.RawTransaction{
			{RawTransaction: CreatePaymentTx(&userWallet, &receiverWallet, txnTypes.XRPCurrencyAmount(5000000)).Flatten()},
			{RawTransaction: CreatePaymentTx(&userWallet, &receiverWallet, txnTypes.XRPCurrencyAmount(5000000)).Flatten()},
		},
	}
	batchTx.SetAllOrNothingFlag()

	flattenedBatchTx := batchTx.Flatten()
	fmt.Println("⏳ Autofilling flattened batch transaction...")
	if err := client.Autofill(&flattenedBatchTx); err != nil {
		fmt.Println("Autofill error:", err)
		return
	}

	fmt.Println("⏳ Signing batch transaction...")
	response, err := client.SubmitTxAndWait(flattenedBatchTx, &types.SubmitOptions{
		Autofill: false,
		Wallet:   &userWallet,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("✅ Batch transaction submitted")
	fmt.Printf("🌐 Hash: %s\n", response.Hash.String())
	fmt.Printf("🌐 Validated: %t\n", response.Validated)
	fmt.Println()

	// Check final balances
	finalUserBalance, err := client.GetXrpBalance(userWallet.ClassicAddress)
	if err != nil {
		finalUserBalance = "0"
	}
	finalReceiverBalance, err := client.GetXrpBalance(receiverWallet.ClassicAddress)
	if err != nil {
		finalReceiverBalance = "0"
	}

	fmt.Printf("💳 User final balance: %s XRP\n", finalUserBalance)

	fmt.Printf("💳 Receiver final balance: %s XRP\n", finalReceiverBalance)

	fmt.Println()
	fmt.Printf("Multisig Batch transaction test\n")

	// Create test batch transaction
	multiBatchTx := &transaction.Batch{
		BaseTx: transaction.BaseTx{
			Account:         txnTypes.Address(userWallet.ClassicAddress),
			TransactionType: transaction.BatchTx,
		},
		RawTransactions: []txnTypes.RawTransaction{
			{RawTransaction: CreatePaymentTx(&userWallet, &receiverWallet, txnTypes.XRPCurrencyAmount(5000000)).Flatten()},
			{RawTransaction: CreatePaymentTx(&user2Wallet, &receiverWallet, txnTypes.XRPCurrencyAmount(5000000)).Flatten()},
		},
		BatchSigners: []txnTypes.BatchSigner{
			{
				BatchSigner: txnTypes.BatchSignerData{
					Account:       txnTypes.Address(user2Wallet.ClassicAddress),
					SigningPubKey: user2Wallet.PublicKey,
				},
			},
		},
	}
	multiBatchTx.SetAllOrNothingFlag()

	flattenedMultiBatchTx := multiBatchTx.Flatten()
	fmt.Println("⏳ Autofilling flattened multi batch transaction...")
	if err := client.AutofillMultisigned(&flattenedMultiBatchTx, 1); err != nil {
		fmt.Println("Autofill error:", err)
		return
	}

	fmt.Println("⏳ Signing multi batch transaction...")
	if err := wallet.SignMultiBatch(user2Wallet, &flattenedMultiBatchTx, nil); err != nil {
		fmt.Println("Signing error:", err)
		return
	}

	response, err = client.SubmitTxAndWait(flattenedMultiBatchTx, &types.SubmitOptions{
		Autofill: false,
		Wallet:   &userWallet,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("✅ Multisig Batch transaction submitted")
	fmt.Printf("🌐 Hash: %s\n", response.Hash.String())
	fmt.Printf("🌐 Validated: %t\n", response.Validated)
	fmt.Println()

	// Check final balances
	finalUser2Balance, err := client.GetXrpBalance(user2Wallet.ClassicAddress)
	if err != nil {
		finalUser2Balance = "0"
	}
	finalUserBalance, err = client.GetXrpBalance(userWallet.ClassicAddress)
	if err != nil {
		finalUserBalance = "0"
	}
	finalReceiverBalance, err = client.GetXrpBalance(receiverWallet.ClassicAddress)
	if err != nil {
		finalReceiverBalance = "0"
	}
	fmt.Printf("💳 User final balance: %s XRP\n", finalUserBalance)
	fmt.Printf("💳 User2 final balance: %s XRP\n", finalUser2Balance)
	fmt.Printf("💳 Receiver final balance: %s XRP\n", finalReceiverBalance)
}
