package doze

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type Request resty.Request

func (r *Request) Context() context.Context {
	return r.cast().Context()
}

func (r *Request) SetPathParameter(parameter, value string) *Request {
	return (*Request)(r.cast().SetPathParam(parameter, value))
}

func (r *Request) SetQueryParameter(parameter, value string) *Request {
	return (*Request)(r.cast().SetQueryParam(parameter, value))
}

func (r *Request) SetHeader(header, value string) *Request {
	return (*Request)(r.cast().SetHeader(header, value))
}

func (r *Request) SetBody(body interface{}) *Request {
	return (*Request)(r.cast().SetBody(body))
}

func (r *Request) SetContentLength() *Request {
	return (*Request)(r.cast().SetContentLength(true))
}

func (request *Request) Get(url string) (interface{}, error) {
	return checkResponse(request.cast().Get(url))
}

func (request *Request) Patch(url string) (interface{}, error) {
	return checkResponse(request.cast().Patch(url))
}

func (request *Request) Post(url string) (interface{}, error) {
	return checkResponse(request.cast().Post(url))
}

func (request *Request) Put(url string) (interface{}, error) {
	return checkResponse(request.cast().Put(url))
}

func (request *Request) Delete(url string) (interface{}, error) {
	return checkResponse(request.cast().Delete(url))
}

// Cast the request to the underlying resty.Request type
func (request *Request) cast() *resty.Request {
	return (*resty.Request)(request)
}

func checkResponse(response *resty.Response, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, response.Error().(error)
	}
	return response.Result(), nil
}
