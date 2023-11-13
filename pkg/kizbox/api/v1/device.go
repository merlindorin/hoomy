package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ApiDevices struct {
	cl Client
}

func NewApiDevices(cl Client) *ApiDevices {
	return &ApiDevices{cl: cl}
}

func (a *ApiDevices) List(ctx context.Context, devices *[]Device) (*http.Response, error) {
	body, res, err := a.cl.Do(ctx, http.MethodGet, "/setup/devices", nil)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(body, &devices)
}

func (a *ApiDevices) Get(ctx context.Context, deviceURL string, d *Device) (*http.Response, error) {
	body, res, err := a.cl.Do(ctx, http.MethodGet, fmt.Sprintf("/setup/devices/%s", url.PathEscape(deviceURL)), nil)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(body, d)
}

func (a *ApiDevices) States(ctx context.Context, deviceURL string, s *[]State) (*http.Response, error) {
	body, res, err := a.cl.Do(ctx, http.MethodGet, fmt.Sprintf("/setup/devices/%s/states", url.PathEscape(deviceURL)), nil)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(body, s)
}

func (a *ApiDevices) State(ctx context.Context, deviceURL string, stateName string, s *State) (*http.Response, error) {
	body, res, err := a.cl.Do(ctx, http.MethodGet, fmt.Sprintf("/setup/devices/%s/states/%s", url.PathEscape(deviceURL), stateName), nil)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(body, s)
}

func (a *ApiDevices) Controllables(ctx context.Context, controllableName string, s *[]string) (*http.Response, error) {
	body, res, err := a.cl.Do(ctx, http.MethodGet, fmt.Sprintf("/setup/devices/controllables/%s", url.PathEscape(controllableName)), nil)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(body, s)
}
