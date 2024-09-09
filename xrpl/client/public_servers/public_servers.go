package publicservers

// WSPublicServers is a struct representing the enum-like structure.
type WSPublicServers struct {
	mainnet mainnetUrls
	testnet TestnetUrls
	devnet  DevnetUrls
}

// getter to get mainnet
func (w WSPublicServers) Mainnet() mainnetUrls {
	return w.mainnet
}

// getter to get testnet
func (w WSPublicServers) Testnet() TestnetUrls {
	return w.testnet
}

// getter to get devnet
func (w WSPublicServers) Devnet() DevnetUrls {
	return w.devnet
}

type MainnetUrl string

// mainnetUrls is a struct representing the subfields for the Mainnet enum.
type mainnetUrls struct {
	xrpLedgerFoundation MainnetUrl
	rippleS1            MainnetUrl
	rippleS2            MainnetUrl
}

const (
	xrpLedgerFoundation MainnetUrl = "wss://xrplcluster.com/"
	rippleS1            MainnetUrl = "wss://s1.ripple.com/"
	rippleS2            MainnetUrl = "wss://s2.ripple.com/"
)

// Full history server cluster with CORS support - xrplcluster.com.
func (m mainnetUrls) XRPLedgerFoundation() MainnetUrl {
	return xrpLedgerFoundation
}

// General purpose server cluster - s1.ripple.com.
func (m mainnetUrls) RippleS1() MainnetUrl {
	return m.rippleS1
}

// Full-history server cluster - s2.ripple.com
func (m mainnetUrls) RippleS2() MainnetUrl {
	return m.rippleS2
}

// TestnetUrls is a struct representing the subfields for the Testnet enum.
type TestnetUrls struct {
	rippleTestnet TestnetUrl
	xrplLabs      TestnetUrl
	clio          TestnetUrl
}

type TestnetUrl string

const (
	rippleTestnet TestnetUrl = "wss://s.altnet.rippletest.net:51233"
	xrplLabs      TestnetUrl = "wss://testnet.xrpl-labs.com"
	clio          TestnetUrl = "wss://clio.altnet.rippletest.net:51233"
)

// Testnet public server - s.altnet.rippletest.net:51233
func (t TestnetUrls) Ripple() TestnetUrl {
	return t.rippleTestnet
}

// Testnet public server with CORS support - testnet.xrpl-labs.com
func (t TestnetUrls) XRPLLabs() TestnetUrl {
	return t.xrplLabs
}

// Testnet public server with Clio - clio.altnet.rippletest.net:51233
func (t TestnetUrls) Clio() TestnetUrl {
	return t.clio
}

// DevnetUrls is a struct representing the subfields for the Devnet enum.
type DevnetUrls struct {
	rippleDevnet     DevnetUrl
	rippleDevnetClio DevnetUrl
}

type DevnetUrl string

const (
	rippleDevnet     DevnetUrl = "wss://s.devnet.rippletest.net:51233"
	rippleDevnetClio DevnetUrl = "wss://clio.devnet.rippletest.net:51233"
)

// Devnet public server - s.devnet.rippletest.net:51233
func (d DevnetUrls) RippleDevnet() DevnetUrl {
	return d.rippleDevnet
}

// Devnet public server with Clio - clio.devnet.rippletest.net:51233
func (d DevnetUrls) RippleDevnetClio() DevnetUrl {
	return d.rippleDevnetClio
}

// func to initialize WSPublicServers struct
func NewWSPublicServers() WSPublicServers {
	return WSPublicServers{
		mainnet: mainnetUrls{
			xrpLedgerFoundation: xrpLedgerFoundation,
			rippleS1:            rippleS1,
			rippleS2:            rippleS2,
		},
		testnet: TestnetUrls{
			rippleTestnet: rippleTestnet,
			xrplLabs:      xrplLabs,
			clio:          clio,
		},
		devnet: DevnetUrls{
			rippleDevnet:     rippleDevnet,
			rippleDevnetClio: rippleDevnetClio,
		},
	}
}
