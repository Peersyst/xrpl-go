package publicservers

// ****************
// Websocket URLs
// ****************

type WSPublicServersUrls struct {
	mainnet mainnetWsUrls
	testnet testnetWsUrls
	devnet  devnetWsUrls
}

// The mainnet public servers.
func (w WSPublicServersUrls) Mainnet() mainnetWsUrls {
	return w.mainnet
}

// The testnet public servers.
func (w WSPublicServersUrls) Testnet() testnetWsUrls {
	return w.testnet
}

// The devnet public servers.
func (w WSPublicServersUrls) Devnet() devnetWsUrls {
	return w.devnet
}

type MainnetWsUrl string

// mainnetWsUrls is a struct representing the subfields for the Mainnet enum.
type mainnetWsUrls struct {
	xrpLedgerFoundation MainnetWsUrl
	rippleS1            MainnetWsUrl
	rippleS2            MainnetWsUrl
}

const (
	xrpLedgerFoundationWs MainnetWsUrl = "wss://xrplcluster.com/"
	rippleS1Ws            MainnetWsUrl = "wss://s1.ripple.com"
	rippleS2Ws            MainnetWsUrl = "wss://s2.ripple.com"
)

// Full history server cluster with CORS support - xrplcluster.com.
func (m mainnetWsUrls) XRPLedgerFoundation() MainnetWsUrl {
	return m.xrpLedgerFoundation
}

// General purpose server cluster - s1.ripple.com.
func (m mainnetWsUrls) RippleS1() MainnetWsUrl {
	return m.rippleS1
}

// Full-history server cluster - s2.ripple.com
func (m mainnetWsUrls) RippleS2() MainnetWsUrl {
	return m.rippleS2
}

// testnetWsUrls is a struct representing the subfields for the Testnet enum.
type testnetWsUrls struct {
	rippleTestnet TestnetWsUrl
	xrplLabs      TestnetWsUrl
	clio          TestnetWsUrl
}

type TestnetWsUrl string

const (
	rippleTestnet TestnetWsUrl = "wss://s.altnet.rippletest.net:51233"
	xrplLabs      TestnetWsUrl = "wss://testnet.xrpl-labs.com"
	clio          TestnetWsUrl = "wss://clio.altnet.rippletest.net:51233"
)

// Testnet public server - s.altnet.rippletest.net:51233
func (t testnetWsUrls) Ripple() TestnetWsUrl {
	return t.rippleTestnet
}

// Testnet public server with CORS support - testnet.xrpl-labs.com
func (t testnetWsUrls) XRPLLabs() TestnetWsUrl {
	return t.xrplLabs
}

// Testnet public server with Clio - clio.altnet.rippletest.net:51233
func (t testnetWsUrls) Clio() TestnetWsUrl {
	return t.clio
}

// devnetWsUrls is a struct representing the subfields for the Devnet enum.
type devnetWsUrls struct {
	rippleDevnet     DevnetUrl
	rippleDevnetClio DevnetUrl
}

type DevnetUrl string

const (
	rippleDevnet     DevnetUrl = "wss://s.devnet.rippletest.net:51233"
	rippleDevnetClio DevnetUrl = "wss://clio.devnet.rippletest.net:51233"
)

// Devnet public server - s.devnet.rippletest.net:51233
func (d devnetWsUrls) RippleDevnet() DevnetUrl {
	return d.rippleDevnet
}

// Devnet public server with Clio - clio.devnet.rippletest.net:51233
func (d devnetWsUrls) RippleDevnetClio() DevnetUrl {
	return d.rippleDevnetClio
}

// func to initialize WSPublicServersUrls struct
func NewWSPublicServersUrls() WSPublicServersUrls {
	return WSPublicServersUrls{
		mainnet: mainnetWsUrls{
			xrpLedgerFoundation: xrpLedgerFoundationWs,
			rippleS1:            rippleS1Ws,
			rippleS2:            rippleS2Ws,
		},
		testnet: testnetWsUrls{
			rippleTestnet: rippleTestnet,
			xrplLabs:      xrplLabs,
			clio:          clio,
		},
		devnet: devnetWsUrls{
			rippleDevnet:     rippleDevnet,
			rippleDevnetClio: rippleDevnetClio,
		},
	}
}

// ****************
// JSON RPC URLs
// ****************

