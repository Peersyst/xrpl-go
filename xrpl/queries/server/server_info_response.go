package server

import "github.com/Peersyst/xrpl-go/xrpl/transaction/types"

type InfoResponse struct {
	Info Info `json:"info"`
}

type Info struct {
	AmendmentBlocked        bool                      `json:"amendment_blocked,omitempty"`
	BuildVersion            string                    `json:"build_version"`
	ClosedLedger            *LedgerInfo               `json:"closed_ledger,omitempty"`
	CompleteLedgers         string                    `json:"complete_ledgers"`
	HostID                  string                    `json:"hostid"`
	IOLatencyMS             uint                      `json:"io_latency_ms"`
	JQTransOverflow         string                    `json:"jq_trans_overflow"`
	LastClose               *Close                    `json:"last_close"`
	Load                    *Load                     `json:"load,omitempty"`
	LoadFactor              uint                      `json:"load_factor"`
	LoadFactorLocal         uint                      `json:"load_factor_local,omitempty"`
	LoadFactorNet           uint                      `json:"load_factor_net,omitempty"`
	LoadFactorCluster       uint                      `json:"load_factor_cluster,omitempty"`
	LoadFactorFeeEscelation uint                      `json:"load_factor_fee_escelation,omitempty"`
	LoadFactorFeeQueue      uint                      `json:"load_factor_fee_queue,omitempty"`
	LoadFactorServer        uint                      `json:"load_factor_server,omitempty"`
	Peers                   uint                      `json:"peers,omitempty"`
	PubkeyNode              string                    `json:"pubkey_node"`
	PubkeyValidator         string                    `json:"pubkey_validator,omitempty"`
	Reporting               *Reporting                `json:"reporting,omitempty"`
	ServerState             string                    `json:"server_state"`
	ServerStateDurationUS   string                    `json:"server_state_duration_us"`
	StateAccounting         map[string]InfoAccounting `json:"state_accounting"`
	Time                    string                    `json:"time"`
	Uptime                  uint                      `json:"uptime"`
	ValidatedLedger         *LedgerInfo               `json:"validated_ledger,omitempty"`
	ValidationQuorum        uint                      `json:"validation_quorum"`
	ValidatorListExpires    string                    `json:"validator_list_expires,omitempty"`
}

type InfoAccounting struct {
	DurationUS  string `json:"duration_us"`
	Transitions string `json:"transitions"`
}

type Reporting struct {
	ETLSources      []ETLSource `json:"etl_sources"`
	IsWriter        bool        `json:"is_writer"`
	LastPublishTime string      `json:"last_publish_time"`
}

type ETLSource struct {
	Connected              bool   `json:"connected"`
	GRPCPort               string `json:"grpc_port"`
	IP                     string `json:"ip"`
	LastMessageArrivalTime string `json:"last_message_arrival_time"`
	ValidatedLedgersRange  string `json:"validated_ledgers_range"`
	WebsocketPort          string `json:"websocket_port"`
}

type Load struct {
	// TODO determine job types array format
	JobTypes []interface{} `json:"job_types"`
	Threads  uint          `json:"threads"`
}

type Close struct {
	ConvergeTimeS float32 `json:"converge_time_s"`
	Proposers     uint    `json:"proposers"`
}

type LedgerInfo struct {
	Age            uint          `json:"age"`
	BaseFeeXRP     float32       `json:"base_fee_xrp"`
	Hash           types.Hash256 `json:"hash"`
	ReserveBaseXRP float32       `json:"reserve_base_xrp"`
	ReserveIncXRP  float32       `json:"reserve_inc_xrp"`
	Seq            uint          `json:"seq"`
}
