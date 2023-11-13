package kizbox

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"

	v1 "github.com/merlindorin/hoomy/pkg/kizbox/api/v1"
)

const (
	apiPath = "/enduser-mobile-web/1/enduserAPI"
)

type Client struct {
	addr string
	cl   *http.Client

	apiKey string
	V1     V1
}

type V1 struct {
	Version   *v1.ApiVersion
	Gateway   *v1.ApiGateway
	Devices   *v1.ApiDevices
	Execution *v1.ApiExecution
	Event     *v1.ApiEvent
}

func NewClient(addr, apiKey string) *Client {
	cl := http.DefaultClient
	cl.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	apiClient := &Client{
		addr:   addr,
		cl:     cl,
		apiKey: apiKey,
	}

	apiClient.V1 = V1{
		Version:   v1.NewApiVersion(apiClient),
		Gateway:   v1.NewApiGateway(apiClient),
		Devices:   v1.NewApiDevices(apiClient),
		Execution: v1.NewApiExecution(apiClient),
		Event:     v1.NewApiEvent(apiClient),
	}

	return apiClient
}

func (cl *Client) Do(ctx context.Context, method, path string, b io.Reader) (body []byte, res *http.Response, err error) {
	u, err := url.JoinPath(fmt.Sprintf("https://%s", cl.addr), apiPath, path)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, u, b)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err = cl.cl.Do(req)
	if err != nil {
		return nil, res, err
	}

	if res.StatusCode >= 400 {
		return nil, res, fmt.Errorf("unexcpected response status code %d: %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	return body, res, nil
}
