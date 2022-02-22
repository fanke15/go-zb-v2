package private

import (
	"github.com/fanke15/go-zb-v2/client/websocket"
	"github.com/fanke15/go-zb-v2/constant"
	"github.com/fanke15/go-zb-v2/handler"
)

//私有数据订阅的频道
const (
	//资金相关
	TradeOrderChange      = "Trade.orderChange"
	TradeOrder            = "Trade.order"
	TradeGetOrder         = "Trade.getOrder"
	TradeCancelOrder      = "Trade.cancelOrder"
	TradeBatchCancelOrder = "Trade.batchCancelOrder"
	TradeCancelAllOrder   = "Trade.cancelAllOrders"

	TradeGetUndoneOrder = "Trade.getUndoneOrders"
	TradeGetAllOrder    = "Trade.getAllOrders"
	TradeGetTradeList   = "Trade.getTradeList"
	TradeHistory        = "Trade.tradeHistory"
)

type baseTrade struct {
	baseSubscriber

	Symbol string `json:"symbol"`
}

func SubscribeTradeOrderChange(c *websocket.WSClient, symbol string, futuresAccountType int, handler handler.Handler) {
	sub := new(baseTrade)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = TradeOrderChange
	sub.FuturesAccountType = futuresAccountType

	sub.Symbol = symbol

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

type Order struct {
	baseSubscriber

	Symbol        string  `json:"symbol"`
	Price         float64 `json:"price"`
	Amount        float64 `json:"amount"`
	ActionType    int     `json:"actionType"` //由于baseSubscriber已经有一个action，所以在此换个名字. 1 限价, 2 市价, 3 IOC, 4 只做 maker, 5 FOK
	EntrustType   int     `json:"entrustType"`
	Side          int     `json:"side"`
	ClientOrderId string  `json:"clientOrderId,omitempty"`
	MatchPrice    int     `json:"matchPrice,omitempty"`
}

func SubscribeTradeOrder(c *websocket.WSClient, symbol string, price, amount float64, actionType, entrustType, side int, clientOrderId string, handler handler.Handler) {
	sub := new(Order)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = TradeOrder

	sub.Symbol = symbol
	sub.Price = price
	sub.Amount = amount
	sub.ActionType = actionType
	sub.EntrustType = entrustType
	sub.Side = side
	if clientOrderId != "" {
		sub.ClientOrderId = clientOrderId
	}

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

type GetOrder struct {
	baseSubscriber

	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId,omitempty"`
	ClientOrderId string `json:"clientOrderId,omitempty"`
}

func SubscribeTradeGetOrder(c *websocket.WSClient, symbol string, orderId int64, clientOrderId string, handler handler.Handler) {
	sub := new(GetOrder)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = TradeGetOrder

	sub.Symbol = symbol

	if orderId > 0 {
		sub.OrderId = orderId
	}
	if clientOrderId != "" {
		sub.ClientOrderId = clientOrderId
	}

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

type CancelOrder struct {
	baseSubscriber

	Symbol  string `json:"symbol"`
	OrderId int64  `json:"orderId"`
}

func SubscribeTradeCancelOrder(c *websocket.WSClient, symbol string, orderId int64, handler handler.Handler) {
	sub := new(GetOrder)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = TradeCancelOrder

	sub.Symbol = symbol
	sub.OrderId = orderId

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

type BatchCancelOrder struct {
	baseSubscriber

	Symbol         string   `json:"symbol"`
	OrderIds       []int64  `json:"orderIds,omitempty"`
	ClientOrderIds []string `json:"clientOrderIds,omitempty"`
}

func SubscribeTradeBatchCancelOrder(c *websocket.WSClient, symbol string, orderIds []int64, clientOrderIds []string, handler handler.Handler) {
	sub := new(BatchCancelOrder)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = TradeBatchCancelOrder

	sub.Symbol = symbol
	if orderIds != nil {
		sub.OrderIds = orderIds
	}
	if clientOrderIds != nil {
		sub.ClientOrderIds = clientOrderIds
	}

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

func SubscribeTradeCancelAllOrder(c *websocket.WSClient, symbol string, handler handler.Handler) {
	sub := new(baseTrade)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = TradeCancelAllOrder

	sub.Symbol = symbol

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

type GetUndoneOrders struct {
	baseSubscriber

	Symbol string `json:"symbol"`

	PageNum  int `json:"pageNum,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

func SubscribeTradeGetUndoneOrder(c *websocket.WSClient, symbol string, handler handler.Handler) {
	sub := new(GetUndoneOrders)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = TradeGetUndoneOrder

	sub.Symbol = symbol

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

func SubscribeTradeGetUndoneOrderEx(c *websocket.WSClient, symbol string, pageNum, pageSize int, handler handler.Handler) {
	sub := new(GetUndoneOrders)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = TradeGetUndoneOrder

	sub.Symbol = symbol
	sub.PageNum = pageNum
	sub.PageSize = pageSize

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

type GetAllOrder struct {
	baseSubscriber

	Symbol string `json:"symbol"`

	StartTime int64 `json:"startTime,omitempty"`
	EndTime   int64 `json:"endTime,omitempty"`
	PageNum   int   `json:"pageNum,omitempty"`
	PageSize  int   `json:"pageSize,omitempty"`
}

func SubscribeTradeGetAllOrder(c *websocket.WSClient, symbol string, handler handler.Handler) {
	sub := new(GetAllOrder)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = TradeGetAllOrder

	sub.Symbol = symbol

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}
func SubscribeTradeGetAllOrderEx(c *websocket.WSClient, symbol string, startTime, endTime int64, pageNum, pageSize int, handler handler.Handler) {
	sub := new(GetAllOrder)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = TradeGetAllOrder

	sub.Symbol = symbol
	sub.StartTime = startTime
	sub.EndTime = endTime
	sub.PageNum = pageNum
	sub.PageSize = pageSize

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

type GetTradeList struct {
	baseSubscriber

	Symbol  string `json:"symbol"`
	OrderId int64  `json:"orderId"`
}

func SubscribeTradeGetTradeList(c *websocket.WSClient, symbol string, orderId int64, handler handler.Handler) {
	sub := new(GetTradeList)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = TradeGetTradeList

	sub.Symbol = symbol
	sub.OrderId = orderId

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

func SubscribeTradeHistory(c *websocket.WSClient, symbol string, handler handler.Handler) {
	sub := new(baseTrade)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = TradeHistory

	sub.Symbol = symbol

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}
