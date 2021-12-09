package packages

type List struct {
	StartTime int64  `json:"start_time" structs:",omitempty"`
	EndTime   int64  `json:"end_time" structs:",omitempty"`
	TimeType  string `json:"time_type" structs:",omitempty"`
	Logistics string `json:"logistics" structs:",omitempty"`
	Status    string `json:"status" structs:",omitempty"`
	PageNo    int64  `json:"page_no" structs:",omitempty"`
	PageSize  int64  `json:"page_size" structs:",omitempty"`
}
