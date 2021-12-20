package packages

const (
	METHOD_CANCEL_APPLY_LIST = "package.getCancelApplyList"
)

// 审核状态：0 所有，1 待审核，2 已审核
type Status int

const (
	STATUS_0 = iota // 所有
	STATUS_1        // 待审核
	STATUS_2        // 已审核
)

// ReqCancelList 取消订单列表参数
type ReqCancelList struct {
	Logistics  LogisticsMode `json:"logistics" param:"logistics"`   // 物流模式
	Status     Status        `json:"status" param:"status"`         // 审核状态
	PageNo     int           `json:"pageNo" param:"pageNo"`         // 从1开始，限制小于100
	PageSize   int           `json:"pageSize" param:"pageSize"`     // 限制小于100
	StartTime  int64         `json:"startTime" param:"startTime"`   // 时间范围起点
	EndTime    int64         `json:"endTime" param:"endTime"`       // 时间范围终点
	TimeType   TimeType      `json:"timeType" param:"timeType"`     // startTime/endTime对应的时间类型，1 创建时间 限制 end-start<=24h，2 更新时间 限制 end-start<=30min
	UseHasNext bool          `json:"useHasNext" param:"useHasNext"` // true表示分页不返会total 返回 hasNext = true 表示仍有数据，false 返回total
	PackageID  string        `json:"packageId" param:"packageId"`   // 包裹号
}

// RspCancelList 取消订单列表返回
type RspCancelList struct {
	Total            int  `json:"total"`    // 总数 useHasNext=true时为0
	PageNo           int  `json:"pageNo"`   // 当前页数
	PageSize         int  `json:"pageSize"` // 页大小
	HasNext          bool `json:"hasNext"`  // useHasNext=true时 hasNext小时仍有数据
	CancelRecordList []struct {
		PackageId     string        `json:"packageId"`     // 包裹号
		CreatedTime   int64         `json:"createdTime"`   // 申请取消时间
		Status        Status        `json:"status"`        // 审核状态 1 待审核 2 已审核
		CancelReason  string        `json:"cancelReason"`  // 取消原因
		AuditTime     int64         `json:"auditTime"`     // 审核时间
		LastAuditTime int64         `json:"lastAuditTime"` // 最晚审核时间，超过此时间会自动审核
		Logistics     LogisticsMode `json:"logistics"`     // 物流模式
		AuditResult   string        `json:"auditResult"`   // 审核结果
		AuditReason   string        `json:"auditReason"`   // 审核理由
		CancelId      string        `json:"cancelId"`      // 取消ID，标识一次取消申请
		UpdateTime    int64         `json:"updateTime"`    // 更新时间
		Operator      string        `json:"operator"`      // 操作人
	} `json:"cancelRecordList"` // 取消申请列表
	MaxPageNo int `json:"maxPageNo"` // 最大页码数
}
