package v1

import (
	"context"
	"net/http"
)

type ApiVersion struct {
	cl Client
}

func NewApiVersion(cl Client) *ApiVersion {
	return &ApiVersion{cl: cl}
}

func (a *ApiVersion) Get(ctx context.Context, v *Version) (*http.Response, error) {
	return a.cl.Do(ctx, WithMethod(http.MethodGet), WithPath("/apiVersion"), WithUnmarshalBody(v))
}
