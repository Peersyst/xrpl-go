"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[420],{7928:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>l,default:()=>h,frontMatter:()=>o,metadata:()=>i,toc:()=>a});const i=JSON.parse('{"id":"xrpl/websocket","title":"websocket","description":"Overview","source":"@site/docs/xrpl/websocket.md","sourceDirName":"xrpl","slug":"/xrpl/websocket","permalink":"/docs/xrpl/websocket","draft":false,"unlisted":false,"editUrl":"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/xrpl/websocket.md","tags":[],"version":"current","frontMatter":{},"sidebar":"tutorialSidebar","previous":{"title":"wallet","permalink":"/docs/xrpl/wallet"}}');var r=n(5105),s=n(3461);const o={},l="websocket",c={},a=[{value:"Overview",id:"overview",level:2},{value:"Config",id:"config",level:2},{value:"Host",id:"host",level:3},{value:"FaucetProvider",id:"faucetprovider",level:3},{value:"MaxRetries",id:"maxretries",level:3},{value:"RetryDelay",id:"retrydelay",level:3},{value:"FeeCushion",id:"feecushion",level:3},{value:"MaxFeeXRP",id:"maxfeexrp",level:3},{value:"Connection",id:"connection",level:2},{value:"Methods",id:"methods",level:2},{value:"Request",id:"request",level:3},{value:"Autofill/AutofillMultisigned",id:"autofillautofillmultisigned",level:3},{value:"Submit/SubmitMultisigned",id:"submitsubmitmultisigned",level:3},{value:"SubmitAndWait",id:"submitandwait",level:3},{value:"Queries",id:"queries",level:2},{value:"Examples",id:"examples",level:2},{value:"How to send a payment transaction",id:"how-to-send-a-payment-transaction",level:3}];function d(e){const t={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",header:"header",li:"li",p:"p",pre:"pre",ul:"ul",...(0,s.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(t.header,{children:(0,r.jsx)(t.h1,{id:"websocket",children:"websocket"})}),"\n",(0,r.jsx)(t.h2,{id:"overview",children:"Overview"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"websocket"})," package provides a WebSocket client for interacting with the XRPL network via its WebSocket API. This client handles the communication with XRPL nodes, allowing you to:"]}),"\n",(0,r.jsxs)(t.ul,{children:["\n",(0,r.jsx)(t.li,{children:"Send requests to query the ledger state."}),"\n",(0,r.jsx)(t.li,{children:"Submit transactions to the network."}),"\n",(0,r.jsx)(t.li,{children:"Receive responses and handle errors."}),"\n",(0,r.jsx)(t.li,{children:"Manage the connections configuration."}),"\n"]}),"\n",(0,r.jsx)(t.h2,{id:"config",children:"Config"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"websocket"})," package provides a ",(0,r.jsx)(t.code,{children:"Config"})," struct that allows you to configure the WebSocket client. Every time you create a new ",(0,r.jsx)(t.code,{children:"Client"}),", you need to pass a ",(0,r.jsx)(t.code,{children:"Config"})," struct as an argument. You can initialize a ",(0,r.jsx)(t.code,{children:"Config"})," struct using the ",(0,r.jsx)(t.code,{children:"NewClientConfig"})," function."]}),"\n",(0,r.jsxs)(t.p,{children:[(0,r.jsx)(t.code,{children:"Config"})," struct follows the options pattern, so you can pass different options to the ",(0,r.jsx)(t.code,{children:"NewClientConfig"})," function:"]}),"\n",(0,r.jsx)(t.h3,{id:"host",children:"Host"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"WithHost"})," option allows you to set the host of the WebSocket client."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:"func (wc ClientConfig) WithHost(host string) ClientConfig\n"})}),"\n",(0,r.jsx)(t.h3,{id:"faucetprovider",children:"FaucetProvider"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"FaucetProvider"})," option allows you to set the faucet provider of the WebSocket client. There're two predefined faucet providers: ",(0,r.jsx)(t.code,{children:"TestnetFaucetProvider"})," and ",(0,r.jsx)(t.code,{children:"DevnetFaucetProvider"}),". You can also implement your own faucet provider by implementing the ",(0,r.jsx)(t.code,{children:"FaucetProvider"})," interface."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:"func (wc ClientConfig) WithFaucetProvider(fp common.FaucetProvider) ClientConfig\n"})}),"\n",(0,r.jsx)(t.h3,{id:"maxretries",children:"MaxRetries"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"WithMaxRetries"})," option allows you to set the maximum number of retries for a transaction."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:"func (wc ClientConfig) WithMaxRetries(maxRetries int) ClientConfig\n"})}),"\n",(0,r.jsx)(t.h3,{id:"retrydelay",children:"RetryDelay"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"WithRetryDelay"})," option allows you to set the delay between retries for a transaction."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:"func (wc ClientConfig) WithRetryDelay(retryDelay time.Duration) ClientConfig\n"})}),"\n",(0,r.jsx)(t.h3,{id:"feecushion",children:"FeeCushion"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"WithFeeCushion"})," option allows you to set the fee cushion for a transaction."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:"func (wc ClientConfig) WithFeeCushion(feeCushion float32) ClientConfig\n"})}),"\n",(0,r.jsx)(t.h3,{id:"maxfeexrp",children:"MaxFeeXRP"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"WithMaxFeeXRP"})," option allows you to set the maximum fee in XRP that the WebSocket client will use."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:"func (wc ClientConfig) WithMaxFeeXRP(maxFeeXrp float32) ClientConfig\n"})}),"\n",(0,r.jsx)(t.h2,{id:"connection",children:"Connection"}),"\n",(0,r.jsxs)(t.p,{children:["As the ",(0,r.jsx)(t.code,{children:"websocket"})," package is a WebSocket client, it needs to be connected to a WebSocket server. The ",(0,r.jsx)(t.code,{children:"Client"})," type exposes the following methods to connect to a WebSocket server:"]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:"// Connection methods\nfunc (c *Client) Connect() error\nfunc (c *Client) Disconnect() error\n\n// Connection status\nfunc (c *Client) IsConnected() bool\n\n// Connection\nfunc (c *Client) Conn() *websocket.Conn\n"})}),"\n",(0,r.jsxs)(t.p,{children:["So, for example, if you want to connect to the ",(0,r.jsx)(t.code,{children:"devnet"})," ledger, you can do it this way:"]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:'client := websocket.NewClient(websocket.NewClientConfig().WithHost("wss://s.altnet.rippletest.net:51233"))\ndefer client.Disconnect()\n\nerr := client.Connect()\nif err != nil {\n    // ...\n}\n\nif !client.IsConnected() {\n    // ...\n}\n'})}),"\n",(0,r.jsx)(t.h2,{id:"methods",children:"Methods"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"Client"})," type exposes the following methods to interact with the XRPL network:"]}),"\n",(0,r.jsx)(t.h3,{id:"request",children:"Request"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"Request"})," method is used to send a request to the server and returns the response. This method is mostly used to send client ",(0,r.jsx)(t.a,{href:"/docs/xrpl/queries",children:(0,r.jsx)(t.code,{children:"queries"})})," to the server."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:"func (c *Client) Request(reqParams XRPLRequest) (*ClientResponse, error)\n"})}),"\n",(0,r.jsx)(t.h3,{id:"autofillautofillmultisigned",children:"Autofill/AutofillMultisigned"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"Autofill"})," method is used to autofill some fields in a flat transaction. This method is useful for adding dynamic fields like ",(0,r.jsx)(t.code,{children:"LastLedgerSequence"})," or ",(0,r.jsx)(t.code,{children:"Fee"}),". It returns an error if the transaction is not valid or some internall call fails. There's also a ",(0,r.jsx)(t.code,{children:"AutofillMultisigned"})," method that works the same way but for multisigned transactions."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:"func (c *Client) Autofill(tx *transaction.FlatTransaction) error\nfunc (c *Client) AutofillMultisigned(tx *transaction.FlatTransaction, nSigners uint64) error\n"})}),"\n",(0,r.jsx)(t.h3,{id:"submitsubmitmultisigned",children:"Submit/SubmitMultisigned"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"Submit"})," method is used to submit a transaction to the XRPL network. It returns a ",(0,r.jsx)(t.code,{children:"TxResponse"})," struct containing the transaction result for the blob submitted. ",(0,r.jsx)(t.code,{children:"txBlob"})," must be signed. There's also a ",(0,r.jsx)(t.code,{children:"SubmitMultisigned"})," method that works the same way but for multisigned transactions."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:"func (c *Client) Submit(txBlob string, failHard bool) (*requests.SubmitResponse, error)\nfunc (c *Client) SubmitMultisigned(txBlob string, failHard bool) (*requests.SubmitMultisignedResponse, error)\n"})}),"\n",(0,r.jsx)(t.h3,{id:"submitandwait",children:"SubmitAndWait"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"SubmitAndWait"})," method is used to submit a transaction to the XRPL network and wait for it to be included in a ledger. It returns a ",(0,r.jsx)(t.code,{children:"TxResponse"})," struct containing the transaction result for the blob submitted."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:"func (c *Client) SubmitAndWait(txBlob string, failHard bool) (*requests.TxResponse, error)\n"})}),"\n",(0,r.jsx)(t.h2,{id:"queries",children:"Queries"}),"\n",(0,r.jsxs)(t.p,{children:["The ",(0,r.jsx)(t.code,{children:"websocket"})," package provides query wrappers that allows you to send client ",(0,r.jsx)(t.a,{href:"/docs/xrpl/queries",children:(0,r.jsx)(t.code,{children:"queries"})})," to the server."]}),"\n",(0,r.jsx)(t.h2,{id:"examples",children:"Examples"}),"\n",(0,r.jsx)(t.h3,{id:"how-to-send-a-payment-transaction",children:"How to send a payment transaction"}),"\n",(0,r.jsxs)(t.p,{children:["This example shows how to send a payment transaction to the XRPL testnet with the ",(0,r.jsx)(t.code,{children:"websocket"})," package."]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-go",children:'package main\n\nimport (\n\t"fmt"\n\t"strconv"\n\n\t"github.com/Peersyst/xrpl-go/xrpl/currency"\n\t"github.com/Peersyst/xrpl-go/xrpl/faucet"\n\t"github.com/Peersyst/xrpl-go/xrpl/transaction"\n\t"github.com/Peersyst/xrpl-go/xrpl/transaction/types"\n\t"github.com/Peersyst/xrpl-go/xrpl/wallet"\n\t"github.com/Peersyst/xrpl-go/pkg/crypto"\n\t"github.com/Peersyst/xrpl-go/xrpl/websocket"\n)\n\nfunc main() {\n\n\t// Create a new websocket client with a testnet faucet provider\n\tclient := websocket.NewClient(\n\t\twebsocket.NewClientConfig().\n\t\t\tWithHost("wss://s.altnet.rippletest.net:51233").\n\t\t\tWithFaucetProvider(faucet.NewTestnetFaucetProvider()),\n\t)\n\tdefer client.Disconnect()\n\n\t// Connect to the testnet\n\tif err := client.Connect(); err != nil {\n\t\tfmt.Println(err)\n\t\treturn\n\t}\n\n\t// Check if the client is connected\n\tif !client.IsConnected() {\n\t\treturn\n\t}\n\n\t// Create a new wallet with the ed25519 algorithm\n\tw, err := wallet.New(crypto.ED25519())\n\tif err != nil {\n\t\tfmt.Println(err)\n\t\treturn\n\t}\n\n\t// Fund the wallet with the testnet faucet\n\tif err := client.FundWallet(&w); err != nil {\n\t\tfmt.Println(err)\n\t\treturn\n\t}\n\n\t// Convert the amount to drops\n\txrpAmount, err := currency.XrpToDrops("1")\n\tif err != nil {\n\t\tfmt.Println(err)\n\t\treturn\n\t}\n\n\txrpAmountInt, err := strconv.ParseInt(xrpAmount, 10, 64)\n\tif err != nil {\n\t\tfmt.Println(err)\n\t\treturn\n\t}\n\n\tp := &transaction.Payment{\n\t\tBaseTx: transaction.BaseTx{\n\t\t\tAccount: types.Address(w.GetAddress()),\n\t\t},\n\t\tDestination: "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe",\n\t\tAmount:      types.XRPCurrencyAmount(xrpAmountInt),\n\t\tDeliverMax:  types.XRPCurrencyAmount(xrpAmountInt),\n\t}\n\n\tflattenedTx := p.Flatten()\n\n\t// Autofill the transaction with the client\'s config\n\tif err := client.Autofill(&flattenedTx); err != nil {\n\t\tfmt.Println(err)\n\t\treturn\n\t}\n\n\t// Sign the transaction with the wallet\n\ttxBlob, _, err := w.Sign(flattenedTx)\n\tif err != nil {\n\t\tfmt.Println(err)\n\t\treturn\n\t}\n\n\t// Submit the transaction to the network and wait for it to be included in a ledge\n\tres, err := client.SubmitAndWait(txBlob, false)\n\tif err != nil {\n\t\tfmt.Println(err)\n\t\treturn\n\t}\n}\n\n'})})]})}function h(e={}){const{wrapper:t}={...(0,s.R)(),...e.components};return t?(0,r.jsx)(t,{...e,children:(0,r.jsx)(d,{...e})}):d(e)}},3461:(e,t,n)=>{n.d(t,{R:()=>o,x:()=>l});var i=n(8101);const r={},s=i.createContext(r);function o(e){const t=i.useContext(s);return i.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function l(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:o(e.components),i.createElement(s.Provider,{value:t},e.children)}}}]);