package v1

import (
	"context"
	"net/http"
	"net/url"
)

type ApiDevices struct {
	cl Client
}

func NewApiDevices(cl Client) *ApiDevices {
	return &ApiDevices{cl: cl}
}

func (receiver *ApiDevices) List(ctx context.Context, v *[]Device) (*http.Response, error) {
	return receiver.cl.Do(ctx, WithMethod(http.MethodGet), WithPath("/setup/devices"), WithUnmarshalBody(v))
}

func (receiver *ApiDevices) Get(ctx context.Context, deviceURL string, v *Device) (*http.Response, error) {
	return receiver.cl.Do(ctx, WithMethod(http.MethodGet), WithPath("/setup/devices/%s", url.PathEscape(deviceURL)), WithUnmarshalBody(v))
}

func (receiver *ApiDevices) States(ctx context.Context, deviceURL string, v *[]State) (*http.Response, error) {
	return receiver.cl.Do(ctx, WithMethod(http.MethodGet), WithPath("/setup/devices/%s/states", url.PathEscape(deviceURL)), WithUnmarshalBody(v))
}

func (receiver *ApiDevices) State(ctx context.Context, deviceURL string, stateName string, v *State) (*http.Response, error) {
	return receiver.cl.Do(ctx, WithMethod(http.MethodGet), WithPath("/setup/devices/%s/states/%s", url.PathEscape(deviceURL), stateName), WithUnmarshalBody(v))
}

func (receiver *ApiDevices) Controllables(ctx context.Context, controllableName string, v *[]string) (*http.Response, error) {
	return receiver.cl.Do(ctx, WithMethod(http.MethodGet), WithPath("/setup/devices/controllables/%s", url.PathEscape(controllableName)), WithUnmarshalBody(v))
}
