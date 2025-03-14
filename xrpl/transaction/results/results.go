package results

// TxResult represents the result code of a transaction
type TxResult string

const (
	// Success results - https://xrpl.org/docs/references/protocol/transactions/transaction-results/tes-success

	// The transaction was applied and forwarded to other servers.
	// If this appears in a validated ledger, then the transaction's success is final.
	TesSUCCESS TxResult = "tesSUCCESS"

	// ------------------------------------------------------------------------------------------------
	// tec codes ⬇️ - https://xrpl.org/docs/references/protocol/transactions/transaction-results/tec-codes

	// These codes indicate that the transaction failed, but it was applied to a ledger to apply the transaction cost. They have numerical values in the range 100 to 199. It is recommended to use the text code, not the numeric value.

	// Transactions with tec codes destroy the XRP paid as a transaction cost, and consume a sequence number. For the most part, the transactions take no other action, but there are some exceptions. For example, a transaction that results in tecOVERSIZE still cleans up some unfunded offers. Always look at the transaction metadata to see precisely what a transaction did.
	// ------------------------------------------------------------------------------------------------

	// The transaction tried to create an object (such as an Offer or a Check) whose provided Expiration time has already passed.
	TecEXPIRED   TxResult = "tecEXPIRED"
	TecDUPLICATE TxResult = "tecDUPLICATE"

	// ------------------------------------------------------------------------------------------------
	// tem codes ⬇️ - https://xrpl.org/docs/references/protocol/transactions/transaction-results/tem-codes
	//
	// These codes indicate that the transaction was malformed, and cannot succeed according to the XRP Ledger protocol.
	// They have numerical values in the range -299 to -200. The exact code for any given error is subject to change, so don't rely on it.
	// ------------------------------------------------------------------------------------------------

	// The transaction is otherwise invalid. For example, the transaction ID may not be the right format, the signature may not be formed properly, or something else went wrong in understanding the transaction.
	TemINVALID TxResult = "temINVALID"

	// ------------------------------------------------------------------------------------------------
	// tef codes ⬇️ - https://xrpl.org/docs/references/protocol/transactions/transaction-results/tef-codes
	//
	// These codes indicate that the transaction failed and was not included in a ledger, but the transaction could have succeeded in some theoretical ledger.
	// Typically this means that the transaction can no longer succeed in any future ledger. They have numerical values in the range -199 to -100. The exact code for any given error is subject to change, so don't rely on it.
	// ------------------------------------------------------------------------------------------------

	// The sequence number of the transaction is lower than the current sequence number of the account sending the transaction.
	//revive:disable-next-line:var-naming
	TefPAST_SEQ TxResult = "tefPAST_SEQ"
)

// String returns the string representation of the result
func (t TxResult) String() string {
	return string(t)
}
