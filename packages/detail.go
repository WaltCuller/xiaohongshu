package packages

const (
	URL_PACKAGES_DETAIL    = "/ark/open_api/v0/packages/{package_id}" // 订单详情路由
	METHOD_PACKAGES_DETAIL = "package.getPackageDetail"
)

// PackageAfterSalesStatus 售后状态
type PackageAfterSalesStatus int

const (
	NONE     PackageAfterSalesStatus = iota + 1 // 无售后
	HANDLING                                    // 售后处理中
	COMPLETE                                    // 售后完成（含取消）
)

// Int ...
func (p PackageAfterSalesStatus) Int() int {
	return int(p)
}

// CancelStatus 申请取消状态
type CancelStatus int

const (
	CANCEL_0 CancelStatus = iota // 未申请取消
	CANCEL_1                     // 取消处理中
)

// Int ...
func (c CancelStatus) Int() int {
	return int(c)
}

// SellerRemarkFlag 商家标记优先级
type SellerRemarkFlag int

const (
	GRAY   SellerRemarkFlag = iota + 1 // 灰
	RED                                // 红
	YELLOW                             // 黄
	GREEN                              // 绿
	BLUE                               // 蓝
	P                                  // 紫
)

// Int ...
func (s SellerRemarkFlag) Int() int {
	return int(s)
}

// IsGift 是否是赠品
type IsGift bool

const (
	GIFT_TRUE  IsGift = true  // 赠品
	GIFT_FALSE        = false // 非赠品
)

// Bool ...
func (i IsGift) Bool() bool {
	return bool(i)
}

// IsChannel 是否是渠道商品
type IsChannel bool

const (
	CHANNEL_TRUE  IsChannel = true  // 渠道商品
	CHANNEL_FALSE           = false // 非渠道商品
)

// Bool ...
func (i IsChannel) Bool() bool {
	return bool(i)
}

// ReqDetail 订单详情请求
type ReqDetail struct {
	PackageId string `json:"packageId" mapstructure:"packageId"` // 包裹号
}

