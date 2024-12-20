package server

import "github.com/Peersyst/xrpl-go/xrpl/queries/server/types"

// ############################################################################
// Feature All Request
// ############################################################################

// The feature command returns information about amendments this server knows
// about, including whether they are enabled.
type FeatureAllRequest struct {
}

func (*FeatureAllRequest) Method() string {
	return "feature"
}

// TODO: Implement V2
func (*FeatureAllRequest) Validate() error {
	return nil
}

// ############################################################################
// Feature All Response
// ############################################################################

// The feature command returns information about amendments this server knows
// about, including whether they are enabled.
type FeatureAllResponse struct {
	Features map[string]types.FeatureStatus `json:"features"`
}

// ############################################################################
// Feature One Request
// ############################################################################

type FeatureOneRequest struct {
	Feature string `json:"feature"`
}

func (*FeatureOneRequest) Method() string {
	return "feature"
}

// TODO: Implement V2
func (*FeatureOneRequest) Validate() error {
	return nil
}

// ############################################################################
// Feature One Response
// ############################################################################

// The expected response from the feature method.
type FeatureResponse map[string]types.FeatureStatus
