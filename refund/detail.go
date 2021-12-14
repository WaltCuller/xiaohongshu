package refund

const (
	URL_REFUND_DETAIL    = "/ark/open_api/v0/refund/{returns_id}" // 售后详情路由
	METHOD_REFUND_DETAIL = "afterSale.getAfterSaleDetail"
)

// ReceiveAbnormalType 收货异常类型
type ReceiveAbnormalType uint

const (
	RECEIVE_1 ReceiveAbnormalType = iota + 1
	RECEIVE_2
	RECEIVE_4 = 2 * iota
	RECEIVE_5
	RECEIVE_8 = 2 * iota
	RECEIVE_9
)

// Uint ...
func (r ReceiveAbnormalType) Uint() uint {
	return uint(r)
}

// RefundStatus 退款状态，0：暂未触发退款；108 触发退款；1 退款中；2 退款成功；3 退款失败；101 转帐超时；401 已取消；201 待审核；301 审核通过；302 审核不通过；402 自动关闭
type RefundStatus uint

const (
	REFUND_0   RefundStatus = iota // 暂未触发退款
	REFUND_1                       // 退款中
	REFUND_2                       // 退款成功
	REFUND_3                       // 退款失败
	REFUND_101 RefundStatus = 101  // 转账超时
	REFUND_108 RefundStatus = 108  // 触发退款
	REFUND_201 RefundStatus = 201  // 待审核
	REFUND_301 RefundStatus = 301  // 审核通过
	REFUND_302 RefundStatus = 302  // 审核不通过
	REFUND_401 RefundStatus = 401  // 已取消
	REFUND_402 RefundStatus = 402  // 自动关闭
)

// Uint ...
func (r RefundStatus) Uint() uint {
	return uint(r)
}

// ReqDetail 售后详情请求
type ReqDetail struct {
	AfterSaleId string `json:"afterSaleId"` // 售后ID
}

