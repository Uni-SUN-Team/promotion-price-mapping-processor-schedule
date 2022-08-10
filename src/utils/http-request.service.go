package utils

import (
	"bytes"
	"net/http"
	"time"
	"unisun/api/class-room-price-mapping-processor-schedule/src/constants"
)

type HTTPRequestAdapter struct {
}

func New() *HTTPRequestAdapter {
	return &HTTPRequestAdapter{}
}

func (*HTTPRequestAdapter) HTTPRequest(url string, method string, payload []byte) (*http.Response, error) {
	var request *http.Request
	var err error
	var body *bytes.Buffer
	timeout := time.Duration(5 * time.Second)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}
	switch method {
	case constants.GET:
		body = bytes.NewBuffer(nil)
	case constants.POST:
		body = bytes.NewBuffer(payload)
	case constants.PUT:
		body = bytes.NewBuffer(payload)
	case constants.DELETE:
		body = bytes.NewBuffer(nil)
	default:
		body = bytes.NewBuffer(nil)
	}
	if err != nil {
		return nil, err
	}
	request, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
