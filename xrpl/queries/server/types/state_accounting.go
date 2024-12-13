package types

type StateAccountingFinal struct {
	Disconnected InfoAccounting `json:"disconnected"`
	Connected    InfoAccounting `json:"connected"`
	Full         InfoAccounting `json:"full"`
	Syncing      InfoAccounting `json:"syncing"`
	Tracking     InfoAccounting `json:"tracking"`
}

type InfoAccounting struct {
	DurationUS  string `json:"duration_us"`
	Transitions string `json:"transitions"`
}
