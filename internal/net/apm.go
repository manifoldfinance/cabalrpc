package net

import (
	"github.com/opentracing/opentracing-go"
	"net/http"
)

func NewApmHttpClient(core RawClient) RawClient {
	return apmHttpClient{core: core}
}

type apmHttpClient struct {
	core RawClient
}

func (a apmHttpClient) Do(req *http.Request) (*http.Response, error) {
	span, _ := opentracing.StartSpanFromContext(req.Context(), "http-post")
	span = span.SetTag("url", req.URL.String())
	defer span.Finish()
	return a.core.Do(req)
}
