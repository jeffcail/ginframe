package httprequest

import (
	"bytes"
	"github.com/jeffcail/gorequest"
)

type HttpRequest struct{}

// Get http get
func (hr *HttpRequest) Get(url string, header map[string]string, params map[string]interface{}) ([]byte, error) {
	return gorequest.Get(url, header, params)
}

// Post http post
func (hr *HttpRequest) Post(url string, header map[string]string, params map[string]interface{}) ([]byte, error) {
	return gorequest.Post(url, header, params)
}

// PostMultipart http post multipart
func (hr *HttpRequest) PostMultipart(url string, header map[string]string, payload *bytes.Buffer) ([]byte, error) {
	return gorequest.PostMultipart(url, header, payload)
}
