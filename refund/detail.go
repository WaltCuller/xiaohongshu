package refund

const (
	URL_REFUND_DETAIL = "/ark/open_api/v0/refund/{returns_id}"
)

type ReceiveAbnormalType uint

const (
	RECEIVE_1 ReceiveAbnormalType = iota + 1
	RECEIVE_2
	RECEIVE_4 = 2 * iota
	RECEIVE_5
	RECEIVE_8 = 2 * iota
	RECEIVE_9
)

func (r ReceiveAbnormalType) Uint() uint {
	return uint(r)
}

type RefundStatus uint

const (
	REFUND_1 RefundStatus = iota + 1
	REFUND_2
	REFUND_108 RefundStatus = iota + 106
)

func (r RefundStatus) Uint() uint {
	return uint(r)
}

type Detail struct {
}

type DetailRsp struct {
	ReturnsId            string         `json:"returns_id" mapstructure:"returns_id"`
	ReturnType           ReturnType     `json:"return_type" mapstructure:"return_type"`
	ReasonId             int            `json:"reason_id" mapstructure:"reason_id"`
	Reason               string         `json:"reason" mapstructure:"reason"`
	Status               ReturnStatus   `json:"status" mapstructure:"status"`
	SubStatus            ReturnSubStats `json:"sub_status" mapstructure:"sub_status"`
	PackageId            string         `json:"package_id" mapstructure:"package_id"`
	ExchangePackageId    string         `json:"exchange_package_id" mapstructure:"exchange_package_id"`
	OrderId              string         `json:"order_id" mapstructure:"order_id"`
	UserId               string         `json:"user_id" mapstructure:"user_id"`
	CreatedAt            int64          `json:"created_at" mapstructure:"created_at"`
	UpdatedAt            int64          `json:"updated_at" mapstructure:"updated_at"`
	ReturnExpressNo      string         `json:"return_express_no" mapstructure:"return_express_no"`
	ReturnExpressCompany string         `json:"return_express_company" mapstructure:"return_express_company"`
	ReturnAddress        string         `json:"return_address" mapstructure:"return_address"`
	Items                []struct {
		ItemId            string  `json:"item_id" mapstructure:"item_id"`
		ItemName          string  `json:"item_name" mapstructure:"item_name"`
		Skucode           string  `json:"skucode" mapstructure:"skucode"`
		Image             string  `json:"image" mapstructure:"image"`
		Price             float64 `json:"price" mapstructure:"price"`
		BoughtCount       int     `json:"bought_count" mapstructure:"bought_count"`
		AppliedCount      int     `json:"applied_count" mapstructure:"applied_count"`
		ReturnedCount     int     `json:"returned_count" mapstructure:"returned_count"`
		ReturnPrice       float64 `json:"return_price" mapstructure:"return_price"`
		ExchangeItemId    string  `json:"exchange_item_id" mapstructure:"exchange_item_id"`
		ExchangeItemName  string  `json:"exchange_item_name" mapstructure:"exchange_item_name"`
		ExchangeItemImage string  `json:"exchange_item_image" mapstructure:"exchange_item_image"`
		ExchangeSkucode   string  `json:"exchange_skucode" mapstructure:"exchange_skucode"`
		Barcode           string  `json:"barcode" mapstructure:"barcode"`
		ExchangeBarcode   string  `json:"exchange_barcode" mapstructure:"exchange_barcode"`
	} `json:"items" mapstructure:"items"`
	ProofPhotos             []string            `json:"proof_photos" mapstructure:"proof_photos"`
	Desc                    string              `json:"desc" mapstructure:"desc"`
	Note                    string              `json:"note" mapstructure:"note"`
	ShipNeeded              IsShipNeeded        `json:"ship_needed" mapstructure:"ship_needed"`
	Refunded                bool                `json:"refunded" mapstructure:"refunded"`
	RefundStatus            int                 `json:"refund_status" mapstructure:"refund_status"`
	RefundFee               float64             `json:"refund_fee" mapstructure:"refund_fee"`
	RefundTime              int64               `json:"refund_time" mapstructure:"refund_time"`
	FillExpressTime         int64               `json:"fill_express_time" mapstructure:"fill_express_time"`
	ExpressSignTime         int64               `json:"express_sign_time" mapstructure:"express_sign_time"`
	AutoReceiveDeadline     int64               `json:"auto_receive_deadline" mapstructure:"auto_receive_deadline"`
	ReceiveAbnormalType     ReceiveAbnormalType `json:"receive_abnormal_type" mapstructure:"receive_abnormal_type"`
	ReturnExpressRefundable bool                `json:"return_express_refundable" mapstructure:"return_express_refundable"`
	ReturnExpressRefunded   bool                `json:"return_express_refunded" mapstructure:"return_express_refunded"`
	UseFastRefund           bool                `json:"use_fast_refund" mapstructure:"use_fast_refund"`
	ExpectRefundFee         float64             `json:"expect_refund_fee" mapstructure:"expect_refund_fee"`
}
