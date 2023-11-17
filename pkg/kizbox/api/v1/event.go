package v1

import (
	"context"
	"net/http"
)

type ApiEvent struct {
	cl Client
}

func NewApiEvent(cl Client) *ApiEvent {
	return &ApiEvent{cl: cl}
}

func (receiver *ApiEvent) Register(ctx context.Context, v *EventRegister) (*http.Response, error) {
	return receiver.cl.DoParams(ctx, WithMethod(http.MethodPost), WithPath("/events/register"), WithUnmarshalBody(v))
}

func (receiver *ApiEvent) Fetch(ctx context.Context, eventRegister EventRegister, v *[]map[string]interface{}) (*http.Response, error) {
	return receiver.cl.DoParams(ctx, WithMethod(http.MethodPost), WithPath("/events/%s/fetch", eventRegister.ID), WithUnmarshalBody(v))
}

func (receiver *ApiEvent) Unregister(ctx context.Context, listenerID string) (*http.Response, error) {
	return receiver.cl.DoParams(ctx, WithMethod(http.MethodPost), WithPath("/events/%s/unregister", listenerID))
}
