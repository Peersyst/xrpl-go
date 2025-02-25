package types

import "github.com/Peersyst/xrpl-go/xrpl/queries/common"

// The LedgerStream only sends ledgerClosed messages when the consensus process declares
// a new validated ledger. The message identifies the ledger and provides some information
// about its contents.
type LedgerStream struct {
	// `ledgerClosed` indicates this is from the ledger stream
	Type Type `json:"type"`
	// The reference transaction cost as of this ledger version, in drops of XRP. If this
	// ledger version includes a SetFee pseudo-transaction the new transaction cost applies
	// starting with the following ledger version.
	FeeBase int `json:"fee_base"`
	// (May be omitted) The reference transaction cost in "fee units". If the XRPFees
	// amendment is enabled, this field is permanently omitted as it will no longer be relevant.
	FeeRef int `json:"fee_ref"`
	// The identifying hash of the ledger version that was closed.
	LedgerHash common.LedgerHash `json:"ledger_hash"`
	// The ledger index of the ledger that was closed.
	LedgerIndex common.LedgerIndex `json:"ledger_index"`
	// The time this ledger was closed, in seconds since the Ripple Epoch.
	LedgerTime uint64 `json:"ledger_time"`
	// The minimum reserve, in drops of XRP, that is required for an account. If this ledger
	// version includes a SetFee pseudo-transaction the new base reserve applies starting with
	// the following ledger version.
	ReserveBase uint `json:"reserve_base"`
	// The owner reserve for each object an account owns in the ledger, in drops of XRP. If
	// the ledger includes a SetFee pseudo-transaction the new owner reserve applies after
	// this ledger.
	ReserveInc uint `json:"reserve_inc"`
	// Number of new transactions included in this ledger version.
	TxnCount int `json:"txn_count"`
	// (May be omitted) Range of ledgers that the server has available. This may be a disjoint
	// sequence such as 24900901-24900984,24901116-24901158. This field is not returned if the
	// server is not connected to the network, or if it is connected but has not yet obtained
	// a ledger from the network.
	ValidatedLedgers string `json:"validated_ledgers,omitempty"`
}
