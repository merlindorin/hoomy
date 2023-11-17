package client

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

type ApiClient struct {
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

func NewClient(addr, apiKey string) *ApiClient {
	cl := http.DefaultClient
	cl.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	apiClient := &ApiClient{
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

func (cl *ApiClient) DoParams(ctx context.Context, params ...v1.WithParam) (res *http.Response, err error) {
	p := &v1.Params{
		Method:           "GET",
		Path:             "/",
		Body:             nil,
		ResponseHandlers: nil,
	}

	for _, param := range params {
		param.Apply(p)
	}

	u, err := url.JoinPath(fmt.Sprintf("https://%s", cl.addr), apiPath, p.Path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, p.Method, u, p.Body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err = cl.cl.Do(req)
	if err != nil {
		return res, err
	}

	if res.StatusCode >= 400 {
		return res, fmt.Errorf("unexcpected response status code %d: %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	for _, responseHandler := range p.ResponseHandlers {
		err = responseHandler(res)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (cl *ApiClient) Do(ctx context.Context, method, path string, b io.Reader) (res *http.Response, err error) {
	return cl.DoParams(ctx, v1.WithMethod(method), v1.WithPath(path), v1.WithBody(b))
}
