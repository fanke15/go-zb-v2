package private

import (
	"github.com/fanke15/go-zb-v2/handler"
)

func (c *Client) GetAccount(convertUnit string, futuresAccountType int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["convertUnit"] = convertUnit

	byte, err := c.Get("/Server/api/v2/Fund/getAccount", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetBill(currency string, futuresAccountType int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["currency"] = currency
	byte, err := c.Get("/Server/api/v2/Fund/getBill", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetBillEx(currency string, futuresAccountType int, billType int, pageNum, pageSize int, startTime, endTime int64, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["currency"] = currency
	params["type"] = billType
	params["pageNum"] = pageNum
	params["pageSize"] = pageSize
	params["startTime"] = startTime
	params["endTime"] = endTime
	byte, err := c.Get("/Server/api/v2/Fund/getBill", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetBillTypeList(handler handler.Handler) {
	byte, err := c.Get("/Server/api/v2/Fund/getBillTypeList", nil)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

//保证金变动历史
func (c *Client) GetMarginHistory(symbol string, futuresAccountType int, marginType int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["symbol"] = symbol
	params["type"] = marginType

	byte, err := c.Get("/Server/api/v2/Fund/marginHistory", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetMarginHistoryEx(symbol string, futuresAccountType int, marginType int, pageNum, pageSize int, startTime, endTime int64, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["symbol"] = symbol
	params["type"] = marginType
	params["pageNum"] = pageNum
	params["pageSize"] = pageSize
	params["startTime"] = startTime
	params["endTime"] = endTime

	byte, err := c.Get("/Server/api/v2/Fund/marginHistory", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) Balance(currencyName string, futuresAccountType int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	if currencyName != "" {
		params["currencyName"] = currencyName
	}
	byte, err := c.Get("/Server/api/v2/Fund/balance", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}
