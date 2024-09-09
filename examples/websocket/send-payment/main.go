package main

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/Peersyst/xrpl-go/xrpl"
	public_servers "github.com/Peersyst/xrpl-go/xrpl/client/public_servers"
	"github.com/Peersyst/xrpl-go/xrpl/client/websocket"
	"github.com/Peersyst/xrpl-go/xrpl/faucet"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/utils"
)

func main() {
	wallet, err := xrpl.NewWalletFromSeed("sEdSMVV4dJ1JbdBxmakRR4Puu3XVZz2", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	receiverWallet, err := xrpl.NewWalletFromSeed("sEd7d8Ci9nevdLCeUMctF3uGXp9WQqJ", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	// init public urls for websocket client
	publicServers := public_servers.NewServerUrls()

	client := websocket.NewWebsocketClient(
		websocket.NewWebsocketClientConfig().
			WithHost(publicServers.TestnetWebSocket().Ripple()).
			WithFaucetProvider(faucet.NewTestnetFaucetProvider()),
	)

	balance, err := client.GetXrpBalance(wallet.GetAddress())

	if err != nil || balance == "0" {
		fmt.Println("Balance: 0")
		fmt.Println("Funding wallet")
		err = client.FundWallet(&wallet)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	balance, _ = client.GetXrpBalance(wallet.GetAddress())

	fmt.Println("Balance: ", balance)

	amount, err := utils.XrpToDrops("1")
	if err != nil {
		fmt.Println(err)
		return
	}

	amountUint, err := strconv.ParseUint(amount, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	payment := transactions.Payment{
		BaseTx: transactions.BaseTx{
			Account: types.Address(wallet.GetAddress()),
			Memos: []transactions.MemoWrapper{
				{
					Memo: transactions.Memo{
						MemoData:   hex.EncodeToString([]byte("Hello, World!")),
						MemoFormat: hex.EncodeToString([]byte("plain")),
						MemoType:   hex.EncodeToString([]byte("message")),
					},
				},
				{
					Memo: transactions.Memo{
						MemoData:   hex.EncodeToString([]byte("Hello, World 2!")),
						MemoFormat: hex.EncodeToString([]byte("text/plain")),
						MemoType:   hex.EncodeToString([]byte("message2")),
					},
				},
			},
		},
		Destination: types.Address(receiverWallet.GetAddress()),
		Amount:      types.XRPCurrencyAmount(amountUint),
	}

	flatTx := payment.Flatten()

	err = client.Autofill(&flatTx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Transaction autofilled")

	txBlob, _, err := wallet.Sign(flatTx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Transaction signed")
	fmt.Println("Transaction submitted\n")

	response, err := client.SubmitTransactionBlob(txBlob, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Transaction engine result:", response.EngineResult)
	fmt.Println("Transaction accepted:", response.Accepted)
	fmt.Println("Transaction hash:", response.Tx["hash"])
}
