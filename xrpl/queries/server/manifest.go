package server

import (
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/version"
)

type ManifestDetails struct {
	Domain       string `json:"domain"`
	EphemeralKey string `json:"ephemeral_key"`
	MasterKey    string `json:"master_key"`
	Seq          uint   `json:"seq"`
}

// ############################################################################
// Request
// ############################################################################

// The manifest method reports the current "manifest" information for a given
// validator public key. The "manifest" is the public portion of that
// validator's configured token.
type ManifestRequest struct {
	common.BaseRequest
	PublicKey string `json:"public_key"`
}

func (*ManifestRequest) Method() string {
	return "manifest"
}

func (*ManifestRequest) APIVersion() int {
	return version.RippledAPIV2
}

// TODO: Implement V2
func (*ManifestRequest) Validate() error {
	return nil
}

// ############################################################################
// Response
// ############################################################################

// The expected response from the manifest method.
type ManifestResponse struct {
	Details   ManifestDetails `json:"details,omitempty"`
	Manifest  string          `json:"manifest,omitempty"`
	Requested string          `json:"requested"`
}
