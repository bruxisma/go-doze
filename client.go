package doze

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type Client resty.Client

func (client *Client) Request(ctx context.Context, result interface{}) *Request {
	request := (*resty.Client)(client).NewRequest().
		ExpectContentType("application/json").
		SetContext(ctx).
		SetResult(result)
	return (*Request)(request)
}

func (client *Client) ToResty() *resty.Client {
	return (*resty.Client)(client)
}
