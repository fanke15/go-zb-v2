package public

import (
	"github.com/fanke15/go-zb-v2/handler"
)

//用于获取公共行情
func (c *Client) GetWholeDepth(symbol string, size int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["size"] = size

	byte, err := c.Get("/api/public/v1/depth", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetKLine(symbol string, period string, size int, handler handler.Handler) {
	c.getKLine("/api/public/v1/kline", symbol, period, size, handler)
}

func (c *Client) GetMarkKline(symbol string, period string, size int, handler handler.Handler) {
	c.getKLine("/api/public/v1/markKline", symbol, period, size, handler)
}

func (c *Client) GetIndexKLine(symbol string, period string, size int, handler handler.Handler) {
	c.getKLine("/api/public/v1/indexKline", symbol, period, size, handler)
}

func (c *Client) getKLine(klineType string, symbol string, period string, size int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["period"] = period
	params["size"] = size

	byte, err := c.Get(klineType, params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetTrade(symbol string, size int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["size"] = size

	byte, err := c.Get("/api/public/v1/trade", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetTicker(symbol string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol

	byte, err := c.Get("/api/public/v1/ticker", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetAllTicker(handler handler.Handler) {
	byte, err := c.Get("/api/public/v1/ticker", nil)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetMarkPrice(symbol string, handler handler.Handler) {
	c.getPrice("/api/public/v1/markPrice", symbol, handler)
}

func (c *Client) GetIndexPrice(symbol string, handler handler.Handler) {
	c.getPrice("/api/public/v1/indexPrice", symbol, handler)
}

func (c *Client) getPrice(priceType string, symbol string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol

	byte, err := c.Get(priceType, params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetFundingRate(symbol string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol

	byte, err := c.Get("/api/public/v1/fundingRate", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

//查询市场强平订单
func (c *Client) GetAllForceOrders(symbol string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol

	byte, err := c.Get("/Server/api/v2/allForceOrders", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}
func (c *Client) GetAllForceOrdersEx(symbol string, startTime, endTime int64, limit int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}

	byte, err := c.Get("/Server/api/v2/allForceOrders", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

//大户账户数多空比
func (c *Client) GetTopLongShortAccountRatio(symbol, period string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["period"] = period //"5m","15m","30m","1h","2h","4h","6h","12h","1d"

	byte, err := c.Get("/Server/api/v2/data/topLongShortAccountRatio", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}
func (c *Client) GetTopLongShortAccountRatioEx(symbol, period string, startTime, endTime int64, limit int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["period"] = period
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}

	byte, err := c.Get("/Server/api/v2/data/topLongShortAccountRatio", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

//大户持仓量多空比
func (c *Client) GetTopLongShortPositionRatio(symbol, period string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["period"] = period

	byte, err := c.Get("/Server/api/v2/data/topLongShortPositionRatio", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}
func (c *Client) GetTopLongShortPositionRatioEx(symbol, period string, startTime, endTime int64, limit int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["period"] = period
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if limit > 0 {
		params["limit"] = limit
	}

	byte, err := c.Get("/Server/api/v2/data/topLongShortPositionRatio", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

//zb现货兑换价格
func (c *Client) GetSpotPrice(symbol, side string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["side"] = side //asks, bids

	byte, err := c.Get("/api/public/v1/spotPrice", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}
