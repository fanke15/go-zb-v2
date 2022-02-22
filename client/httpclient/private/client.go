package private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fanke15/go-zb-v2/client"
	"github.com/fanke15/go-zb-v2/client/httpclient"
	"net/http"
	"strings"

	. "github.com/fanke15/go-zb-v2/log"
)

type Client struct {
	*httpclient.HttpClient

	ApiKey     string
	SecretKey  string
	FutureType string
}

func NewPrivateClient(endpoint, futureType, apiKey, secretKey string) *Client {
	var client Client
	client.HttpClient = httpclient.NewHttpClient(endpoint)
	client.ApiKey = apiKey
	client.SecretKey = secretKey
	client.FutureType = futureType

	if client.FutureType != "" {
		client.FutureType = "/" + client.FutureType
	}

	return &client
}

func (c *Client) Get(path string, params map[string]interface{}) (*json.RawMessage, error) {
	path = c.FutureType + path

	url := c.Endpoint + path
	if params != nil {
		var urlSuffix []string
		for k, v := range params {
			urlSuffix = append(urlSuffix, k+"="+fmt.Sprint(v))
		}
		join := strings.Join(urlSuffix, "&")
		url = url + "?" + join
	}
	Log.Info().Msg(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Log.Error().Msg(err.Error())
		return nil, err
	}

	ts := client.GetISOTime()
	sign := client.Sign2(ts, "GET", path, params, c.SecretKey)
	req.Header.Set("ZB-APIKEY", c.ApiKey)
	req.Header.Set("ZB-TIMESTAMP", ts)
	req.Header.Set("ZB-SIGN", sign)
	req.Header.Set("ZB-LAN", "cn")

	resp, err := c.Client.Do(req)
	if err != nil {
		Log.Error().Msg(err.Error())
		return nil, err
	}
	return c.HandleResponse(resp)
}

func (c *Client) Post(path string, params map[string]interface{}) (*json.RawMessage, error) {
	path = c.FutureType + path

	url := c.Endpoint + path
	Log.Info().Msg(url)

	marshal, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshal))
	if err != nil {
		Log.Error().Msg(err.Error())
		return nil, err
	}

	ts := client.GetISOTime()
	sign := client.Sign2(ts, "POST", path, params, c.SecretKey)
	req.Header.Set("ZB-APIKEY", c.ApiKey)
	req.Header.Set("ZB-TIMESTAMP", ts)
	req.Header.Set("ZB-SIGN", sign)
	req.Header.Set("ZB-LAN", "cn")

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		Log.Error().Msg(err.Error())
		return nil, err
	}
	return c.HandleResponse(resp)
}
