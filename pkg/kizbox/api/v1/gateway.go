package v1

import (
	"context"
	"net/http"
)

type ApiGateway struct {
	cl Client
}

func NewApiGateway(cl Client) *ApiGateway {
	return &ApiGateway{cl: cl}
}

func (a *ApiGateway) List(ctx context.Context, v *[]Gateway) (*http.Response, error) {
	return a.cl.Do(ctx, WithMethod(http.MethodGet), WithPath("/setup/gateways"), WithUnmarshalBody(v))
}