type JsonRpcPublicServersUrls struct {
	mainnet mainnetRpcUrls
	testnet testnetRpcUrls
	devnet  devnetRpcUrls
}

// The mainnet public servers.
func (j JsonRpcPublicServersUrls) Mainnet() mainnetRpcUrls {
	return j.mainnet
}

// The testnet public servers.
func (j JsonRpcPublicServersUrls) Testnet() testnetRpcUrls {
	return j.testnet
}

// The devnet public servers.
func (j JsonRpcPublicServersUrls) Devnet() devnetRpcUrls {
	return j.devnet
}

type MainnetRpcUrl string

// mainnetWsUrls is a struct representing the subfields for the Mainnet enum.
type mainnetRpcUrls struct {
	xrpLedgerFoundation MainnetRpcUrl
	rippleS1            MainnetRpcUrl
	rippleS2            MainnetRpcUrl
}

const (
	xrpLedgerFoundationRpc MainnetRpcUrl = "https://xrplcluster.com"
	rippleS1Rpc            MainnetRpcUrl = "https://s1.ripple.com:51234"
	rippleS2Rpc            MainnetRpcUrl = "https://s2.ripple.com:51234"
)

// Full history server cluster with CORS support - xrplcluster.com.
func (m mainnetRpcUrls) XRPLedgerFoundation() MainnetRpcUrl {
	return m.xrpLedgerFoundation
}

// General purpose server cluster - s1.ripple.com.
func (m mainnetRpcUrls) RippleS1() MainnetRpcUrl {
	return m.rippleS1
}

// Full-history server cluster - s2.ripple.com
func (m mainnetRpcUrls) RippleS2() MainnetRpcUrl {
	return m.rippleS2
}

// testnetWsUrls is a struct representing the subfields for the Testnet enum.
type testnetRpcUrls struct {
	rippleTestnet TestnetRpcUrl
	xrplLabs      TestnetRpcUrl
	clio          TestnetRpcUrl
}

type TestnetRpcUrl string

const (
	rippleTestnetRpc TestnetRpcUrl = "https://s.altnet.rippletest.net:51234"
	xrplLabsRpc      TestnetRpcUrl = "https://testnet.xrpl-labs.com"
	clioRpc          TestnetRpcUrl = "https://clio.altnet.rippletest.net:51234"
)

// Testnet public server - s.altnet.rippletest.net:51234
func (t testnetRpcUrls) Ripple() TestnetRpcUrl {
	return t.rippleTestnet
}

// Testnet public server with CORS support - testnet.xrpl-labs.com
func (t testnetRpcUrls) XRPLLabs() TestnetRpcUrl {
	return t.xrplLabs
}

// Testnet public server with Clio - clio.altnet.rippletest.net:51234
func (t testnetRpcUrls) Clio() TestnetRpcUrl {
	return t.clio
}

// devnetWsUrls is a struct representing the subfields for the Devnet enum.
type devnetRpcUrls struct {
	rippleDevnet     DevnetRpcUrl
	rippleDevnetClio DevnetRpcUrl
}

type DevnetRpcUrl string

const (
	rippleDevnetRpc     DevnetRpcUrl = "https://s.devnet.rippletest.net:51234"
	rippleDevnetClioRpc DevnetRpcUrl = "https://clio.devnet.rippletest.net:51234"
)

// Devnet public server - s.devnet.rippletest.net:51234
func (d devnetRpcUrls) RippleDevnet() DevnetRpcUrl {
	return d.rippleDevnet
}

// Devnet public server with Clio - clio.devnet.rippletest.net:51234
func (d devnetRpcUrls) RippleDevnetClio() DevnetRpcUrl {
	return d.rippleDevnetClio
}

// func to initialize JsonRpcPublicServersUrls struct
func NewJsonRpcPublicServersUrls() JsonRpcPublicServersUrls {
	return JsonRpcPublicServersUrls{
		mainnet: mainnetRpcUrls{
			xrpLedgerFoundation: xrpLedgerFoundationRpc,
			rippleS1:            rippleS1Rpc,
			rippleS2:            rippleS2Rpc,
		},
		testnet: testnetRpcUrls{
			rippleTestnet: rippleTestnetRpc,
			xrplLabs:      xrplLabsRpc,
			clio:          clioRpc,
		},
		devnet: devnetRpcUrls{
			rippleDevnet:     rippleDevnetRpc,
			rippleDevnetClio: rippleDevnetClioRpc,
		},
	}
}
