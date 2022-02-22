package private

import (
	"github.com/fanke15/go-zb-v2/handler"
	"github.com/fanke15/go-zb-v2/types"
)

//actionType: 1 限价, 2 市价, 3 IOC, 4 只做 maker, 5 FOK
func (c *Client) Order(symbol string, price, amount float64, actionType, entrustType, side int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["price"] = price
	params["amount"] = amount
	params["action"] = actionType
	params["entrustType"] = entrustType
	params["side"] = side
	byte, err := c.Post("/Server/api/v2/trade/order", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) BatchOrder(orders []types.Order, handler handler.Handler) {
	params := make(map[string]interface{})
	params["orderDatas"] = orders

	byte, err := c.Post("/Server/api/v2/trade/batchOrder", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}
func (c *Client) CancelOrder(symbol string, orderId int64, clientOrderId string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	if orderId > 0 {
		params["orderId"] = orderId
	}
	if clientOrderId != "" {
		params["clientOrderId"] = clientOrderId
	}

	byte, err := c.Post("/Server/api/v2/trade/cancelOrder", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) BatchCancelOrder(symbol string, orderIds []int64, clientOrderIds []string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	if orderIds != nil {
		params["orderIds"] = orderIds
	}
	if clientOrderIds != nil {
		params["clientOrderIds"] = clientOrderIds
	}

	byte, err := c.Post("/Server/api/v2/trade/batchCancelOrder", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) CancelAllOrder(symbol string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	byte, err := c.Post("/Server/api/v2/trade/cancelAllOrders", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetUndoneOrder(symbol string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	byte, err := c.Get("/Server/api/v2/trade/getUndoneOrders", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetUndoneOrderEx(symbol string, pageNum, pageSize int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["pageNum"] = pageNum
	params["pageSize"] = pageSize
	byte, err := c.Get("/Server/api/v2/trade/getUndoneOrders", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetOrder(symbol string, orderId int64, clientOrderId string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	if orderId > 0 {
		params["orderId"] = orderId
	}
	if clientOrderId != "" {
		params["clientOrderId"] = clientOrderId
	}
	byte, err := c.Get("/Server/api/v2/trade/getOrder", params)
	if err != nil {
		return
	}
	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetAllOrder(symbol string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	byte, err := c.Get("/Server/api/v2/trade/getAllOrders", params)
	if err != nil {
		return
	}
	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetAllOrderEx(symbol string, startTime, endTime int64, pageNum, pageSize int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["startTime"] = startTime
	params["endTime"] = endTime
	params["pageNum"] = pageNum
	params["pageSize"] = pageSize
	byte, err := c.Get("/Server/api/v2/trade/getAllOrders", params)
	if err != nil {
		return
	}
	if handler != nil {
		handler(*byte)
	}
}
func (c *Client) GetTradeList(symbol string, orderId int64, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["orderId"] = orderId
	byte, err := c.Get("/Server/api/v2/trade/getTradeList", params)
	if err != nil {
		return
	}
	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) TradeHistory(symbol string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	byte, err := c.Get("/Server/api/v2/trade/tradeHistory", params)
	if err != nil {
		return
	}
	if handler != nil {
		handler(*byte)
	}
}
func (c *Client) OrderAlgo(symbol string, side, orderType int, amount, triggerPrice, algoPrice float64, priceType, bizType int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["side"] = side
	params["orderType"] = orderType
	params["amount"] = amount
	params["triggerPrice"] = triggerPrice
	params["algoPrice"] = algoPrice
	params["priceType"] = priceType
	params["bizType"] = bizType
	byte, err := c.Post("/Server/api/v2/trade/orderAlgo", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}
func (c *Client) CancelAlgos(symbol string, ids []string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["ids"] = ids

	byte, err := c.Post("/Server/api/v2/trade/cancelAlgos", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}
func (c *Client) GetOrderAlgos(symbol string, handler handler.Handler) {
	c.GetOrderAlgosEx(symbol, 0, 0, 0, 0, 0, 0, 0, 0, handler)
}
func (c *Client) GetOrderAlgosEx(symbol string, side, orderType, bizType, status int, startTime, endTime int64, pageNum, pageSize int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	if side > 0 {
		params["side"] = side
	}
	if orderType > 0 {
		params["orderType"] = orderType
	}
	if bizType > 0 {
		params["bizType"] = bizType
	}
	if status > 0 {
		params["status"] = status
	}
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}
	if pageNum > 0 {
		params["pageNum"] = pageNum
	}
	if pageSize > 0 {
		params["pageSize"] = pageSize
	}
	byte, err := c.Get("/Server/api/v2/trade/getOrderAlgos", params)
	if err != nil {
		return
	}
	if handler != nil {
		handler(*byte)
	}
}
