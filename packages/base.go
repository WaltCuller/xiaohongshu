package packages

type LogisticsMode string

const (
	EXPRESS        LogisticsMode = "red_express"
	DOMESTIC_TRADE LogisticsMode = "red_domestic_trade"
	STANDARD       LogisticsMode = "red_standard"
	AUTO           LogisticsMode = "red_auto"
	BOX            LogisticsMode = "red_box"
	BONDED         LogisticsMode = "red_bonded"
)

type OrderStatus string

const (
	WAITING  OrderStatus = "waiting"
	SHIPPED  OrderStatus = "shipped"
	RECEIVED OrderStatus = "received"
)
