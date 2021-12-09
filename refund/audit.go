package refund

type Audit struct {
	ReturnsId          string        `json:"returns_id,omitempty"`
	OperateTime        int64         `json:"operate_time,omitempty"`
	AuditResult        int           `json:"audit_result,omitempty"`
	AuditDescription   string        `json:"audit_description,omitempty"`
	RejectReason       int           `json:"reject_reason,omitempty"`
	AutoRefundDisabled bool          `json:"auto_refund_disabled,omitempty"`
	ShipNeeded         int           `json:"ship_needed,omitempty"`
	ReceiverInfo       *ReceiverInfo `json:"receiver_info,omitempty"`
}

type ReceiverInfo struct {
	Code     string `json:"code,omitempty"`
	Country  string `json:"country,omitempty"`
	Province string `json:"province,omitempty"`
	City     string `json:"city,omitempty"`
	District string `json:"district,omitempty"`
	Street   string `json:"street,omitempty"`
}
