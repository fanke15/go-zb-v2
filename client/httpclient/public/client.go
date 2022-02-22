package public

import (
	"github.com/fanke15/go-zb-v2/client/httpclient"
)

type Client struct {
	*httpclient.HttpClient
}

func NewPublicClient(endpoint string) *Client {
	var client Client
	client.HttpClient = httpclient.NewHttpClient(endpoint)

	return &client
}
