package transaction

import (
	"github.com/Peersyst/xrpl-go/v1/xrpl/ledger-entry-types"
)

// Vote on the trading fee for an Automated Market Maker instance. Up to 8 accounts can vote in proportion to the amount of the AMM's LP Tokens they hold.
// Each new vote re-calculates the AMM's trading fee based on a weighted average of the votes.
//
// Example:
//
// ```json
//
//	{
//	    "Account" : "rJVUeRqDFNs2xqA7ncVE6ZoAhPUoaJJSQm",
//	    "Asset" : {
//	        "currency" : "XRP"
//	    },
//	    "Asset2" : {
//	        "currency" : "TST",
//	        "issuer" : "rP9jPyP5kyvFRb6ZiRghAGw5u8SGAmU4bd"
//	    },
//	    "Fee" : "10",
//	    "Flags" : 2147483648,
//	    "Sequence" : 8,
//	    "TradingFee" : 600,
//	    "TransactionType" : "AMMVote"
//	}
//
// ```
type AMMVote struct {
	BaseTx
	// The definition for one of the assets in the AMM's pool. In JSON, this is an object with currency and issuer fields (omit issuer for XRP).
	Asset ledger.Asset
	// The definition for the other asset in the AMM's pool. In JSON, this is an object with currency and issuer fields (omit issuer for XRP).
	Asset2 ledger.Asset
	// The proposed fee to vote for, in units of 1/100,000; a value of 1 is equivalent to 0.001%. The maximum value is 1000, indicating a 1% fee.
	TradingFee uint16
}

// TxType returns the type of the transaction (AMMVote).
func (*AMMVote) TxType() TxType {
	return AMMVoteTx
}

// Flatten returns the flattened map of the AMMVote transaction.
func (a *AMMVote) Flatten() FlatTransaction {
	// Add BaseTx fields
	flattened := a.BaseTx.Flatten()

	// Add AMMDelete-specific fields
	flattened["TransactionType"] = "AMMVote"
	flattened["Asset"] = a.Asset.Flatten()
	flattened["Asset2"] = a.Asset2.Flatten()
	flattened["TradingFee"] = a.TradingFee

	return flattened
}

// Validates the AMMVote struct and make sure all the fields are correct.
func (a *AMMVote) Validate() (bool, error) {
	_, err := a.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	if ok, err := IsAsset(a.Asset); !ok {
		return false, err
	}

	if ok, err := IsAsset(a.Asset2); !ok {
		return false, err
	}

	if a.TradingFee > 1000 {
		return false, ErrAMMTradingFeeTooHigh
	}

	return true, nil
}
