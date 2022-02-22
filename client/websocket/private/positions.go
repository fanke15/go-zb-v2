package private

import (
	"github.com/fanke15/go-zb-v2/client/websocket"
	"github.com/fanke15/go-zb-v2/constant"
	"github.com/fanke15/go-zb-v2/handler"
)

//私有数据订阅的频道
const (
	//资金相关
	PositionsChange           = "Positions.change"
	PositionsGetPositions     = "Positions.getPositions"
	PositionsMarginInfo       = "Positions.marginInfo"
	PositionsUpdateMargin     = "Positions.updateMargin"
	PositionsGetSetting       = "Positions.getSetting"
	PositionsSetLeverage      = "Positions.setLeverage"
	PositionsSetPositionsMode = "Positions.setPositionsMode"
	PositionsSetMarginMode    = "Positions.setMarginMode"
	PositionsGetNominalValue  = "Positions.getNominalValue"
)

type Positions struct {
	baseSubscriber

	Symbol string `json:"symbol,omitempty"`
}

func SubscribePositionsChange(c *websocket.WSClient, symbol string, futuresAccountType int, handler handler.Handler) {
	sub := new(Positions)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = PositionsChange
	sub.FuturesAccountType = futuresAccountType

	sub.Symbol = symbol

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

func SubscribeGetPositions(c *websocket.WSClient, symbol string, futuresAccountType int, handler handler.Handler) {
	sub := new(Positions)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = PositionsGetPositions
	sub.FuturesAccountType = futuresAccountType

	sub.Symbol = symbol

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

type Margin struct {
	baseSubscriber

	PositionsId int64 `json:"positionsId"`

	Amount     float64 `json:"amount,omitempty"`
	MarginType int     `json:"type,omitempty"`
}

func SubscribeMarginInfo(c *websocket.WSClient, positionsId int64, futuresAccountType int, handler handler.Handler) {
	sub := new(Margin)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = PositionsMarginInfo
	sub.FuturesAccountType = futuresAccountType

	sub.PositionsId = positionsId

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

func SubscribeUpdateMargin(c *websocket.WSClient, positionsId int64, futuresAccountType int, amount float64, marginType int, handler handler.Handler) {
	sub := new(Margin)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = PositionsUpdateMargin
	sub.FuturesAccountType = futuresAccountType

	sub.PositionsId = positionsId
	sub.Amount = amount
	sub.MarginType = marginType

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

func SubscribeGetSetting(c *websocket.WSClient, symbol string, futuresAccountType int, handler handler.Handler) {
	sub := new(Positions)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = PositionsGetSetting
	sub.FuturesAccountType = futuresAccountType

	sub.Symbol = symbol

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

type Leverage struct {
	baseSubscriber

	Symbol   string `json:"symbol"`
	Leverage int    `json:"leverage"`
}

func SubscribeSetLeverage(c *websocket.WSClient, symbol string, futuresAccountType int, leverage int, handler handler.Handler) {
	sub := new(Leverage)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = PositionsSetLeverage
	sub.FuturesAccountType = futuresAccountType

	sub.Symbol = symbol
	sub.Leverage = leverage

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

type PositionsMode struct {
	baseSubscriber

	Symbol        string `json:"symbol"`
	PositionsMode int    `json:"positionsMode"`
}

func SubscribeSetPositionsMode(c *websocket.WSClient, symbol string, futuresAccountType int, positionsMode int, handler handler.Handler) {
	sub := new(PositionsMode)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = PositionsSetPositionsMode
	sub.FuturesAccountType = futuresAccountType

	sub.Symbol = symbol
	sub.PositionsMode = positionsMode

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

type MarginMode struct {
	baseSubscriber

	Symbol     string `json:"symbol"`
	MarginMode int    `json:"marginMode"`
}

func SubscribeSetMarginMode(c *websocket.WSClient, symbol string, futuresAccountType int, marginMode int, handler handler.Handler) {
	sub := new(MarginMode)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = PositionsSetMarginMode
	sub.FuturesAccountType = futuresAccountType

	sub.Symbol = symbol
	sub.MarginMode = marginMode

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}

type NominalValue struct {
	baseSubscriber

	Symbol string `json:"symbol"`
	Side   int    `json:"side"`
}

func SubscribeGetNominalValue(c *websocket.WSClient, symbol string, futuresAccountType int, side int, handler handler.Handler) {
	sub := new(NominalValue)
	sub.Action = constant.SUBSCRIBE
	sub.Channel = PositionsGetNominalValue
	sub.FuturesAccountType = futuresAccountType

	sub.Symbol = symbol
	sub.Side = side

	subscribe(c, sub)
	c.AddHandler(sub.Channel, handler)
}
