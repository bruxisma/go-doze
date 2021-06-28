package doze

import "github.com/go-resty/resty/v2"

type Request resty.Request

// Cast the request to the underlying resty.Request type
func (request *Request) Cast() *resty.Request {
	return (*resty.Request)(request)
}

func (r *Request) SetPathParameter(parameter, value string) *Request {
	return (*Request)(r.Cast().SetPathParam(parameter, value))
}

func (r *Request) SetQueryParameter(parameter, value string) *Request {
	return (*Request)(r.Cast().SetQueryParam(parameter, value))
}

func (r *Request) SetBody(body interface{}) *Request {
	return (*Request)(r.Cast().SetBody(body))
}

func (request *Request) Get(url string) (interface{}, error) {
	return checkResponse(request.Cast().Get(url))
}

func (request *Request) Patch(url string) (interface{}, error) {
	return checkResponse(request.Cast().Patch(url))
}

func (request *Request) Post(url string) (interface{}, error) {
	return checkResponse(request.Cast().Post(url))
}

func (request *Request) Put(url string) (interface{}, error) {
	return checkResponse(request.Cast().Put(url))
}

func (request *Request) Delete(url string) (interface{}, error) {
	return checkResponse(request.Cast().Delete(url))
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
