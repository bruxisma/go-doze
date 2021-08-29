package doze

import (
	"context"
	"time"

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

func NewClient() *Client {
	client := resty.New().
		SetContentLength(true).
		SetTimeout(time.Duration(1 * time.Minute))
	return (*Client)(client)
}

func (client *Client) ToResty() *resty.Client {
	return (*resty.Client)(client)
}

func (client *Client) SetAuthToken(token string) *Client {
	client.ToResty().SetAuthToken(token)
	return client
}

func (client *Client) SetAuthScheme(scheme string) *Client {
	client.ToResty().SetAuthScheme(scheme)
	return client
}

func (client *Client) SetBaseURL(url string) *Client {
	client.ToResty().SetHostURL(url)
	return client
}

func (client *Client) SetDebug() *Client {
	client.ToResty().SetDebug(true)
	return client
}

func (client *Client) SetError(error interface{}) *Client {
	client.ToResty().SetError(error)
	return client
}

func (client *Client) SetUserAgent(agent string) *Client {
	client.ToResty().SetHeader("User-Agent", agent)
	return client
}
