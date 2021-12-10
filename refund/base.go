package refund

type ReturnStatus uint

// 售后状态 不传/0:全部 1:待审核 2:待用户寄回 3:待收货 4:完成 5:取消 6:关闭 9:拒绝 9999:删除
const (
	RETURN_NONE ReturnStatus = iota
	RETURN_1
	RETURN_2
	RETURN_3
	RETURN_4
	RETURN_5
	RETURN_6
	_
	_
	RETURN_9
	RETURN_9999 ReturnStatus = 9999
)

func (r ReturnStatus) Uint() uint {
	return uint(r)
}

type ReturnSubStats uint

// 售后子状态 301-待审核 302-快递已签收 304-收货异常
const (
	SUB_301 ReturnSubStats = iota + 301
	SUB_302
	_
	SUB_304
)

func (r ReturnSubStats) Uint() uint {
	return uint(r)
}

type ReturnType uint

// 售后类型 不传/0:全部 1:退货退款 2:换货 3:仅退款
const (
	TYPE_NONE ReturnType = iota
	TYPE_1
	TYPE_2
	TYPE_3
)

func (r ReturnType) Uint() uint {
	return uint(r)
}

type IsShipNeeded bool

const (
	SHIP_NEEDLESS = iota
	SHIP_NEEDED
)

func (i IsShipNeeded) Bool() bool {
	return bool(i)
}
