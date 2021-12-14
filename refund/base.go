package refund

// ReturnStatus 售后状态
type ReturnStatus uint

const (
	RETURN_NONE ReturnStatus = iota // 不传默认，全部
	RETURN_1                        // 待审核
	RETURN_2                        // 待用户寄回
	RETURN_3                        // 待收货
	RETURN_4                        // 完成
	RETURN_5                        // 取消
	RETURN_6                        // 关闭
	_
	_
	RETURN_9                        // 拒绝
	RETURN_9999 ReturnStatus = 9999 // 删除
)

// Uint ...
func (r ReturnStatus) Uint() uint {
	return uint(r)
}

// ReturnSubStats 售后子状态
type ReturnSubStats uint

const (
	SUB_301 ReturnSubStats = iota + 301 // 待审核
	SUB_302                             // 快递已签收
	_
	SUB_304 // 收获异常
)

// Uint ...
func (r ReturnSubStats) Uint() uint {
	return uint(r)
}

// ReturnType 售后类型
type ReturnType uint

const (
	TYPE_NONE ReturnType = iota // 不传/0:全部
	TYPE_1                      // 退货退款
	TYPE_2                      // 换货
	TYPE_3                      // 仅退款（old
	TYPE_4                      // 仅退款（new
)

// Uint ...
func (r ReturnType) Uint() uint {
	return uint(r)
}

// IsShipNeeded 是否寄回
type IsShipNeeded bool

const (
	SHIP_NEEDLESS = iota // 无需寄回
	SHIP_NEEDED          // 需寄回
)

// Bool ...
func (i IsShipNeeded) Bool() bool {
	return bool(i)
}
