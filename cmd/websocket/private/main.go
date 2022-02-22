package main

import (
	"github.com/fanke15/go-zb-v2/client/websocket"
	"github.com/fanke15/go-zb-v2/client/websocket/private"
	"github.com/fanke15/go-zb-v2/constant"
	handler "github.com/fanke15/go-zb-v2/handler/private"
	"sync"
	"time"
)

func main() {
	url := "wss://futures.zb.today/ws/private/api/v2"
	//url := "ws://172.16.100.179:9012/ws/private/api/v2"
	//url := "ws://127.0.0.1:9012/ws/private/api/v2"
	c := websocket.NewWSClient(url, time.Second*1000)

	futuresAccountType := constant.BaseUsdt

	private.Login(c)

	////--------Fund--------
	currry := "USD"
	private.SubscribeFundChange(c, currry, futuresAccountType, handler.FundChangeHandler)
	private.SubscribeFundBalance(c, currry, futuresAccountType, handler.BalanceHandler)
	private.SubscribeFundGetBill(c, currry, futuresAccountType, handler.GetBillHandler)

	convertUnit := private.CNY
	private.SubscribeAssetChange(c, convertUnit, futuresAccountType, handler.AssetChangeHandler)
	private.SubscribeAssetInfo(c, convertUnit, futuresAccountType, handler.AssetInfoHandler)

	////--------Position--------
	symbol := "ETH_USDT"
	private.SubscribePositionsChange(c, symbol, futuresAccountType, handler.PositionsChangeHandler)
	private.SubscribeGetPositions(c, symbol, futuresAccountType, handler.GetPositionsHandler)
	private.SubscribeMarginInfo(c, 6785805193549195299, futuresAccountType, handler.MarginInfoHandler)
	private.SubscribeUpdateMargin(c, 6785805193549195299, futuresAccountType, 1, 1, handler.UpdateMarginHandler)

	private.SubscribeGetSetting(c, symbol, futuresAccountType, handler.GetSettingHandler)
	private.SubscribeSetLeverage(c, symbol, futuresAccountType, 12, handler.SetLeverageHandler)
	//private.SubscribeSetPositionsMode(c, symbol, futuresAccountType, constant.BiDirection, handler.SetPositionsModeHandler) 	//暂不支持SetPositionsMode接口
	//private.SubscribeSetMarginMode(c, symbol, futuresAccountType, constant.Isolated, handler.SetMarginModeHandler)			//暂不支持SetMarginMode接口
	private.SubscribeGetNominalValue(c, symbol, futuresAccountType, constant.SideOpenLong, handler.GetNominalValueHandler)

	////--------Trade--------
	private.SubscribeTradeOrderChange(c, symbol, futuresAccountType, handler.OrderChangeHandler)
	private.SubscribeTradeOrder(c, symbol, 2022.1, 1, 1, 1, 1, "", handler.OrderHandler)
	private.SubscribeTradeGetAllOrder(c, symbol, handler.GetAllOrderHandler)
	private.SubscribeTradeGetUndoneOrder(c, symbol, handler.GetUndoneOrderHandler)
	private.SubscribeTradeGetOrder(c, symbol, 6786121423673892864, "", handler.GetOrderHandler)
	private.SubscribeTradeGetTradeList(c, symbol, 6785805407710355456, handler.GetTradeListHandler)
	private.SubscribeTradeHistory(c, symbol, handler.TradeHistoryHandler)
	private.SubscribeTradeCancelOrder(c, symbol, 6786125344224059392, handler.CancelOrderHandler)
	private.SubscribeTradeBatchCancelOrder(c, symbol, []int64{6786125434909106176, 6786125435013963786}, nil, handler.BatchCancelOrderHandler)
	private.SubscribeTradeCancelAllOrder(c, symbol, handler.CancelAllOrderHandler)

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
