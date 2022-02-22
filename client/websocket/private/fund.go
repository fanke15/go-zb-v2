package private

import (
	"github.com/fanke15/go-zb-v2/client/websocket"
	"github.com/fanke15/go-zb-v2/constant"
	"github.com/fanke15/go-zb-v2/handler"
)

//私有数据订阅的频道
const (
	//资金相关
	FundChange      = "Fund.change"
	FundBalance     = "Fund.balance"
	FundGetBill     = "Fund.getBill"
	FundAssetChange = "Fund.assetChange"
	FundAssetInfo   = "Fund.assetInfo"
)

//convertUnit可选范围
const (
	CNY = "cny"
	USD = "usd"
	BTC = "btc"
)

/**
资金相关
*/

type baseFund struct {
	baseSubscriber

	Currency string `json:"currency, omitempty"`
}

func subscribeFund(c *websocket.WSClient, channel string, currency string, futuresAccountType int) {
	sub := new(baseFund)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = channel
	sub.FuturesAccountType = futuresAccountType

	if currency != "" {
		sub.Currency = currency
	}

	subscribe(c, sub)
}

func SubscribeFundChange(c *websocket.WSClient, currency string, futuresAccountType int, handler handler.Handler) {
	channel := FundChange
	subscribeFund(c, channel, currency, futuresAccountType)
	c.AddHandler(channel, handler)
}

func SubscribeFundBalance(c *websocket.WSClient, currency string, futuresAccountType int, handler handler.Handler) {
	channel := FundBalance
	subscribeFund(c, channel, currency, futuresAccountType)
	c.AddHandler(channel, handler)
}

type bill struct {
	baseFund

	BillType  int   `json:"type, omitempty"`
	StartTime int64 `json:"startTime, omitempty"`
	EndTime   int64 `json:"endTime, omitempty"`
	PageNum   int   `json:"pageNum, omitempty"`
	PageSize  int   `json:"pageSize, omitempty"`
}

func SubscribeFundGetBill(c *websocket.WSClient, currency string, futuresAccountType int, handler handler.Handler) {
	sub := new(bill)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = FundGetBill
	sub.FuturesAccountType = futuresAccountType
	sub.Currency = currency
	subscribe(c, sub)

	c.AddHandler(FundGetBill, handler)
}

func SubscribeFundGetBillEx(c *websocket.WSClient, currency string, futuresAccountType, billType int, pageNum, pageSize int, startTime, endTime int64, handler handler.Handler) {
	sub := new(bill)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = FundGetBill
	sub.FuturesAccountType = futuresAccountType

	sub.Currency = currency

	sub.BillType = billType
	sub.PageNum = pageNum
	sub.PageSize = pageSize
	sub.StartTime = startTime
	sub.EndTime = endTime

	subscribe(c, sub)
	c.AddHandler(FundGetBill, handler)
}

type asset struct {
	baseSubscriber

	convertUnit string
}

func SubscribeAssetChange(c *websocket.WSClient, convertUnit string, futuresAccountType int, handler handler.Handler) {
	sub := new(asset)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = FundAssetChange
	sub.FuturesAccountType = futuresAccountType

	sub.convertUnit = convertUnit

	subscribe(c, sub)
	c.AddHandler(FundAssetChange, handler)
}

func SubscribeAssetInfo(c *websocket.WSClient, convertUnit string, futuresAccountType int, handler handler.Handler) {
	sub := new(asset)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = FundAssetInfo
	sub.FuturesAccountType = futuresAccountType

	sub.convertUnit = convertUnit

	subscribe(c, sub)
	c.AddHandler(FundAssetInfo, handler)
}
