package transaction

type TxType string

const (
	AccountSetTx           TxType = "AccountSet"
	AccountDeleteTx        TxType = "AccountDelete"
	AMMBidTx               TxType = "AMMBid"
	AMMCreateTx            TxType = "AMMCreate"
	AMMDepositTx           TxType = "AMMDeposit"
	AMMVoteTx              TxType = "AMMVote"
	AMMWithdrawTx          TxType = "AMMWithdraw"
	CheckCancelTx          TxType = "CheckCancel"
	CheckCashTx            TxType = "CheckCash"
	CheckCreateTx          TxType = "CheckCreate"
	ClawbackTx             TxType = "Clawback"
	DepositPreauthTx       TxType = "DepositPreauth"
	EscrowCancelTx         TxType = "EscrowCancel"
	EscrowCreateTx         TxType = "EscrowCreate"
	EscrowFinishTx         TxType = "EscrowFinish"
	NFTokenAcceptOfferTx   TxType = "NFTokenAcceptOffer"
	NFTokenBurnTx          TxType = "NFTokenBurn"
	NFTokenCancelOfferTx   TxType = "NFTokenCancelOffer"
	NFTokenCreateOfferTx   TxType = "NFTokenCreateOffer"
	NFTokenMintTx          TxType = "NFTokenMint"
	OfferCreateTx          TxType = "OfferCreate"
	OfferCancelTx          TxType = "OfferCancel"
	PaymentTx              TxType = "Payment"
	PaymentChannelClaimTx  TxType = "PaymentChannelClaim"
	PaymentChannelCreateTx TxType = "PaymentChannelCreate"
	PaymentChannelFundTx   TxType = "PaymentChannelFund"
	SetRegularKeyTx        TxType = "SetRegularKey"
	SignerListSetTx        TxType = "SignerListSet"
	TrustSetTx             TxType = "TrustSet"
	TicketCreateTx         TxType = "TicketCreate"
	HashedTx               TxType = "HASH"   // TX stored as a string, rather than complete tx obj
	BinaryTx               TxType = "BINARY" // TX stored as a string, json tagged as 'tx_blob'
)

func (t TxType) String() string {
	return string(t)
}
