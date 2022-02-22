package private

import (
	"encoding/json"
	"github.com/fanke15/go-zb-v2/client"
	"github.com/fanke15/go-zb-v2/client/websocket"
	"github.com/fanke15/go-zb-v2/constant"
	. "github.com/fanke15/go-zb-v2/log"
)

/**
用于websocket私有数据，需要用户login才能订阅
*/
const (
	//test
	//ApiKey    = "a55caded-eef9-426b-af7c-faf53c78e2ae"
	//SecretKey = "3b5b1e07-383e-42dc-ae85-90449241c0dc" //网站拿到的原始的没有经过sha1加密的SecretKey，

	//online test
	ApiKey    = "3576549f-e1ce-4317-b83e-48f192cf9e23"
	SecretKey = "6ee95f06-0365-4e09-a9db-b38d2d7111b0" //网站拿到的原始的没有经过sha1加密的SecretKey，
)

//每个订阅都需要的参数
type baseSubscriber struct {
	Action             string `json:"action"`
	Channel            string `json:"channel"`
	FuturesAccountType int    `json:"futuresAccountType,omitempty"`
}

func subscribe(c *websocket.WSClient, sub interface{}) bool {
	data, err := json.Marshal(sub)
	if err != nil {
		Log.Error().Msg(err.Error())
		return false
	}

	Log.Info().Msg(string(data))
	c.Write(data)

	return true
}

type login struct {
	Action string `json:"action"`

	APIKey    string `json:"ZB-APIKEY"`
	Timestamp string `json:"ZB-TIMESTAMP"`
	Sign      string `json:"ZB-SIGN"`
}

func Login(c *websocket.WSClient) {
	sub := new(login)
	sub.Action = constant.LOGIN
	sub.APIKey = ApiKey
	sub.Timestamp = client.GetISOTime()
	//sub.Sign = client.Sign(sub.Timestamp, "GET", "login", nil, SecretKey)
	sub.Sign = client.Sign2(sub.Timestamp, "GET", "login", nil, SecretKey)

	subscribe(c, sub)
}
