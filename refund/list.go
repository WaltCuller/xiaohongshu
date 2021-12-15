package refund

const (
	URL_REFUND_LIST    = "/ark/open_api/v0/refund/list" // 售后列表路由
	METHOD_REFUND_LIST = "afterSale.listAfterSaleApi"
)

// TimeType startTime/endTime对应的时间类型
type TimeType int

const (
	CREATED TimeType = iota + 1 // 创建时间
	UPDATED                     // 更新时间
)

// ReqList 售后列表请求
type ReqList struct {
	Status     ReturnStatus `json:"status" param:"status"`         // 售后状态 1：待审核；2:待用户寄回；3:待收货；4:完成；5:取消；6:关闭；9:拒绝；9999:删除，不传默认全部
	PageNo     int          `json:"pageNo" param:"pageNo"`         // 返回页码 默认 1，页码从 1 开始 PS：当前采用分页返回，数量和页数会一起传，如果不传，则采用 默认值
	PageSize   int          `json:"pageSize" param:"pageSize"`     // 返回数量，默认50最大100
	StartTime  int64        `json:"startTime" param:"startTime"`   // 查询时间起点
	EndTime    int64        `json:"endTime" param:"endTime"`       // 查询时间终点
	TimeType   TimeType     `json:"timeType" param:"timeType"`     // 时间类型，1：根据创建时间查询 end-start<=24h；2：根据更新时间查询 end-start<=30min
	UseHasNext bool         `json:"useHasNext" param:"useHasNext"` // 是否返回所有数据,true 不返会total 返回 hasNext = true 表示仍有数据，false 返回total
	ReasonId   string       `json:"reasonId" param:"reasonId"`     // 编号
	ReturnType ReturnType   `json:"returnType" param:"returnType"` // 售后类型 不传/0:全部；1:退货退款；2:换货；3:仅退款(old) 4:仅退款(new) 理论上不会有3出现
}

// RspList 售后列表返回
type RspList struct {
	Total               int  `json:"total" mapstructure:"total"`       // 查询到的总数，useHasNext=true时为0
	PageNo              int  `json:"pageNo" mapstructure:"pageNo"`     // 当前页数
	PageSize            int  `json:"pageSize" mapstructure:"pageSize"` // 页大小
	HaxNext             bool `json:"haxNext" mapstructure:"haxNext"`   // 是否有下一页
	SimpleAfterSaleList []struct {
		ReturnsId                string              `json:"returnsId" mapstructure:"returnsId"`                               // 售后ID
		ReturnType               int                 `json:"returnType" mapstructure:"returnType"`                             // 售后类型
		ReasonId                 int                 `json:"reasonId" mapstructure:"reasonId"`                                 // 售后原因ID
		Reason                   string              `json:"reason" mapstructure:"reason"`                                     // 售后原因
		Status                   ReturnStatus        `json:"status" mapstructure:"status"`                                     // 售后状态 1:待审核 2:待用户寄回 3:待收货 4:完成 5:取消 6:关闭 9:拒绝 9999:删除
		SubStatus                ReturnSubStats      `json:"subStatus" mapstructure:"subStatus"`                               // 售后子状态 301-待审核 302-快递已签收 304-收货异常
		ReceiveAbnormalType      ReceiveAbnormalType `json:"receiveAbnormalType" mapstructure:"receiveAbnormalType"`           // 收货异常类型
		PackageId                string              `json:"packageId" mapstructure:"packageId"`                               // 包裹ID
		ExchangePackageId        string              `json:"exchangePackageId" mapstructure:"exchangePackageId"`               // 换货包裹ID
		OrderId                  string              `json:"orderId" mapstructure:"orderId"`                                   // 订单ID
		UserId                   string              `json:"userId" mapstructure:"userId"`                                     // 用户ID
		CreatedTime              int64               `json:"createdTime" mapstructure:"createdTime"`                           // 售后创建时间戳（毫秒）
		ReturnExpressNo          string              `json:"returnExpressNo" mapstructure:"returnExpressNo"`                   // 售后快递单号
		ReturnExpressCompany     string              `json:"returnExpressCompany" mapstructure:"returnExpressCompany"`         // 售后快递公司
		ReturnAddress            string              `json:"returnAddress" mapstructure:"returnAddress"`                       // 售后退货地址
		ShipNeeded               IsShipNeeded        `json:"shipNeeded" mapstructure:"shipNeeded"`                             // 是否需要寄回 1-需要 0-不需要
		Refunded                 bool                `json:"refunded" mapstructure:"refunded"`                                 // 是否已退款
		RefundStatus             RefundStatus        `json:"refundStatus" mapstructure:"refundStatus"`                         // 退款状态 108触发退款 1退款中 3退款失败 2退款成功 401已取消 101已创建 201待审核 301审核通过 302审核不通过 402自动关闭
		AutoReceiveDeadline      int64               `json:"autoReceiveDeadline" mapstructure:"autoReceiveDeadline"`           // 自动确认收货时间
		UseFastRefund            bool                `json:"useFastRefund" mapstructure:"useFastRefund"`                       // 是否急速退款
		UpdateTime               int64               `json:"updateTime" mapstructure:"updateTime"`                             // 售后更新时间戳（毫秒）
		ReturnExpressCompanyCode string              `json:"returnExpressCompanyCode" mapstructure:"returnExpressCompanyCode"` // 退货快递公司编号
	} `json:"simpleAfterSaleList" mapstructure:"simpleAfterSaleList"` // 售后信息列表
	MaxPageNo int `json:"maxPageNo" mapstructure:"maxPageNo"` // 最大页码数
}
