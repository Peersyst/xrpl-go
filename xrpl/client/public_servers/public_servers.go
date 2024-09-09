package publicservers

// ****************
// Public Servers URLs
// ****************

// ServerUrls holds both WebSocket and JSON-RPC URLs for different networks.
type ServerUrls struct {
	mainnet mainnetUrls
	testnet testnetUrls
	devnet  devnetUrls
}

// mainnetUrls holds the WebSocket and JSON-RPC URLs for the mainnet.
type mainnetUrls struct {
	webSocket mainnetWebSocketUrls
	jsonRpc   mainnetJsonRpcUrls
}

// testnetUrls holds the WebSocket and JSON-RPC URLs for the testnet.
type testnetUrls struct {
	webSocket testnetWebSocketUrls
	jsonRpc   testnetJsonRpcUrls
}

// devnetUrls holds the WebSocket and JSON-RPC URLs for the devnet.
type devnetUrls struct {
	webSocket devnetWebSocketUrls
	jsonRpc   devnetJsonRpcUrls
}

// WebSocket URLs Structs

type mainnetWebSocketUrls struct {
	xrpLedgerFoundation string
	rippleS1            string
	rippleS2            string
}

type testnetWebSocketUrls struct {
	ripple   string
	xrplLabs string
	clio     string
}

type devnetWebSocketUrls struct {
	ripple     string
	rippleClio string
}

// JSON-RPC URLs Structs

type mainnetJsonRpcUrls struct {
	xrpLedgerFoundation string
	rippleS1            string
	rippleS2            string
}

type testnetJsonRpcUrls struct {
	ripple   string
	xrplLabs string
	clio     string
}

type devnetJsonRpcUrls struct {
	ripple     string
	rippleClio string
}

// NewServerUrls initializes WebSocket and JSON-RPC URLs for all networks, considering the structural differences.
func NewServerUrls() ServerUrls {
	return ServerUrls{
		mainnet: mainnetUrls{
			webSocket: mainnetWebSocketUrls{
				xrpLedgerFoundation: "wss://xrplcluster.com/",
				rippleS1:            "wss://s1.ripple.com",
				rippleS2:            "wss://s2.ripple.com",
			},
			jsonRpc: mainnetJsonRpcUrls{
				xrpLedgerFoundation: "https://xrplcluster.com",
				rippleS1:            "https://s1.ripple.com:51234",
				rippleS2:            "https://s2.ripple.com:51234",
			},
		},
		testnet: testnetUrls{
			webSocket: testnetWebSocketUrls{
				ripple:   "wss://s.altnet.rippletest.net:51233",
				xrplLabs: "wss://testnet.xrpl-labs.com",
				clio:     "wss://clio.altnet.rippletest.net:51233",
			},
			jsonRpc: testnetJsonRpcUrls{
				ripple:   "https://s.altnet.rippletest.net:51234",
				xrplLabs: "https://testnet.xrpl-labs.com",
				clio:     "https://clio.altnet.rippletest.net:51234",
			},
		},
		devnet: devnetUrls{
			webSocket: devnetWebSocketUrls{
				ripple:     "wss://s.devnet.rippletest.net:51233",
				rippleClio: "wss://clio.devnet.rippletest.net:51233",
			},
			jsonRpc: devnetJsonRpcUrls{
				ripple:     "https://s.devnet.rippletest.net:51234",
				rippleClio: "https://clio.devnet.rippletest.net:51234",
			},
		},
	}
}

// Accessor (getter) methods for ServerUrls struct

func (s ServerUrls) MainnetWebSocket() mainnetWebSocketUrls {
	return s.mainnet.webSocket
}

func (s ServerUrls) MainnetJsonRpc() mainnetJsonRpcUrls {
	return s.mainnet.jsonRpc
}

func (s ServerUrls) TestnetWebSocket() testnetWebSocketUrls {
	return s.testnet.webSocket
}

func (s ServerUrls) TestnetJsonRpc() testnetJsonRpcUrls {
	return s.testnet.jsonRpc
}

func (s ServerUrls) DevnetWebSocket() devnetWebSocketUrls {
	return s.devnet.webSocket
}

func (s ServerUrls) DevnetJsonRpc() devnetJsonRpcUrls {
	return s.devnet.jsonRpc
}

// Getter methods for WebSocket URLs

// Full history server cluster with CORS support - wss://xrplcluster.com/
func (m mainnetWebSocketUrls) XRPLedgerFoundation() string {
	return m.xrpLedgerFoundation
}

// General purpose server cluster - wss://s1.ripple.com
func (m mainnetWebSocketUrls) RippleS1() string {
	return m.rippleS1
}

// Full-history server cluster - wss://s2.ripple.com
func (m mainnetWebSocketUrls) RippleS2() string {
	return m.rippleS2
}

// Testnet public server - wss://s.altnet.rippletest.net:51233
func (t testnetWebSocketUrls) Ripple() string {
	return t.ripple
}

// Testnet public server with CORS support - wss://testnet.xrpl-labs.com
func (t testnetWebSocketUrls) XRPLLabs() string {
	return t.xrplLabs
}

// Testnet public server with Clio - wss://clio.altnet.rippletest.net:51233
func (t testnetWebSocketUrls) Clio() string {
	return t.clio
}

// Devnet public server - wss://s.devnet.rippletest.net:51233
func (d devnetWebSocketUrls) Ripple() string {
	return d.ripple
}

// Devnet public server with Clio - wss://clio.devnet.rippletest.net:51233
func (d devnetWebSocketUrls) RippleDevnetClio() string {
	return d.rippleClio
}

// Getter methods for JSON-RPC URLs

// Full history server cluster with CORS support - https://xrplcluster.com
func (m mainnetJsonRpcUrls) XRPLedgerFoundation() string {
	return m.xrpLedgerFoundation
}

// General purpose server cluster - https://s1.ripple.com:51234
func (m mainnetJsonRpcUrls) RippleS1() string {
	return m.rippleS1
}

// Full-history server cluster - https://s2.ripple.com:51234
func (m mainnetJsonRpcUrls) RippleS2() string {
	return m.rippleS2
}

// Testnet public server - https://s.altnet.rippletest.net:51234
func (t testnetJsonRpcUrls) Ripple() string {
	return t.ripple
}

// Testnet public server with CORS support - https://testnet.xrpl-labs.com
func (t testnetJsonRpcUrls) XRPLLabs() string {
	return t.xrplLabs
}

// Testnet public server with Clio - https://clio.altnet.rippletest.net:51234
func (t testnetJsonRpcUrls) Clio() string {
	return t.clio
}

// Devnet public server - https://s.devnet.rippletest.net:51234
func (d devnetJsonRpcUrls) Ripple() string {
	return d.ripple
}

// Devnet public server with Clio - https://clio.devnet.rippletest.net:51234
func (d devnetJsonRpcUrls) RippleClio() string {
	return d.rippleClio
}
