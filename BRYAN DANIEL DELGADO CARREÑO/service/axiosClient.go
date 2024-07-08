package services

import (
	"io/ioutil"
	"net/http"
)

type AxiosHttpClient struct{}

func (c *AxiosHttpClient) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
