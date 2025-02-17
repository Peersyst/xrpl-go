# transaction

## Overview

As the [`queries`](/docs/xrpl/queries) package contains the types of XRPL queries, this package contains all transaction types available in the XRPL. Contains all transaction structs to build and sign transactions with wallets and clients.

## Transaction types

These are the transaction types available in the XRPL:

- [AccountDelete](https://xrpl.org/docs/references/protocol/transactions/types/accountdelete)
- [AccountSet](https://xrpl.org/docs/references/protocol/transactions/types/accountset)
- [AMMBid](https://xrpl.org/docs/references/protocol/transactions/types/ammbid)
- [AMMCreate](https://xrpl.org/docs/references/protocol/transactions/types/ammcreate)
- [AMMDelete](https://xrpl.org/docs/references/protocol/transactions/types/ammdelete)
- [AMMDeposit](https://xrpl.org/docs/references/protocol/transactions/types/ammdeposit)
- [AMMVote](https://xrpl.org/docs/references/protocol/transactions/types/ammvote)
- [AMMWithdraw](https://xrpl.org/docs/references/protocol/transactions/types/ammwithdraw)
- [CheckCancel](https://xrpl.org/docs/references/protocol/transactions/types/checkcancel)
- [CheckCash](https://xrpl.org/docs/references/protocol/transactions/types/checkcash)
- [CheckCreate](https://xrpl.org/docs/references/protocol/transactions/types/checkcreate)
- [Clawback](https://xrpl.org/docs/references/protocol/transactions/types/clawback)
- [DepositPreauth](https://xrpl.org/docs/references/protocol/transactions/types/depositpreauth)
- [DIDDelete](https://xrpl.org/docs/references/protocol/transactions/types/diddelete)
- [DIDSet](https://xrpl.org/docs/references/protocol/transactions/types/didset)
- [EscrowCancel](https://xrpl.org/docs/references/protocol/transactions/types/escrowcancel)
- [EscrowCreate](https://xrpl.org/docs/references/protocol/transactions/types/escrowcreate)
- [EscrowFinish](https://xrpl.org/docs/references/protocol/transactions/types/escrowfinish)
- [NFTokenAcceptOffer](https://xrpl.org/docs/references/protocol/transactions/types/nftokenacceptoffer)
- [NFTokenBurn](https://xrpl.org/docs/references/protocol/transactions/types/nftokenburn)
- [NFTokenCancelOffer](https://xrpl.org/docs/references/protocol/transactions/types/nftokencanceloffer)
- [NFTokenCreateOffer](https://xrpl.org/docs/references/protocol/transactions/types/nftokencreateoffer)
- [NFTokenMint](https://xrpl.org/docs/references/protocol/transactions/types/nftokenmint)
- [OfferCancel](https://xrpl.org/docs/references/protocol/transactions/types/offercancel)
- [OfferCreate](https://xrpl.org/docs/references/protocol/transactions/types/offercreate)
- [OracleDelete](https://xrpl.org/docs/references/protocol/transactions/types/oracledelete)
- [OracleSet](https://xrpl.org/docs/references/protocol/transactions/types/oracleset)
- [PaymentChannelClaim](https://xrpl.org/docs/references/protocol/transactions/types/paymentchannelclaim)
- [PaymentChannelCreate](https://xrpl.org/docs/references/protocol/transactions/types/paymentchannelcreate)
- [PaymentChannelFund](https://xrpl.org/docs/references/protocol/transactions/types/paymentchannelfund)
- [Payment](https://xrpl.org/docs/references/protocol/transactions/types/payment)
- [SetRegularKey](https://xrpl.org/docs/references/protocol/transactions/types/setregularkey)
- [SignerListSet](https://xrpl.org/docs/references/protocol/transactions/types/signerlistset)
- [TicketCreate](https://xrpl.org/docs/references/protocol/transactions/types/ticketcreate)
- [TrustSet](https://xrpl.org/docs/references/protocol/transactions/types/trustset)
- [XChainAccountCreateCommit](https://xrpl.org/docs/references/protocol/transactions/types/xchainaccountcreatecommit)
- [XChainAddAccountCreateAttestation](https://xrpl.org/docs/references/protocol/transactions/types/xchainaddaccountcreateattestation)
- [XChainAddClaimAttestation](https://xrpl.org/docs/references/protocol/transactions/types/xchainaddclaimattestation)
- [XChainClaim](https://xrpl.org/docs/references/protocol/transactions/types/xchainclaim)
- [XChainCommit](https://xrpl.org/docs/references/protocol/transactions/types/xchaincommit)
- [XChainCreateBridge](https://xrpl.org/docs/references/protocol/transactions/types/xchaincreatebridge)
- [XChainCreateClaimID](https://xrpl.org/docs/references/protocol/transactions/types/xchaincreateclaimid)
- [XChainModifyBridge](https://xrpl.org/docs/references/protocol/transactions/types/xchainmodifybridge)

## Usage

To use the `transaction` package, you need to import it in your project:

```go
import "github.com/Peersyst/xrpl-go/xrpl/transaction"
```