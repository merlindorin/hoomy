package v1

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiGateway struct {
	cl Client
}

func NewApiGateway(cl Client) *ApiGateway {
	return &ApiGateway{cl: cl}
}

func (a *ApiGateway) List(ctx context.Context, c chan *Gateway) (*http.Response, error) {
	body, res, err := a.cl.Do(ctx, http.MethodGet, "/setup/gateways", nil)
	if err != nil {
		return res, err
	}

	var gateways []*Gateway

	err = json.Unmarshal(body, &gateways)
	if err != nil {
		return res, err
	}

	for _, gateway := range gateways {
		c <- gateway
	}

	return res, nil
}
