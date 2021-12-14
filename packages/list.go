package packages

const (
	URL_PACKAGES_LIST    = "/ark/open_api/v0/packages" // 订单列表路由
	METHOD_PACKAGES_LIST = "package.getPackageList"
)

// TimeType startTime/endTime对应的时间类型
type TimeType int

const (
	CREATED TimeType = iota + 1 // 创建时间
	UPDATED                     // 更新时间
)

// Int ...
func (t TimeType) Int() int {
	return int(t)
}

// PackageType 包裹类型
type PackageType int

const (
	NULL         PackageType = iota // 全部
	NORMAL                          // 普通
	PREORDER                        // 定金预售
	FULLPREORDER                    // 全额预售
	DELAY                           // 延迟发货
	CHANGE                          // 换货补发
)

// Int ...
func (p PackageType) Int() int {
	return int(p)
}

// PackageStatus 包裹状态
type PackageStatus int

const (
	ALL PackageStatus = iota // 全部
	_1                       // 已下单代付款
	_2                       // 已支付处理中
	_3                       // 清关中
	_4                       // 待发货
	_5                       // 部分发货
	_6                       // 待收货
	_7                       // 已完成
	_8                       // 已关闭
	_9                       // 已取消
	_10                      // 换货申请中
)

// Int ...
func (p PackageStatus) Int() int {
	return int(p)
}

// ReqList 订单列表参数
type ReqList struct {
	StartTime     int           `json:"startTime" param:"startTime"`         // 时间范围起点
	EndTime       int           `json:"endTime" param:"endTime"`             // 时间范围终点
	TimeType      TimeType      `json:"timeType" param:"timeType"`           // startTime/endTime对应的时间类型，1 创建时间 限制 end-start<=24h、2 更新时间 限制 end-start<=30min 倒序拉取 最后一页到第一页
	PackageType   PackageType   `json:"packageType" param:"packageType"`     // 包裹类型，0/null 全部 1 普通 normal 2 定金预售 3 全款预售 4 延迟发货 5 换货补发
	PackageStatus PackageStatus `json:"packageStatus" param:"packageStatus"` // 包裹状态，0全部 1已下单待付款 2已支付处理中 3清关中 4待发货 5部分发货 6待收货 7已完成 8已关闭 9已取消 10换货申请中
	PageNo        int           `json:"pageNo" param:"pageNo"`               // 页码，默认1，限制100
	PageSize      int           `json:"pageSize" param:"pageSize"`           // 查询总数，默认50，限制100
}

// RspList 订单列表返回
type RspList struct {
	Total       int `json:"total" mapstructure:"total"`         // 包裹总数
	PageNo      int `json:"pageNo" mapstructure:"pageNo"`       // 当前页数 请求参数中的值
	PageSize    int `json:"pageSize" mapstructure:"pageSize"`   // 页大小 请求参数中的值
	MaxPageNo   int `json:"maxPageNo" mapstructure:"maxPageNo"` // 最大页码数 方便 直接从最后一页拉数据
	PackageList []struct {
		PackageId               string                  `json:"packageId" mapstructure:"packageId"`                                // 包裹号
		OrderId                 string                  `json:"orderId" mapstructure:"orderId"`                                    // 母单号
		PackageType             PackageType             `json:"packageType" mapstructure:"package_type"`                           // 包裹类型，1普通 2定金预售 3全款预售 4延迟发货 5换货补发
		PackageStatus           PackageStatus           `json:"packageStatus" mapstructure:"package_status"`                       // 包裹状态，1已下单待付款 2已支付处理中 3清关中 4待发货 5部分发货 6待收货 7已完成 8已关闭 9已取消 10换货申请中
		PackageAfterSalesStatus PackageAfterSalesStatus `json:"packageAfterSalesStatus" mapstructure:"package_after_sales_status"` // 售后状态，1无售后 2售后处理中 3售后完成(含取消)
		CancelStatus            CancelStatus            `json:"cancelStatus" mapstructure:"cancel_status"`                         // 申请取消状态，0未申请取消 1取消处理中
		CreatedTime             int64                   `json:"createdTime" mapstructure:"created_time"`                           // 创建时间 单位ms
		PaidTime                int64                   `json:"paidTime" mapstructure:"paid_time"`                                 // 支付时间 单位ms
		UpdateTime              int64                   `json:"updateTime" mapstructure:"update_time"`                             // 更新时间 单位ms
		DeliveryTime            int64                   `json:"deliveryTime" mapstructure:"delivery_time"`                         // 包裹发货时间 单位ms
		CancelTime              int64                   `json:"cancelTime" mapstructure:"cancel_time"`                             // 包裹取消时间 单位ms
		FinishTime              int64                   `json:"finishTime" mapstructure:"finish_time"`                             // 包裹完成时间 单位ms
		PromiseLastDeliveryTime int64                   `json:"promiseLastDeliveryTime" mapstructure:"promise_last_delivery_time"` // 承诺最晚发货时间 单位ms
		PlanInfoId              string                  `json:"planInfoId" mapstructure:"plan_info_id"`                            // 物流方案id
		PlanInfoName            string                  `json:"planInfoName" mapstructure:"plan_info_name"`                        // 物流方案名称
		ReceiverCountryId       string                  `json:"receiverCountryId" mapstructure:"receiver_country_id"`              // 收件人国家id
		ReceiverCountryName     string                  `json:"receiverCountryName" mapstructure:"receiver_country_name"`          // 目前仅 中国
		ReceiverProvinceId      string                  `json:"receiverProvinceId" mapstructure:"receiver_province_id"`            // 收件人省份id
		ReceiverProvinceName    string                  `json:"receiverProvinceName" mapstructure:"receiver_province_name"`        // 收件人省份
		ReceiverCityId          string                  `json:"receiverCityId" mapstructure:"receiver_city_id"`                    // 收件人城市id
		ReceiverCityName        string                  `json:"receiverCityName" mapstructure:"receiver_city_name"`                // 收件人城市
		ReceiverDistrictId      string                  `json:"receiverDistrictId" mapstructure:"receiver_district_id"`            // 收件人区县id
		ReceiverDistrictName    string                  `json:"receiverDistrictName" mapstructure:"receiver_district_name"`        // 收件人区县名称
		CustomerRemark          string                  `json:"customerRemark" mapstructure:"customer_remark"`                     // 用户备注
		SellerRemark            string                  `json:"sellerRemark" mapstructure:"seller_remark"`                         // 商家标记备注
		SellerRemarkFlag        SellerRemarkFlag        `json:"sellerRemarkFlag" mapstructure:"seller_remark_flag"`                // 商家标记优先级，ark订单列表展示旗子颜色 1灰旗 2红旗 3黄旗 4绿旗 5蓝旗 6紫旗
		OriginalPackageId       string                  `json:"originalPackageId" mapstructure:"original_package_id"`              // 原始包裹编号，换货包裹的原包裹
		Logistics               LogisticsMode           `json:"logistics" mapstructure:"logistics"`                                // 物流模式: red_express三方备货直邮(备货海外仓),red_domestic_trade(三方备货内贸),red_standard(三方备货保税仓),red_auto(三方自主发货),red_box(三方小包),red_bonded(三方保税)
	} `json:"packageList" mapstructure:"package_list"` // 包裹信息
}
