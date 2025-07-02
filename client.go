package dingtalk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/taadis/dingtalk/sign"
)

const (
	defaultDialTimeout = 5 * time.Second
	defaultKeepAlive   = 5 * time.Second
)

type Client struct {
	//
	httpClient *http.Client
	//
	options *Options
}

func NewClient(opts ...Option) *Client {
	return &Client{
		httpClient: newHttpClient(),
		options:    newOptions(opts...),
	}
}

func newHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   defaultDialTimeout,
				KeepAlive: defaultKeepAlive,
			}).DialContext,
		},
		Timeout: defaultDialTimeout,
	}
	return client
}

func (c *Client) do(ctx context.Context, callMethod string, endPoint string, header map[string]string, body []byte) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, callMethod, endPoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, fmt.Errorf("response is nil, please check it")
	}

	return response, nil
}

func (c *Client) Send(ctx context.Context, msg Messager) error {
	return c.send(ctx, msg)
}

func (c *Client) send(ctx context.Context, msg Messager) error {
	var (
		uri  string
		resp *http.Response
		err  error
	)
	value := url.Values{}
	value.Set("access_token", c.options.AccessToken)
	if c.options.secret != "" {
		t := time.Now().UnixMilli()
		sign := sign.GenerateSign(t, c.options.secret)
		value.Set("timestamp", fmt.Sprintf("%d", t))
		value.Set("sign", sign)
	}
	uri = Webhook + value.Encode()
	header := map[string]string{
		"Content-type": "application/json",
	}
	reqBody, err := msg.Marshal()
	if err != nil {
		return err
	}
	resp, err = c.do(ctx, http.MethodPost, uri, header, reqBody)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("send msg err. http code: %d, token: %s, msg: %s", resp.StatusCode, c.options.AccessToken, reqBody)
	}
	rspBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var rsp Response
	err = json.Unmarshal(rspBody, &rsp)
	if err != nil {
		return err
	}
	if rsp.ErrCode != 0 {
		return fmt.Errorf("%d - %s", rsp.ErrCode, rsp.ErrMsg)
	}

	return nil
}
