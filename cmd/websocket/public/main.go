package main

import (
	"github.com/fanke15/go-zb-v2/client/websocket"
	"github.com/fanke15/go-zb-v2/client/websocket/public"
	handler "github.com/fanke15/go-zb-v2/handler/public"
	"sync"
	"time"
)

func main() {
	url := "ws://ttfutures.zb.com/ws/public/v1"
	//url := "ws://172.16.100.178:9005/ws/public/v1"
	//url := "wss://futures.zb.today/ws/public/v1"
	//url := "wss://dist.zb.today/ws/public/v1"
	c := websocket.NewWSClient(url, time.Second*10)
	symbol := "ETH_USDT"
	size := 10

	//public.SubscribeWholeDepth(c, symbol, size, handler.DepthHandler)
	//public.SubscribeDepth(c, symbol, size, handler.DepthHandler)
	public.SubscribeKLine(c, symbol, "1M", size, handler.KLineHandler)
	//public.SubscribeTrade(c, symbol, size, handler.TradeHandler)
	//public.SubscribeTicker(c, symbol, handler.TickerWSHandler)
	//public.SubscribeAllTicker(c, handler.AllTickerHandler)
	//public.SubscribeMarkPrice(c, symbol, handler.MarkIndexPriceWSHandler)
	//public.SubscribeIndexPrice(c, symbol, handler.MarkIndexPriceWSHandler)
	//public.SubscribeMarkKLine(c, symbol, "1M", size, handler.KLineHandler)
	//public.SubscribeIndexKLine(c, symbol, "1M", size, handler.KLineHandler)
	//public.SubscribeFundingRate(c, symbol, handler.FundingRateHandler)

	//public.SubscribeAll(client, "ETH_USDT")

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func SubscribeAll(c *websocket.WSClient, symbol string, size int) {
	public.SubscribeWholeDepth(c, symbol, size, handler.DepthHandler)
	public.SubscribeDepth(c, symbol, size, handler.DepthHandler)
	public.SubscribeKLine(c, symbol, "1M", size, handler.KLineHandler)
	public.SubscribeTrade(c, symbol, size, handler.TickerHandler)
	public.SubscribeTicker(c, symbol, handler.TickerHandler)
	public.SubscribeAllTicker(c, handler.AllTickerHandler)
	public.SubscribeMarkPrice(c, symbol, handler.MarkIndexPriceHandler)
	public.SubscribeIndexPrice(c, symbol, handler.MarkIndexPriceHandler)
	public.SubscribeMarkKLine(c, symbol, "1M", size, handler.KLineHandler)
	public.SubscribeIndexKLine(c, symbol, "1M", size, handler.KLineHandler)
	public.SubscribeFundingRate(c, symbol, handler.FundingRateHandler)
}
