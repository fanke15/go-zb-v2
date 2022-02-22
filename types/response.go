package types

import "encoding/json"

type WSResponse struct {
	Channel string           `json:"channel"`
	Data    *json.RawMessage `json:"data,omitempty"`

	Code int    `json:"errorCode,omitempty"`
	Msg  string `json:"errorMsg,omitempty"`
}

type HttpResponse struct {
	Code   int              `json:"code"`
	Desc   string           `json:"desc"`
	CnDesc string           `json:"cnDesc,omitempty"`
	EnDesc string           `json:"enDesc,omitempty"`
	Data   *json.RawMessage `json:"data,omitempty"`
}
