package public

import (
	"encoding/json"
	"github.com/fanke15/go-zb-v2/client/websocket"
	"github.com/fanke15/go-zb-v2/constant"
	"github.com/fanke15/go-zb-v2/handler"
	. "github.com/fanke15/go-zb-v2/log"
)

/**
用于websocket公共行情
*/
type Subscriber struct {
	Action  string `json:"action"`
	Channel string `json:"channel"`
	Size    int    `json:"size",omitempty`
}

func UnSubscribe(c *websocket.WSClient, channel string) {
	sendSubscriber(c, constant.UNSUBSCRIBE, channel, 0)
	c.DeleteHandler(channel)
}

func subscribe(c *websocket.WSClient, channel string, size int) {
	sendSubscriber(c, constant.SUBSCRIBE, channel, size)
}

func sendSubscriber(c *websocket.WSClient, action string, channel string, size int) bool {
	sub := new(Subscriber)
	sub.Action = action
	sub.Channel = channel
	if size > 0 {
		sub.Size = size
	}

	data, err := json.Marshal(sub)
	if err != nil {
		Log.Error().Msg(err.Error())
		return false
	}

	c.Write(data)

	return true
}

func SubscribeWholeDepth(c *websocket.WSClient, symbol string, size int, handler handler.Handler) {
	channel := symbol + "." + "DepthWhole"
	subscribe(c, channel, size)

	c.AddHandler(channel, handler)
}

func SubscribeDepth(c *websocket.WSClient, symbol string, size int, handler handler.Handler) {
	channel := symbol + "." + "Depth"
	subscribe(c, channel, size)
	c.AddHandler(channel, handler)
}

//period可选值为：1M,5M,15M, 30M, 1H, 6H, 1D, 5D
func SubscribeKLine(c *websocket.WSClient, symbol string, period string, size int, handler handler.Handler) {
	channel := symbol + "." + "KLine_" + period
	subscribe(c, channel, size)
	c.AddHandler(channel, handler)
}

func SubscribeTrade(c *websocket.WSClient, symbol string, size int, handler handler.Handler) {
	channel := symbol + "." + "Trade"
	subscribe(c, channel, size)
	c.AddHandler(channel, handler)
}

func SubscribeTicker(c *websocket.WSClient, symbol string, handler handler.Handler) {
	channel := symbol + "." + "Ticker"
	subscribe(c, channel, 0)
	c.AddHandler(channel, handler)
}
func SubscribeAllTicker(c *websocket.WSClient, handler handler.Handler) {
	channel := "All.Ticker"
	subscribe(c, channel, 0)
	c.AddHandler(channel, handler)
}
func SubscribeMarkPrice(c *websocket.WSClient, symbol string, handler handler.Handler) {
	channel := symbol + "." + "mark"
	subscribe(c, channel, 0)
	c.AddHandler(channel, handler)
}

func SubscribeIndexPrice(c *websocket.WSClient, symbol string, handler handler.Handler) {
	channel := symbol + "." + "index"
	subscribe(c, channel, 0)
	c.AddHandler(channel, handler)
}

func SubscribeMarkKLine(c *websocket.WSClient, symbol string, period string, size int, handler handler.Handler) {
	channel := symbol + "." + "mark_" + period
	subscribe(c, channel, size)
	c.AddHandler(channel, handler)
}

func SubscribeIndexKLine(c *websocket.WSClient, symbol string, period string, size int, handler handler.Handler) {
	channel := symbol + "." + "index_" + period
	subscribe(c, channel, size)
	c.AddHandler(channel, handler)
}

func SubscribeFundingRate(c *websocket.WSClient, symbol string, handler handler.Handler) {
	channel := symbol + "." + "FundingRate"
	subscribe(c, channel, 0)
	c.AddHandler(channel, handler)
}
