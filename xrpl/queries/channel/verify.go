package channel

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

// ############################################################################
// Request
// ############################################################################

// The channel_verify method checks the validity of a signature that can be
// used to redeem a specific amount of XRP from a payment channel.
type VerifyRequest struct {
	Amount    types.XRPCurrencyAmount `json:"amount"`
	ChannelID string                  `json:"channel_id"`
	PublicKey string                  `json:"public_key"`
	Signature string                  `json:"signature"`
}

func (*VerifyRequest) Method() string {
	return "channel_verify"
}

// TODO: Implement V2
func (*VerifyRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the channel_verify method.
type VerifyResponse struct {
	SignatureVerified bool `json:"signature_verified"`
}