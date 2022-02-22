package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fanke15/go-zb-v2/constant"
	. "github.com/fanke15/go-zb-v2/log"
	"github.com/fanke15/go-zb-v2/types"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct {
	Client *http.Client

	Endpoint string
}

func NewHttpClient(endpoint string) *HttpClient {
	var c HttpClient
	c.Endpoint = endpoint
	c.Client = &http.Client{
		Timeout: 10 * time.Second,
	}
	return &c
}

func (c *HttpClient) Get(path string, params map[string]interface{}) (*json.RawMessage, error) {
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

	resp, err := c.Client.Get(url)
	if err != nil {
		Log.Error().Msg(err.Error())
		return nil, err
	}
	return c.HandleResponse(resp)
}

func (c *HttpClient) HandleResponse(resp *http.Response) (*json.RawMessage, error) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.Error().Msg(err.Error())
		return nil, err
	}
	Log.Debug().Msg(string(body))

	var response types.HttpResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		Log.Error().Str("body", string(body)).Msg(err.Error())
		return nil, err
	}

	if response.Code != constant.SUCCESS {
		Log.Error().Int("code", response.Code).Str("desc", response.Desc).Msg("")
		return nil, errors.New(response.Desc)
	}

	return response.Data, nil
}
