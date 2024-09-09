package publicservers

// WSPublicServersUrls is a struct representing the enum-like structure.
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
	xrpLedgerFoundation MainnetWsUrl = "wss://xrplcluster.com/"
	rippleS1            MainnetWsUrl = "wss://s1.ripple.com/"
	rippleS2            MainnetWsUrl = "wss://s2.ripple.com/"
)

// Full history server cluster with CORS support - xrplcluster.com.
func (m mainnetWsUrls) XRPLedgerFoundation() MainnetWsUrl {
	return xrpLedgerFoundation
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
			xrpLedgerFoundation: xrpLedgerFoundation,
			rippleS1:            rippleS1,
			rippleS2:            rippleS2,
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
