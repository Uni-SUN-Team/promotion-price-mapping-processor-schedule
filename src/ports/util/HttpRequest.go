package util

import "net/http"

type HttpRequestPort interface {
	HTTPRequest(url string, method string, payload []byte) (*http.Response, error)
}
