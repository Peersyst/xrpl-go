package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type XChainCreateAccountProofSig struct {
	// The amount committed by the XChainAccountCreateCommit transaction on the source chain.
	Amount types.CurrencyAmount
	// The account that should receive this signer's share of the SignatureReward.
	AttestationRewardAccount types.Address
	// The account on the door account's signer list that is signing the transaction.
	AttestationSignerAccount types.Address
	// The destination account for the funds on the destination chain.
	Destination types.Address
	// The public key used to verify the signature.
	PublicKey string
	// A boolean representing the chain where the event occurred.
	WasLockingChainSend uint8
}

type XChainCreateAccountAttestation struct {
	// An attestation from one witness server.
	XChainCreateAccountProofSig XChainCreateAccountProofSig
}

// The XChainOwnedCreateAccountClaimID ledger object is used to collect attestations for creating an
// account via a cross-chain transfer.
// It is created when an XChainAddAccountCreateAttestation transaction adds a signature attesting to
// a XChainAccountCreateCommit transaction and the XChainAccountCreateCount is greater than or equal
// to the current XChainAccountClaimCount on the Bridge ledger object.
// The ledger object is destroyed when all the attestations have been received and the funds have transferred to the new account.
// Example:
// ```json
//
//	{
//	  "LedgerEntryType": "XChainOwnedCreateAccountClaimID",
//	  "LedgerIndex": "5A92F6ED33FDA68FB4B9FD140EA38C056CD2BA9673ECA5B4CEF40F2166BB6F0C",
//	  "Account": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
//	  "XChainAccountCreateCount": "66",
//	  "XChainBridge": {
//	    "IssuingChainDoor": "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
//	    "IssuingChainIssue": {
//	      "currency": "XRP"
//	    },
//	    "LockingChainDoor": "rMAXACCrp3Y8PpswXcg3bKggHX76V3F8M4",
//	    "LockingChainIssue": {
//	      "currency": "XRP"
//	    }
//	  },
//	  "XChainCreateAccountAttestations": [
//	    {
//	      "XChainCreateAccountProofSig": {
//	        "Amount": "20000000",
//	        "AttestationRewardAccount": "rMtYb1vNdeMDpD9tA5qSFm8WXEBdEoKKVw",
//	        "AttestationSignerAccount": "rL8qTrAvZ8Q1o1H9H9Ahpj3xjgmRvFLvJ3",
//	        "Destination": "rBW1U7J9mEhEdk6dMHEFUjqQ7HW7WpaEMi",
//	        "PublicKey": "021F7CC4033EFBE5E8214B04D1BAAEC14808DC6C02F4ACE930A8EF0F5909B0C438",
//	        "SignatureReward": "100",
//	        "WasLockingChainSend": 1
//	      }
//	    }
//	  ]
//	}
//
// ```
type XChainOwnedCreateAccountClaimID struct {
	// The account that owns this object.
	Account types.Address
	// The ledger index is a hash of a unique prefix for XChainOwnedCreateAccountClaimIDs,
	// the actual XChainAccountClaimCount value, and the fields in XChainBridge.
	LedgerIndex string
	// An integer that determines the order that accounts created through cross-chain
	// transfers must be performed. Smaller numbers must execute before larger numbers.
	XChainAccountCreateCount uint64
	// The door accounts and assets of the bridge this object correlates to.
	XChainBridge types.XChainBridge
	// Attestations collected from the witness servers. This includes the parameters needed to recreate the message
	// that was signed, including the amount, destination, signature reward amount, and reward account for that
	// signature. With the exception of the reward account, all signatures must sign the message created with common parameters.
	XChainCreateAccountAttestations []XChainCreateAccountAttestation
}

func (x *XChainOwnedCreateAccountClaimID) EntryType() LedgerEntryType {
	return XChainOwnedCreateAccountClaimIDEntry
}