// RspDetail 订单详情返回
type RspDetail struct {
	PackageId                string                  `json:"packageId" mapstructure:"packageId"`                               // 包裹号
	OrderId                  string                  `json:"orderId" mapstructure:"orderId"`                                   // 母单号
	PackageType              PackageType             `json:"packageType" mapstructure:"packageType"`                           // 包裹类型，1普通 2定金预售 3全款预售 4延迟发货 5换货补发
	PackageStatus            PackageStatus           `json:"packageStatus" mapstructure:"packageStatus"`                       // 包裹状态，1已下单待付款 2已支付处理中 3清关中 4待发货 5部分发货 6待收货 7已完成 8已关闭 9已取消 10换货申请中
	PackageAfterSalesStatus  PackageAfterSalesStatus `json:"packageAfterSalesStatus" mapstructure:"packageAfterSalesStatus"`   // 售后状态，1无售后 2售后处理中 3售后完成(含取消)
	CancelStatus             CancelStatus            `json:"cancelStatus" mapstructure:"cancelStatus"`                         // 申请取消状态，0未申请取消 1取消处理中
	CreatedTime              int64                   `json:"createdTime" mapstructure:"createdTime"`                           // 创建时间 单位ms
	PaidTime                 int64                   `json:"paidTime" mapstructure:"paidTime"`                                 // 支付时间 单位ms
	UpdateTime               int64                   `json:"updateTime" mapstructure:"updateTime"`                             // 更新时间 单位ms
	DeliveryTime             int64                   `json:"deliveryTime" mapstructure:"deliveryTime"`                         // 包裹发货时间 单位ms
	CancelTime               int64                   `json:"cancelTime" mapstructure:"cancelTime"`                             // 包裹取消时间 单位ms
	FinishTime               int64                   `json:"finishTime" mapstructure:"finishTime"`                             // 包裹完成时间 单位ms
	PromiseLastDeliveryTime  int64                   `json:"promiseLastDeliveryTime" mapstructure:"promiseLastDeliveryTime"`   // 承诺最晚发货时间 单位ms
	PlanInfoId               string                  `json:"planInfoId" mapstructure:"planInfoId"`                             // 物流方案id
	PlanInfoName             string                  `json:"planInfoName" mapstructure:"planInfoName"`                         // 物流方案名称
	ReceiverCountryId        string                  `json:"receiverCountryId" mapstructure:"receiverCountryId"`               // 收件人国家id
	ReceiverCountryName      string                  `json:"receiverCountryName" mapstructure:"receiverCountryName"`           // 目前仅 中国
	ReceiverProvinceId       string                  `json:"receiverProvinceId" mapstructure:"receiverProvinceId"`             // 收件人省份id
	ReceiverProvinceName     string                  `json:"receiverProvinceName" mapstructure:"receiverProvinceName"`         // 收件人省份
	ReceiverCityId           string                  `json:"receiverCityId" mapstructure:"receiverCityId"`                     // 收件人城市id
	ReceiverCityName         string                  `json:"receiverCityName" mapstructure:"receiverCityName"`                 // 收件人城市
	ReceiverDistrictId       string                  `json:"receiverDistrictId" mapstructure:"receiverDistrictId"`             // 收件人区县id
	ReceiverDistrictName     string                  `json:"receiverDistrictName" mapstructure:"receiverDistrictName"`         // 收件人区县名称
	CustomerRemark           string                  `json:"customerRemark" mapstructure:"customerRemark"`                     // 用户备注
	SellerRemark             string                  `json:"sellerRemark" mapstructure:"sellerRemark"`                         // 商家标记备注
	SellerRemarkFlag         SellerRemarkFlag        `json:"sellerRemarkFlag" mapstructure:"sellerRemarkFlag"`                 // 商家标记优先级，ark订单列表展示旗子颜色 1灰旗 2红旗 3黄旗 4绿旗 5蓝旗 6紫旗
	PresaleDeliveryStartTime int64                   `json:"presaleDeliveryStartTime" mapstructure:"presaleDeliveryStartTime"` // 预售最早发货时间 单位ms
	PresaleDeliveryEndTime   int64                   `json:"presaleDeliveryEndTime" mapstructure:"presaleDeliveryEndTime"`     // 预售最晚发货时间 单位ms
	ItemList                 []struct {
		ItemId       string `json:"itemId" mapstructure:"itemId"`             // 商品id
		ItemName     string `json:"itemName" mapstructure:"itemName"`         // 商品名称
		Erpcode      string `json:"erpcode" mapstructure:"erpcode"`           // 商家编码(若为组合品，暂不支持组合品的商家编码，但skulist会返回子商品商家编码)
		ItemSpec     string `json:"itemSpec" mapstructure:"itemSpec"`         // 规格
		ItemImage    string `json:"itemImage" mapstructure:"itemImage"`       // 商品图片url
		ItemQuantity int    `json:"itemQuantity" mapstructure:"itemQuantity"` // 商品数量
		SkuList      []struct {
			ItemId                 string `json:"itemId" mapstructure:"itemId"`                                 // 单品商品Id(渠道商品为生成渠道商品的原商品单品id，组合商品为各个子商品的单品id，多包组为对应单包组商品id,商家编码同理)
			ErpCode                string `json:"erpCode" mapstructure:"erpCode"`                               // 商家编码
			Barcode                string `json:"barcode" mapstructure:"barcode"`                               // 商品条码
			SkuCode                string `json:"skuCode" mapstructure:"skuCode"`                               // 商品编码
			Quantity               int    `json:"quantity" mapstructure:"quantity"`                             // 购买数量
			RegisterName           string `json:"registerName" mapstructure:"registerName"`                     // 商品备案名称
			ItemName               string `json:"itemName" mapstructure:"itemName"`                             // 商品名
			PricePerSku            int    `json:"pricePerSku" mapstructure:"pricePerSku"`                       // 单个sku价格
			TaxPerSku              int    `json:"taxPerSku" mapstructure:"taxPerSku"`                           // 单个sku税金
			PaidAmountPerSku       int    `json:"paidAmountPerSku" mapstructure:"paidAmountPerSku"`             // 单个sku实付
			DepositAmountPerSku    int    `json:"depositAmountPerSku" mapstructure:"depositAmountPerSku"`       // 单个sku定金
			MerchantDiscountPerSku int    `json:"merchantDiscountPerSku" mapstructure:"merchantDiscountPerSku"` // 单个sku商家承担优惠
			RedDiscountPerSku      int    `json:"redDiscountPerSku" mapstructure:"redDiscountPerSku"`           // 单个sku平台承担优惠
			RawPricePerSku         int    `json:"rawPricePerSku" mapstructure:"rawPricePerSku"`                 // 单个sku原价
		} `json:"skuList" mapstructure:"skuList"` // 商品sku信息列表（可能出现相同skuCode不同价格的状况）
		TotalPaidAmount       int       `json:"totalPaidAmount" mapstructure:"totalPaidAmount"`             // 总支付金额（考虑总件数）商品总实付
		TotalMerchantDiscount int       `json:"totalMerchantDiscount" mapstructure:"totalMerchantDiscount"` // 商家承担总优惠
		TotalRedDiscount      int       `json:"totalRedDiscount" mapstructure:"totalRedDiscount"`           // 平台承担总优惠
		TotalTaxAmount        int       `json:"totalTaxAmount" mapstructure:"totalTaxAmount"`               // 商品税金
		TotalNetWeight        float32   `json:"totalNetWeight" mapstructure:"totalNetWeight"`               // 商品总净重
		ItemTag               IsGift    `json:"itemTag" mapstructure:"itemTag"`                             // 是否赠品，1 赠品 0 普通商品
		IsChannel             IsChannel `json:"isChannel" mapstructure:"isChannel"`                         // 是否是渠道商品
		Channel               string    `json:"channel" mapstructure:"channel"`                             // -
	} `json:"itemList" mapstructure:"itemList"` // item列表 相同itemid聚合 金额为相同item下sku价格总和 单位 分
	OriginalPackageId    string `json:"originalPackageId" mapstructure:"originalPackageId"`       // 原始关联包裹号(退换包裹的原包裹)
	TotalNetWeightAmount int    `json:"totalNetWeightAmount" mapstructure:"totalNetWeightAmount"` // 订单商品总净重 单位g
	TotalPayAmount       int    `json:"totalPayAmount" mapstructure:"totalPayAmount"`             // 订单实付金额(包含运费) 单位分
	TotalShippingFree    int    `json:"totalShippingFree" mapstructure:"totalShippingFree"`       // 包裹运费 单位分
	Unpack               bool   `json:"unpack" mapstructure:"unpack"`                             // 是否拆包 true已拆包 false未拆包
	ExpressTrackingNo    string `json:"expressTrackingNo" mapstructure:"expressTrackingNo"`       // 快递单号
	ExpressCompanyCode   string `json:"expressCompanyCode" mapstructure:"expressCompanyCode"`     // 快递公司编码
	ReceiverName         string `json:"receiverName" mapstructure:"receiverName"`                 // 收件人姓名 暂不返回 详情通过getReceiverInfo获取
	ReceiverPhone        string `json:"receiverPhone" mapstructure:"receiverPhone"`               // 收件人手机 暂不返回 详情通过getReceiverInfo获取
	ReceiverAddress      string `json:"receiverAddress" mapstructure:"receiverAddress"`           // 收件人地址 暂不返回 详情通过getReceiverInfo获取
	BoundExtendInfo      struct {
		PayNo          string   `json:"payNo" mapstructure:"payNo"`                   // 交易流水号
		PayChannel     string   `json:"payChannel" mapstructure:"payChannel"`         // 交易渠道
		ProductValue   string   `json:"productValue" mapstructure:"productValue"`     // 订单价值（货值，订单商品申价之和（税前价））
		PayAmount      string   `json:"payAmount" mapstructure:"payAmount"`           // 订单支付金额（含运费）
		TaxAmount      string   `json:"taxAmount" mapstructure:"taxAmount"`           // 订单税金
		ShippingFee    string   `json:"shippingFee" mapstructure:"shippingFee"`       // 运费 含运费税
		DiscountAmount string   `json:"discountAmount" mapstructure:"discountAmount"` // 订单优惠
		ZoneCodes      []string `json:"zoneCodes" mapstructure:"zoneCodes"`           // 海关三级地址区域编码
	} `json:"boundExtendInfo" mapstructure:"boundExtendInfo"` // 三方保税节点 金额单位 分
	TransferExtendInfo struct {
		InternationalExpressNo string `json:"internationalExpressNo" mapstructure:"internationalExpressNo"` // 国际快递单号
		OrderDeclaredAmount    string `json:"orderDeclaredAmount" mapstructure:"orderDeclaredAmount"`       // 订单申报金额
		PaintMarker            string `json:"paintMarker" mapstructure:"paintMarker"`                       // 大头笔
		CollectionPlace        string `json:"collectionPlace" mapstructure:"collectionPlace"`               // 集包地
		ThreeSegmentCode       string `json:"threeSegmentCode" mapstructure:"threeSegmentCode"`             // 三段码
	} `json:"transferExtendInfo" mapstructure:"transferExtendInfo"` // 小包转运节点
	OpenAddressId             string `json:"openAddressId" mapstructure:"openAddressId"` // 收件人姓名+手机+地址等计算得出，用来查询收件人详情
	SimpleDeliveryPackageList []struct {
		DeliveryPackageIndex string        `json:"deliveryPackageIndex" mapstructure:"deliveryPackageIndex"` // 发货包裹索引标识 修改快递单号会使用
		Status               PackageStatus `json:"status" mapstructure:"status"`                             // 发货包裹状态,1:已下单待付款 2:已支付处理中 3:清关中 4:待发货 6:待收货 7:已完成 8:已关闭 9:已取消 10:换货申请中
		ExpressTrackingNo    string        `json:"expressTrackingNo" mapstructure:"expressTrackingNo"`       // 拆包快递单号
		ExpressCompanyCode   string        `json:"expressCompanyCode" mapstructure:"expressCompanyCode"`     // 快递公司代码
		ItemIdList           []string      `json:"itemIdList" mapstructure:"itemIdList"`                     // 此发货包裹中有哪些商品，status=4待发货时，列表中的item可以拆包发货。status=6时，列表中的item共享相同的快递公司和单号，修改时一起修改
	} `json:"simpleDeliveryPackageList" mapstructure:"simpleDeliveryPackageList"` // 拆包信息节点
	Logistics LogisticsMode `json:"logistics" mapstructure:"logistics"` // 物流模式red_express三方备货直邮(备货海外仓),red_domestic_trade(三方备货内贸),red_standard(三方备货保税仓),red_auto(三方自主发货),red_box(三方小包),red_bonded(三方保税)
}
