package common

type BaseRequest struct {
	Version int `json:"api_version,omitempty"`
}

func (r *BaseRequest) APIVersion() int {
	return r.Version
}

func (r *BaseRequest) SetAPIVersion(apiVersion int) {
	r.Version = apiVersion
}
