package services

import (
	"github.com/go-resty/resty/v2"
)

type GotHttpClient struct {
	client *resty.Client
}

func NewGotHttpClient() *GotHttpClient {
	return &GotHttpClient{
		client: resty.New(),
	}
}

func (c *GotHttpClient) Get(url string) ([]byte, error) {
	resp, err := c.client.R().Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
