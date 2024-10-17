package ledger

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type XChainClaimProofSig struct {
	// The amount to claim in the XChainCommit transaction on the destination chain.
	Amount types.CurrencyAmount
	// The account that should receive this signer's share of the SignatureReward.
	AttestationRewardAccount types.Address
	// The account on the door account's signer list that is signing the transaction.
	AttestationSignerAccount types.Address
	// The destination account for the funds on the destination chain.
	Destination types.Address `json:",omitempty"`
	// The public key used to verify the signature.
	PublicKey string
	// A boolean representing the chain where the event occurred.
	WasLockingChainSend uint8
}

type XChainClaimAttestation struct {
	// An attestation from one witness server.
	XChainClaimProofSig XChainClaimProofSig
}

// An XChainOwnedClaimID object represents one cross-chain transfer of value and includes information of the account on
// the source chain that locks or burns the funds on the source chain.
// The XChainOwnedClaimID object must be acquired on the destination chain before submitting a XChainCommit on the source chain.
// Its purpose is to prevent transaction replay attacks and is also used as a place to collect attestations from witness servers.
// An XChainCreateClaimID transaction is used to create a new XChainOwnedClaimID. The ledger object is destroyed when the funds are successfully claimed on the destination chain.
// Example:
// ```json
//
//	{
//	  "Account": "rBW1U7J9mEhEdk6dMHEFUjqQ7HW7WpaEMi",
//	  "Flags": 0,
//	  "OtherChainSource": "r9oXrvBX5aDoyMGkoYvzazxDhYoWFUjz8p",
//	  "OwnerNode": "0",
//	  "PreviousTxnID": "1CFD80E9CF232B8EED62A52857DE97438D12230C06496932A81DEFA6E66070A6",
//	  "PreviousTxnLgrSeq": 58673,
//	  "SignatureReward": "100",
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
//	  "XChainClaimAttestations": [
//	    {
//	      "XChainClaimProofSig": {
//	        "Amount": "1000000",
//	        "AttestationRewardAccount": "rfgjrgEJGDxfUY2U8VEDs7BnB1jiH3ofu6",
//	        "AttestationSignerAccount": "rfsxNxZ6xB1nTPhTMwQajNnkCxWG8B714n",
//	        "Destination": "rBW1U7J9mEhEdk6dMHEFUjqQ7HW7WpaEMi",
//	        "PublicKey": "025CA526EF20567A50FEC504589F949E0E3401C13EF76DD5FD1CC2850FA485BD7B",
//	        "WasLockingChainSend": 1
//	      }
//	    },
//	    {
//	      "XChainClaimProofSig": {
//	        "Amount": "1000000",
//	        "AttestationRewardAccount": "rUUL1tP523M8KimERqVS7sxb1tLLmpndyv",
//	        "AttestationSignerAccount": "rEg5sHxZVTNwRL3BAdMwJatkmWDzHMmzDF",
//	        "Destination": "rBW1U7J9mEhEdk6dMHEFUjqQ7HW7WpaEMi",
//	        "PublicKey": "03D40434A6843638681E2F215310EBC4131AFB12EA85985DA073183B732525F7C9",
//	        "WasLockingChainSend": 1
//	      },
//	    }
//	  ],
//	  "XChainClaimID": "b5",
//	  "LedgerEntryType": "XChainOwnedClaimID",
//	  "LedgerIndex": "20B136D7BF6D2E3D610E28E3E6BE09F5C8F4F0241BBF6E2D072AE1BACB1388F5"
//	}
//
// ```
type XChainOwnedClaimID struct {
	// The account that owns this object.
	Account types.Address
	// 	The ledger index is a hash of a unique prefix for XChainOwnedClaimIDs,
	// the actual XChainClaimID value, and the fields in XChainBridge.
	LedgerIndex string
	// The account that must send the corresponding XChainCommit on the source chain. The destination may be specified
	// in the XChainCommit transaction, which means that if the OtherChainSource isn't specified, another account can
	//try to specify a different destination and steal the funds. This also allows tracking only a single set of signatures,
	//since we know which account will send the XChainCommit transaction.
	OtherChainSource types.Address
	// The total amount to pay the witness servers for their signatures. It must be at least the value of SignatureReward in the Bridge ledger object.
	SignatureReward types.CurrencyAmount
	// The door accounts and assets of the bridge this object correlates to.
	XChainBridge types.XChainBridge
	// Attestations collected from the witness servers. This includes the parameters needed to recreate the message
	// that was signed, including the amount, which chain (locking or issuing), optional destination, and reward account for that signature.
	XChainClaimAttestations []XChainClaimAttestation
	//The unique sequence number for a cross-chain transfer.
	XChainClaimID uint64
}

// EntryType returns the type of the ledger entry.
func (*XChainOwnedClaimID) EntryType() LedgerEntryType {
	return XChainOwnedClaimIDEntry
}
