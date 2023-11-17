package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type ApiExecution struct {
	cl Client
}

func NewApiExecution(cl Client) *ApiExecution {
	return &ApiExecution{cl: cl}
}

type Apply struct {
	ID string `json:"id"`
}

func (a *ApiExecution) Apply(ctx context.Context, execute Execute, s *Apply) (*http.Response, error) {
	// patch resource to ensure it is acceptable
	for i, action := range execute.Actions {
		for j, command := range action.Commands {
			if command.Parameters == nil {
				// null parameters is not supported, an empty array must be provided
				execute.Actions[i].Commands[j].Parameters = make([]interface{}, 0)
			}
		}
	}

	ex, err := json.Marshal(execute)
	if err != nil {
		return nil, err
	}

	body, res, err := a.cl.Do(ctx, http.MethodPost, "/exec/apply", bytes.NewReader(ex))
	if err != nil {
		return res, err
	}

	if s != nil {
		return res, json.Unmarshal(body, s)
	}

	return res, nil
}

func (a *ApiExecution) Current(ctx context.Context, s *[]Execution) (*http.Response, error) {
	body, res, err := a.cl.Do(ctx, http.MethodGet, "/exec/current", nil)
	if err != nil {
		return res, err
	}

	if s == nil {
		return res, nil
	}

	return res, json.Unmarshal(body, s)
}

func (a *ApiExecution) Get(ctx context.Context, execID string, s *Execution) (*http.Response, error) {
	path, err := url.JoinPath("/exec", execID)
	if err != nil {
		return nil, err
	}

	body, res, err := a.cl.Do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return res, err
	}

	if s == nil {
		return res, nil
	}

	return res, json.Unmarshal(body, s)
}
