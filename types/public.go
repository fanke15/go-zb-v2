package types

/**
公共行情相关的数据结构
*/

type Depth struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

type KLine [][]interface{}
type Trade [][]interface{}
type AllTicker map[string][]interface{}
type Ticker []interface{}

type FundingRate struct {
	Rate              string `json:"fundingRate"`
	NextCalculateTime string `json:"nextCalculateTime"`
}
type ForceOrder struct {
	Symbol        string  `json:"symbol"`
	Price         float64 `json:"price,string"`
	Amount        float64 `json:"amount,string"`
	TradeAmount   float64 `json:"tradeAmount,string"`
	TradeAvgPrice float64 `json:"tradeAvgPrice,string"`
	Side          string  `json:"side"`
	status        string  `json:"status"`
	Time          int     `json:"time,string"`
}
type TopLongShortAccountRatio struct {
	Symbol         string  `json:"symbol"`
	Timestamp      int     `json:"timestamp,string"`
	LongAccount    int     `json:"longAccount,string"`
	ShortAccount   int     `json:"shortAccount,string"`
	LongShortRatio float64 `json:"longShortRatio,string"`
}

type TopLongShortPositionRatio struct {
	Symbol         string  `json:"symbol"`
	Timestamp      int     `json:"timestamp,string"`
	LongPosition   float64 `json:"longPosition,string"`
	ShortPosition  float64 `json:"shortPosition,string"`
	LongShortRatio float64 `json:"longShortRatio,string"`
}
