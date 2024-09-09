package main

import (
	"fmt"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/Peersyst/xrpl-go/xrpl"
	public_servers "github.com/Peersyst/xrpl-go/xrpl/client/public_servers"
	"github.com/Peersyst/xrpl-go/xrpl/client/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
)

const (
	currencyCode = "FOO"
)

func main() {
	//
	// Configure client
	//
	fmt.Println("⏳ Setting up client...")

	// init public urls for websocket client
	publicServers := public_servers.NewServerUrls()

	client := websocket.NewWebsocketClient(
		websocket.NewWebsocketClientConfig().
			WithHost(publicServers.TestnetWebSocket().Ripple()).
			WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)
	fmt.Println("✅ Client configured!")
	fmt.Println()

	//
	// Configure wallets
	//
	fmt.Println("⏳ Setting up wallets...")
	coldWallet, err := xrpl.NewWallet(addresscodec.ED25519)
	if err != nil {
		fmt.Printf("❌ Error creating cold wallet: %s\n", err)
		return
	}
	err = client.FundWallet(&coldWallet)
	if err != nil {
		fmt.Printf("❌ Error funding cold wallet: %s\n", err)
		return
	}
	fmt.Println("💸 Cold wallet funded!")

	hotWallet, err := xrpl.NewWallet(addresscodec.ED25519)
	if err != nil {
		fmt.Printf("❌ Error creating hot wallet: %s\n", err)
		return
	}
	err = client.FundWallet(&hotWallet)
	if err != nil {
		fmt.Printf("❌ Error funding hot wallet: %s\n", err)
		return
	}
	fmt.Println("💸 Hot wallet funded!")
	fmt.Println()

	fmt.Println("✅ Wallets setup complete!")
	fmt.Println("💳 Cold wallet:", coldWallet.ClassicAddress)
	fmt.Println("💳 Hot wallet:", hotWallet.ClassicAddress)
	fmt.Println()

	//
	// Configure cold address settings
	//
	fmt.Println("⏳ Configuring cold address settings...")
	coldWalletAccountSet := &transactions.AccountSet{
		BaseTx: transactions.BaseTx{
			Account: types.Address(coldWallet.ClassicAddress),
		},
		TickSize:     5,
		TransferRate: 0,
		Domain:       "6578616D706C652E636F6D", // example.com
	}

	coldWalletAccountSet.SetAsfAllowTrustLineClawback()
	coldWalletAccountSet.SetDisallowXRP()

	coldWalletAccountSet.SetRequireDestTag()

	flattenedTx := coldWalletAccountSet.Flatten()

	err = client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error autofilling transaction: %s\n", err)
		return
	}

	txBlob, _, err := coldWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error signing transaction: %s\n", err)
		return
	}

	response, err := client.SubmitTransactionBlob(txBlob, false)
	if err != nil {
		fmt.Printf("❌ Error submitting transaction: %s\n", err)
		return
	}

	if response.EngineResult != "tesSUCCESS" {
		fmt.Println("❌ Cold wallet unfreezing failed!", response.EngineResult)
		fmt.Println("Try again!")
		fmt.Println()
		return
	}

	fmt.Println("✅ Cold address settings configured!")
	fmt.Printf("🌐 Hash: %s\n", response.Tx["hash"])
	fmt.Println()

	//
	// Create trust line from hot to cold address
	//
	fmt.Println("⏳ Creating trust line from hot to cold address...")
	hotColdTrustSet := &transactions.TrustSet{
		BaseTx: transactions.BaseTx{
			Account: types.Address(hotWallet.ClassicAddress),
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Currency: currencyCode,
			Issuer:   types.Address(coldWallet.ClassicAddress),
			Value:    "100000000000000",
		},
	}

	flattenedTx = hotColdTrustSet.Flatten()
	err = client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error autofilling transaction: %s\n", err)
		return
	}

	txBlob, _, err = hotWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error signing transaction: %s\n", err)
		return
	}

	response, err = client.SubmitTransactionBlob(txBlob, false)
	if err != nil {
		fmt.Printf("❌ Error submitting transaction: %s\n", err)
		return
	}

	if response.EngineResult != "tesSUCCESS" {
		fmt.Println("❌ Trust line from hot to cold address creation failed!", response.EngineResult)
		fmt.Println("Try again!")
		fmt.Println()
		return
	}

	fmt.Println("✅ Trust line from hot to cold address created!")
	fmt.Printf("🌐 Hash: %s\n", response.Tx["hash"])
	fmt.Println()

	//
	// Send tokens from cold wallet to hot wallet
	//
	fmt.Println("⏳ Sending tokens from cold wallet to hot wallet...")
	coldToHotPayment := &transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: types.Address(coldWallet.ClassicAddress),
		},
		Amount: types.IssuedCurrencyAmount{
			Currency: currencyCode,
			Issuer:   types.Address(coldWallet.ClassicAddress),
			Value:    "3800",
		},
		Destination:    types.Address(hotWallet.ClassicAddress),
		DestinationTag: 1,
	}

	flattenedTx = coldToHotPayment.Flatten()
	err = client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error autofilling transaction: %s\n", err)
		return
	}

	txBlob, _, err = coldWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error signing transaction: %s\n", err)
		return
	}

	response, err = client.SubmitTransactionBlob(txBlob, false)
	if err != nil {
		fmt.Printf("❌ Error submitting transaction: %s\n", err)
		return
	}

	if response.EngineResult != "tesSUCCESS" {
		fmt.Println("❌ Tokens not sent from cold wallet to hot wallet!", response.EngineResult)
		fmt.Println("Try again!")
		fmt.Println()
		return
	}

	fmt.Println("✅ Tokens sent from cold wallet to hot wallet!")
	fmt.Printf("🌐 Hash: %s\n", response.Tx["hash"])
	fmt.Println()

	//
	// Claw back tokens from customer one
	//
	fmt.Println("⏳ Clawing back tokens from hot wallet...")

	coldWalletClawback := &transactions.Clawback{
		BaseTx: transactions.BaseTx{
			Account: types.Address(coldWallet.ClassicAddress),
		},
		Amount: types.IssuedCurrencyAmount{
			Currency: currencyCode,
			Issuer:   types.Address(hotWallet.ClassicAddress),
			Value:    "50",
		},
	}

	flattenedTx = coldWalletClawback.Flatten()
	err = client.Autofill(&flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error autofilling transaction: %s\n", err)
		return
	}

	txBlob, _, err = coldWallet.Sign(flattenedTx)
	if err != nil {
		fmt.Printf("❌ Error signing transaction: %s\n", err)
		return
	}

	response, err = client.SubmitTransactionBlob(txBlob, false)
	if err != nil {
		fmt.Printf("❌ Error submitting transaction: %s\n", err)
		return
	}

	if response.EngineResult != "tesSUCCESS" {
		fmt.Println("❌ Tokens not clawed back from customer one!", response.EngineResult)
		fmt.Println("Try again!")
		return
	}

	fmt.Println("✅ Tokens clawed back from customer one!")
	fmt.Printf("🌐 Hash: %s\n", response.Tx["hash"])
	fmt.Println()
}
