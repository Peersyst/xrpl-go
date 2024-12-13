package types

type JobType struct {
	JobType    string `json:"job_type"`
	PerSecond  int    `json:"per_second"`
	PeakTime   int    `json:"peak_time,omitempty"`
	AvgTime    int    `json:"avg_time,omitempty"`
	InProgress int    `json:"in_progress,omitempty"`
}
