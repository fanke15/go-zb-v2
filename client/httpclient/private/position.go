package private

import (
	"github.com/fanke15/go-zb-v2/handler"
)

func (c *Client) GetPositions(symbol string, futuresAccountType int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["symbol"] = symbol
	byte, err := c.Get("/Server/api/v2/Positions/getPositions", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) MarginInfo(positionsId int64, futuresAccountType int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["positionsId"] = positionsId
	byte, err := c.Get("/Server/api/v2/Positions/marginInfo", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

//marginType: 1: 增加  0：减少
func (c *Client) UpdateMargin(positionsId int64, futuresAccountType int, amount float64, marginType int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["positionsId"] = positionsId
	params["amount"] = amount
	params["type"] = marginType
	byte, err := c.Post("/Server/api/v2/Positions/updateMargin", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetNominalValue(symbol string, futuresAccountType int, side int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["symbol"] = symbol
	params["side"] = side
	byte, err := c.Get("/Server/api/v2/Positions/getNominalValue", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) SetLeverage(symbol string, futuresAccountType int, leverage int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["symbol"] = symbol
	params["leverage"] = leverage
	byte, err := c.Post("/Server/api/v2/setting/setLeverage", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) SetPositionsMode(symbol string, futuresAccountType int, positionsMode int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["symbol"] = symbol
	params["positionsMode"] = positionsMode
	byte, err := c.Post("/Server/api/v2/setting/setPositionsMode", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) SetMarginMode(symbol string, futuresAccountType int, marginMode int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["symbol"] = symbol
	params["marginMode"] = marginMode
	byte, err := c.Post("/Server/api/v2/setting/setMarginMode", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

func (c *Client) GetPositionSetting(symbol string, futuresAccountType int, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["symbol"] = symbol
	byte, err := c.Get("/Server/api/v2/setting/get", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

//设置自动追加保证金
func (c *Client) SetAutoAppendMargin(symbol string, futuresAccountType int, maxAppendAmount float64, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["symbol"] = symbol
	params["maxAppendAmount"] = maxAppendAmount
	byte, err := c.Post("/Server/api/v2/Positions/setAutoAppendMargin", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}

//设置保证金使用顺序
//marginCoins: "eth,usdt"
func (c *Client) SetMarginCoins(symbol string, futuresAccountType int, marginCoins string, handler handler.Handler) {
	params := make(map[string]interface{})
	params["futuresAccountType"] = futuresAccountType
	params["symbol"] = symbol
	params["marginCoins"] = marginCoins
	byte, err := c.Post("/Server/api/v2/Positions/setMarginCoins", params)
	if err != nil {
		return
	}

	if handler != nil {
		handler(*byte)
	}
}