// RspDetail 售后详情返回
type RspDetail struct {
	ReturnsId            string              `json:"returnsId" mapstructure:"returnsId"`                       // 售后id
	ReturnType           ReturnType          `json:"returnType" mapstructure:"returnType"`                     // 退货类型 1-退货退款, 2-换货, 3:仅退款(old) 4:仅退款(new) 理论上不会有3出现 -1 - 全部
	ReasonId             string              `json:"reasonId" mapstructure:"reasonId"`                         // 售后原因id
	Reason               string              `json:"reason" mapstructure:"reason"`                             // 售后原因说明
	Status               ReturnStatus        `json:"status" mapstructure:"status"`                             // 售后状态 1:待审核 2:待用户寄回 3:待收货 4:完成 5:取消 6:关闭 9:拒绝 9999:删除
	SubStatus            ReturnSubStats      `json:"subStatus" mapstructure:"subStatus"`                       // 售后子状态 301-待审核 302-快递已签收 304-收货异常
	ReceiveAbnormalType  ReceiveAbnormalType `json:"receiveAbnormalType" mapstructure:"receiveAbnormalType"`   // 收货异常类型
	PackageId            string              `json:"packageId" mapstructure:"packageId"`                       // 包裹id
	ExchangePackageId    string              `json:"exchangePackageId" mapstructure:"exchangePackageId"`       // 换货包裹id
	OrderId              string              `json:"orderId" mapstructure:"orderId"`                           // 订单id
	UserId               string              `json:"userId" mapstructure:"userId"`                             // 用户id
	CreatedAt            int                 `json:"createdAt" mapstructure:"createdAt"`                       // 创建时间
	ReturnExpressNo      string              `json:"returnExpressNo" mapstructure:"returnExpressNo"`           // 售后包裹快递单号
	ReturnExpressCompany string              `json:"returnExpressCompany" mapstructure:"returnExpressCompany"` // 售后包裹快递公司
	ReturnAddress        string              `json:"returnAddress" mapstructure:"returnAddress"`               // 售后寄回地址
	ShipNeeded           IsShipNeeded        `json:"shipNeeded" mapstructure:"shipNeeded"`                     // 是否需要寄回 0-否 1-是 -1-全部
	Refunded             bool                `json:"refunded" mapstructure:"refunded"`                         // 是否已退款
	RefundStatus         RefundStatus        `json:"refundStatus" mapstructure:"refundStatus"`                 // 退款状态，0：暂未触发退款；108 触发退款；1 退款中；2 退款成功；3 退款失败；101 转帐超时；401 已取消；201 待审核；301 审核通过；302 审核不通过；402 自动关闭
	AutoReceiveDeadline  int                 `json:"autoReceiveDeadline" mapstructure:"autoReceiveDeadline"`   // 超时自动确认收货的时间
	UseFastRefund        string              `json:"useFastRefund" mapstructure:"useFastRefund"`               // 是否急速退款
	ProofPhotos          []string            `json:"proofPhotos" mapstructure:"proofPhotos"`                   // 照片凭证
	Desc                 string              `json:"desc" mapstructure:"desc"`                                 // 描述
	Note                 string              `json:"note" mapstructure:"note"`                                 // 备注
	RefundTime           int                 `json:"refundTime" mapstructure:"refundTime"`                     // 退款时间
	FillExpressTime      int                 `json:"fillExpressTime" mapstructure:"fillExpressTime"`           // 填写退货快递单时间
	ExpressSignTime      int                 `json:"expressSignTime" mapstructure:"expressSignTime"`           // 退货快递签收时间
	Items                []struct {
		ItemId            string `json:"itemId" mapstructure:"itemId"`                       // 商品id
		ItemName          string `json:"itemName" mapstructure:"itemName"`                   // 商品名称
		Image             string `json:"image" mapstructure:"image"`                         // 商品主图
		Price             int    `json:"price" mapstructure:"price"`                         // 商品价格
		BoughtCount       int    `json:"boughtCount" mapstructure:"boughtCount"`             // 购买数量
		AppliedCount      int    `json:"appliedCount" mapstructure:"appliedCount"`           // 申请退货数量
		ReturnedCount     int    `json:"returnedCount" mapstructure:"returnedCount"`         // 实际收货数量
		RefundedCount     int    `json:"refundedCount" mapstructure:"refundedCount"`         // 退款数量
		ReturnPrice       int    `json:"returnPrice" mapstructure:"returnPrice"`             // 实际退款
		ExchangeItemId    string `json:"exchangeItemId" mapstructure:"exchangeItemId"`       // 换货商品id
		ExchangeItemName  string `json:"exchangeItemName" mapstructure:"exchangeItemName"`   // 换货商品名称
		ExchangeItemImage string `json:"exchangeItemImage" mapstructure:"exchangeItemImage"` // 换货商品主图
		Skucode           string `json:"skucode" mapstructure:"skucode"`                     // 商品代码
		Barcode           string `json:"barcode" mapstructure:"barcode"`                     // 商品条码
		ExchangeSkucode   string `json:"exchangeSkucode" mapstructure:"exchangeSkucode"`     // 换货商品代码
		ExchangeBarcode   string `json:"exchangeBarcode" mapstructure:"exchangeBarcode"`     // 换货商品条码
	} `json:"items" mapstructure:"items"` // 售后商品
	RefundFee                int    `json:"refundFee" mapstructure:"refundFee"`                               // 退款金额，单位元
	ReturnExpressRefundable  bool   `json:"returnExpressRefundable" mapstructure:"returnExpressRefundable"`   // 退货运费是否可退
	ReturnExpressRefunded    string `json:"returnExpressRefunded" mapstructure:"returnExpressRefunded"`       // 退货运费是否已退
	ExpectRefundFee          int    `json:"expectRefundFee" mapstructure:"expectRefundFee"`                   // 期望退款金额，单位元，目前是售后单包含的商品总金额
	UpdatedAt                int    `json:"updatedAt" mapstructure:"updatedAt"`                               // 更新时间
	ReturnExpressCompanyCode string `json:"returnExpressCompanyCode" mapstructure:"returnExpressCompanyCode"` // 退货快递公司编号
}
