package main

import (
	"github.com/fanke15/go-zb-v2/client/httpclient/private"
	"github.com/fanke15/go-zb-v2/constant"
	handler "github.com/fanke15/go-zb-v2/handler/private"
)

func main() {
	apiKey := "99730176-b026-4250-ab01-7eab6080efae"
	secretKey := "22383e66-3be5-425c-b1ad-3d287d79322c"
	//url := "https://futures.zb.com"
	url := "https://fapi.zb.com"
	client := private.NewPrivateClient(url, "usdt", apiKey, secretKey)

	symbol := "ETH_USDT"

	//--------Fund--------
	futuresAccountType := constant.BaseUsdt
	convertUnit := constant.CNY
	client.GetAccount(convertUnit, futuresAccountType, handler.GetAccountHandler)
	client.Balance(constant.USDT, futuresAccountType, handler.BalanceHandler)
	client.Balance("", futuresAccountType, handler.BalanceHandler)

	currency := "USD"
	client.GetBill(currency, futuresAccountType, handler.GetBillHandler)
	client.GetBillTypeList(handler.GetBillTypeListHandler)
	client.GetMarginHistory(symbol, futuresAccountType, constant.ADD, handler.GetMarginHistoryHandler)

	//--------Position--------
	client.GetPositions(symbol, futuresAccountType, handler.GetPositionsHandler)

	var positionsId int64 = 6785805193549195299
	client.MarginInfo(positionsId, futuresAccountType, handler.MarginInfoHandler)
	client.UpdateMargin(positionsId, futuresAccountType, 1, constant.SUBTRACT, handler.UpdateMarginHandler)
	client.SetLeverage(symbol, futuresAccountType, 12, handler.SetLeverageHandler)
	client.SetPositionsMode(symbol, futuresAccountType, constant.OneDirection, handler.SetPositionsModeHandler)
	client.SetMarginMode(symbol, futuresAccountType, constant.Isolated, handler.SetMarginModeHandler)
	client.GetNominalValue(symbol, futuresAccountType, constant.SideOpenLong, handler.GetNominalValueHandler)
	client.GetPositionSetting(symbol, futuresAccountType, handler.GetSettingHandler)
	client.SetAutoAppendMargin(symbol, futuresAccountType, 1212, handler.SetAutoAppendMarginHandler)
	client.SetMarginCoins(symbol, futuresAccountType, "eth,usdt", handler.SetMarginCoinsHandler)

	//--------Trade--------
	client.Order(symbol, 2022.1, 1, 1, 1, 1, handler.OrderHandler)
	client.Order(symbol, 2022.1, 1, 1, 1, 1, handler.OrderHandler)
	client.Order(symbol, 2022.1, 1, 1, 1, 1, handler.OrderHandler)
	client.GetAllOrder(symbol, handler.GetAllOrderHandler)
	client.GetUndoneOrder(symbol, handler.GetUndoneOrderHandler)
	client.GetOrder(symbol, 6787729982828322816, "", handler.GetOrderHandler)
	client.GetTradeList(symbol, 6787729982828322816, handler.GetTradeListHandler)
	client.TradeHistory(symbol, handler.TradeHistoryHandler)
	client.CancelOrder(symbol, 6786125344224059392, "", handler.CancelOrderHandler)
	client.BatchCancelOrder(symbol, []int64{6786125434909106176, 6786125435013963786}, nil, handler.BatchCancelOrderHandler)
	client.CancelAllOrder(symbol, handler.CancelAllOrderHandler)
	client.OrderAlgo(symbol, 1, 1, 2, 1270, 1280, 1, 1, handler.OrderHandler)
	client.CancelAlgos(symbol, []string{}, handler.BatchCancelOrderHandler)
	client.GetOrderAlgos(symbol, handler.AlgoOrderHandler)
}
