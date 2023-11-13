package v1

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiVersion struct {
	cl Client
}

func NewApiVersion(cl Client) *ApiVersion {
	return &ApiVersion{cl: cl}
}

func (a *ApiVersion) Get(ctx context.Context, v *Version) (*http.Response, error) {
	body, res, err := a.cl.Do(ctx, http.MethodGet, "/apiVersion", nil)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(body, v)
}
