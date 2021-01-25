package net

import (
	"bytes"
	"github.com/manifoldfinance/cabalrpc/internal/config"
	"io/ioutil"
	"net/http"
)

type RawClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const contentTypeJson = "application/json"

type RpcClient interface {
	Call([]byte) ([]byte, error)
}

func NewRpcClient(config config.Config) RpcClient {
	var httpClient RawClient
	httpClient = &http.Client{}
	if config.ApmEnabled {
		httpClient = NewApmHttpClient(&http.Client{})
	}
	return rpcClient{
		config: config,
		client: httpClient,
	}
}

type rpcClient struct {
	config config.Config
	client RawClient
}

func (r rpcClient) Call(request []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", r.config.RpcUrl, bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentTypeJson)
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)

}
