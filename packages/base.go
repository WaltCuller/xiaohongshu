package packages

type LogisticsMode string

const (
	EXPRESS        LogisticsMode = "red_express"        // 三方备货直邮(备货海外仓)
	DOMESTIC_TRADE LogisticsMode = "red_domestic_trade" // (三方备货内贸)
	STANDARD       LogisticsMode = "red_standard"       // (三方备货保税仓)
	AUTO           LogisticsMode = "red_auto"           // (三方自主发货)
	BOX            LogisticsMode = "red_box"            // (三方小包)
	BONDED         LogisticsMode = "red_bonded"         // (三方保税)
)

func (l LogisticsMode) String() string {
	return string(l)
}
