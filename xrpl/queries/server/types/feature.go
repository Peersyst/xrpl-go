package types

type FeatureStatus struct {
	Enabled   bool   `json:"enabled"`
	Name      string `json:"name"`
	Supported bool   `json:"supported"`
}
