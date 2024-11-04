package transaction

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
	"github.com/Peersyst/xrpl-go/xrpl/ledger-entry-types"
)

// The SignerListSet transaction creates, replaces, or removes a list of signers that can be used to multi-sign a transaction. This transaction type was introduced by the MultiSign amendment.
//
// Example:
//
// ```json
//
//	{
//	    "Flags": 0,
//	    "TransactionType": "SignerListSet",
//	    "Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
//	    "Fee": "12",
//	    "SignerQuorum": 3,
//	    "SignerEntries": [
//	        {
//	            "SignerEntry": {
//	                "Account": "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
//	                "SignerWeight": 2
//	            }
//	        },
//	        {
//	            "SignerEntry": {
//	                "Account": "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
//	                "SignerWeight": 1
//	            }
//	        },
//	        {
//	            "SignerEntry": {
//	                "Account": "raKEEVSGnKSD9Zyvxu4z6Pqpm4ABH8FS6n",
//	                "SignerWeight": 1
//	            }
//	        }
//	    ]
//	}
//
// `
type SignerListSet struct {
	BaseTx
	// A target number for the signer weights. A multi-signature from this list is valid only if the sum weights of the signatures provided is greater than or equal to this value.
	// To delete a signer list, use the value 0.
	SignerQuorum uint
	// (Omitted when deleting) Array of SignerEntry objects, indicating the addresses and weights of signers in this list.
	// This signer list must have at least 1 member and no more than 32 members.
	// No address may appear more than once in the list, nor may the Account submitting the transaction appear in the list.
	SignerEntries []ledger.SignerEntryWrapper
}

// TxType returns the transaction type for this transaction (SignerListSet).
func (*SignerListSet) TxType() TxType {
	return SignerListSetTx
}

// Flatten returns the flattened map of the SignerListSet transaction.
func (s *SignerListSet) Flatten() FlatTransaction {
	flattened := s.BaseTx.Flatten()

	flattened["TransactionType"] = "SignerListSet"
	flattened["SignerQuorum"] = s.SignerQuorum

	if len(s.SignerEntries) > 0 {
		flattedSignerListEntries := make([]map[string]interface{}, 0)
		for _, signerEntry := range s.SignerEntries {
			flattenedEntry := signerEntry.Flatten()
			if flattenedEntry != nil {
				flattedSignerListEntries = append(flattedSignerListEntries, flattenedEntry)
			}
		}
		flattened["SignerEntries"] = flattedSignerListEntries
	}

	return flattened
}

// At least one account must be part of the SignerList
const MIN_SIGNERS = 1

// A SignerList can have at most 32 signers
const MAX_SIGNERS = 32

// Validate checks if the SignerListSet struct is valid.
func (s *SignerListSet) Validate() (bool, error) {
	_, err := s.BaseTx.Validate()
	if err != nil {
		return false, err
	}

	// All other checks are for if SignerQuorum is greater than 0
	if s.SignerQuorum == 0 {
		return true, nil
	}

	// Check if SignerEntries has at least 1 entry and no more than 32 entries
	if len(s.SignerEntries) < MIN_SIGNERS || len(s.SignerEntries) > MAX_SIGNERS {
		return false, fmt.Errorf("signerEntries must have at least %d entry and no more than %d entries", MIN_SIGNERS, MAX_SIGNERS)
	}

	// Check if WalletLocator is an hexadecimal string for each SignerEntry
	for _, signerEntry := range s.SignerEntries {
		if signerEntry.SignerEntry.WalletLocator != "" && !typecheck.IsHex(signerEntry.SignerEntry.WalletLocator.String()) {
			return false, fmt.Errorf("invalid WalletLocator in SignerEntry, must be an hexadecimal string")
		}
	}

	// Check SignerQuorum is less than or equal to the sum of all SignerWeights
	sumSignerWeights := uint16(0)
	for _, signerEntry := range s.SignerEntries {
		sumSignerWeights += signerEntry.SignerEntry.SignerWeight
	}
	if s.SignerQuorum > uint(sumSignerWeights) {
		return false, fmt.Errorf("signerQuorum must be less than or equal to the sum of all SignerWeights")
	}

	return true, nil
}
