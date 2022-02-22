package constant

const (
	SUCCESS = 10000

	GET  = "GET"
	POST = "POST"

	BaseUsdt = 1 //USDT永续合约
	BaseCoin = 2 //币本位合约

	CNY = "cny"
	USD = "usd"
	BTC = "btc"

	USDT = "usdt"

	//    1开多（买入），2开空（卖出），3平多（卖出），4平空（买入）
	SideOpenLong   = 1 //多
	SideOpenShort  = 2 //多
	SideCloseLong  = 3 //多
	SideCloseShort = 4 //多

	OneDirection = 1
	BiDirection  = 2

	Isolated = 1
	Cross    = 2

	ADD      = 1
	SUBTRACT = 0
)
