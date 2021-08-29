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

func (client *Client) SetAuthToken(token string) {
	client.ToResty().SetAuthToken(token)
}

func (client *Client) SetAuthScheme(scheme string) {
	client.ToResty().SetAuthScheme(scheme)
}

func (client *Client) SetUserAgent(agent string) {
	client.ToResty().SetHeader("User-Agent", agent)
}

func (client *Client) SetBaseURL(url string) {
	client.ToResty().SetHostURL(url)
}

func (client *Client) SetDebug() {
	client.ToResty().SetDebug(true)
}
