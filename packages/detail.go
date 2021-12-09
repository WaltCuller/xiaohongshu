package packages

const (
	URL_PACKAGES_DETAIL = "/ark/open_api/v0/packages/{package_id}"
)

type Detail struct {
}

type DetailRsp struct {
	PackageId   string        `json:"package_id" mapstructure:"package_id"`
	OrderId     string        `json:"order_id" mapstructure:"order_id"`
	Logistics   LogisticsMode `json:"logistics" mapstructure:"logistics"`
	PackageType string        `json:"package_type" mapstructure:"package_type"`
	PresaleInfo struct {
		ShipStartTime int `json:"ship_start_time" mapstructure:"ship_start_time"`
		ShipEndTime   int `json:"ship_end_time" mapstructure:"ship_end_time"`
	} `json:"presale_info" mapstructure:"presale_info"`
	Time                   int         `json:"time" mapstructure:"time"`
	PayTime                int         `json:"pay_time" mapstructure:"pay_time"`
	ConfirmTime            int         `json:"confirm_time" mapstructure:"confirm_time"`
	CreateTime             int         `json:"create_time:" mapstructure:"create_time"`
	ExpressCompanyCode     string      `json:"express_company_code" mapstructure:"express_company_code"`
	ExpressCompanyName     string      `json:"express_company_name" mapstructure:"express_company_name"`
	ExpressNo              string      `json:"express_no" mapstructure:"express_no"`
	Status                 OrderStatus `json:"status" mapstructure:"status"`
	ReceiverName           string      `json:"receiver_name" mapstructure:"receiver_name"`
	ReceiverPhone          string      `json:"receiver_phone" mapstructure:"receiver_phone"`
	ReceiverAddress        string      `json:"receiver_address" mapstructure:"receiver_address"`
	Province               string      `json:"province" mapstructure:"province"`
	City                   string      `json:"city" mapstructure:"city"`
	District               string      `json:"district" mapstructure:"district"`
	TotalNetWeight         int         `json:"total_net_weight" mapstructure:"total_net_weight"`
	PayAmount              int         `json:"pay_amount" mapstructure:"pay_amount"`
	IdNumber               string      `json:"id_number" mapstructure:"id_number"`
	BuyerName              string      `json:"buyer_name" mapstructure:"buyer_name"`
	InternationalExpressNo string      `json:"international_express_no" mapstructure:"international_express_no"`
	DeliveryTimePreference string      `json:"delivery_time_preference" mapstructure:"delivery_time_preference"`
	OrderDeclaredAmount    string      `json:"order_declared_amount" mapstructure:"order_declared_amount"`
	PaintMarker            string      `json:"paint_marker" mapstructure:"paint_marker"`
	ExpressExtend1         string      `json:"express_extend_1" mapstructure:"express_extend_1"`
	ExpressExtend2         string      `json:"express_extend_2" mapstructure:"express_extend_2"`
	ShippingFee            int         `json:"shipping_fee" mapstructure:"shipping_fee"`
	UserNote               string      `json:"user_note" mapstructure:"user_note"`
	SellerNote             string      `json:"seller_note" mapstructure:"seller_note"`
	DeclarationCode        string      `json:"declaration_code" mapstructure:"declaration_code"`
	FrontUrl               string      `json:"front_url" mapstructure:"front_url"`
	BackUrl                string      `json:"back_url" mapstructure:"back_url"`
	ItemList               []struct {
		MerchantDiscount string `json:"merchant_discount" mapstructure:"merchant_discount"`
		RedDiscount      string `json:"red_discount" mapstructure:"red_discount"`
		Barcode          string `json:"barcode" mapstructure:"barcode"`
		Skucode          string `json:"skucode" mapstructure:"skucode"`
		ItemName         string `json:"item_name" mapstructure:"item_name"`
		Type             string `json:"type" mapstructure:"type"`
		Qty              int    `json:"qty" mapstructure:"qty"`
		Price            string `json:"price" mapstructure:"price"`
		PayPrice         string `json:"pay_price" mapstructure:"pay_price"`
		NetWeight        int    `json:"net_weight" mapstructure:"net_weight"`
		RegisterName     string `json:"register_name" mapstructure:"register_name"`
		ArticleNo        string `json:"article_no" mapstructure:"article_no"`
	} `json:"item_list" mapstructure:"item_list"`
}
