package types

type ServerPort struct {
	Port     string   `json:"port"`
	Protocol []string `json:"protocol"`
}
