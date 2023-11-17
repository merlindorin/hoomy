package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiEvent struct {
	cl Client
}

func NewApiEvent(cl Client) *ApiEvent {
	return &ApiEvent{cl: cl}
}

func (a *ApiEvent) Register(ctx context.Context, e *EventRegister) (*http.Response, error) {
	body, res, err := a.cl.Do(ctx, http.MethodPost, "/events/register", nil)
	if err != nil {
		return res, err
	}

	if e == nil {
		return res, nil
	}

	return res, json.Unmarshal(body, e)
}

func (a *ApiEvent) Fetch(ctx context.Context, eventRegister EventRegister, e *[]map[string]interface{}) (*http.Response, error) {
	body, res, err := a.cl.Do(ctx, http.MethodPost, fmt.Sprintf("/events/%s/fetch", eventRegister.ID), nil)
	if err != nil {
		return res, err
	}

	if e == nil {
		return res, nil
	}

	return res, json.Unmarshal(body, e)
}

func (a *ApiEvent) Unregister(ctx context.Context, listenerID string) (*http.Response, error) {
	_, res, err := a.cl.Do(ctx, http.MethodPost, fmt.Sprintf("/events/%s/unregister", listenerID), nil)
	return res, err
}
