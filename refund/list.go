package refund

const (
	URL_REFUND_LIST = "/ark/open_api/v0/refund/list"
)

type List struct {
	Status     ReturnStatus `json:"status" structs:",omitempty"`
	ReturnType ReturnType   `json:"return_type" structs:",omitempty"`
	ReasonID   int          `json:"reason_id" structs:",omitempty"`
	StartTime  int64        `json:"start_time" structs:",omitempty"`
	EndTime    int64        `json:"end_time" structs:",omitempty"`
	Page       int          `json:"page" structs:",omitempty"`
	PageSize   int          `json:"page_size" structs:",omitempty"`
}

type ListRsp struct {
	TotalNumber int `json:"total_number" mapstructure:"total_number"`
	CurrentPage int `json:"current_page" mapstructure:"current_page"`
	TotalPage   int `json:"total_page" mapstructure:"total_page"`
	PageSize    int `json:"page_size" mapstructure:"page_size"`
	RefundList  []struct {
		ReturnsId            string         `json:"returns_id" mapstructure:"returns_id"`
		ReturnType           ReturnType     `json:"return_type" mapstructure:"return_type"`
		ReasonId             int            `json:"reason_id" mapstructure:"reason_id"`
		Reason               string         `json:"reason" mapstructure:"reason"`
		Status               ReturnStatus   `json:"status" mapstructure:"status"`
		SubStatus            ReturnSubStats `json:"sub_status" mapstructure:"sub_status"`
		ReceiveAbnormalType  int            `json:"receive_abnormal_type" mapstructure:"receive_abnormal_type"`
		PackageId            string         `json:"package_id" mapstructure:"package_id"`
		ExchangePackageId    string         `json:"exchange_package_id" mapstructure:"exchange_package_id"`
		OrderId              string         `json:"order_id" mapstructure:"order_id"`
		UserId               string         `json:"user_id" mapstructure:"user_id"`
		CreatedAt            int64          `json:"created_at" mapstructure:"created_at"`
		UpdatedAt            int64          `json:"updated_at" mapstructure:"updated_at"`
		ReturnExpressNo      string         `json:"return_express_no" mapstructure:"return_express_no"`
		ReturnExpressCompany string         `json:"return_express_company" mapstructure:"return_express_company"`
		ReturnAddress        string         `json:"return_address" mapstructure:"return_address"`
		ShipNeeded           IsShipNeeded   `json:"ship_needed" mapstructure:"ship_needed"`
		Refunded             bool           `json:"refunded" mapstructure:"refunded"`
		AutoReceiveDeadline  int64          `json:"auto_receive_deadline" mapstructure:"auto_receive_deadline"`
		UseFastRefund        bool           `json:"use_fast_refund" mapstructure:"use_fast_refund"`
	} `json:"refund_list" mapstructure:"refund_list"`
}
