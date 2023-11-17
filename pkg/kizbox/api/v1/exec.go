package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
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

func (receiver *ApiExecution) Apply(ctx context.Context, execute Execute, s *Apply) (*http.Response, error) {
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

	return receiver.cl.Do(ctx, WithMethod(http.MethodPost), WithPath("/exec/apply"), WithBody(bytes.NewReader(ex)), WithUnmarshalBody(s))
}

func (receiver *ApiExecution) Current(ctx context.Context, s *[]Execution) (*http.Response, error) {
	return receiver.cl.Do(ctx, WithMethod(http.MethodGet), WithPath("/exec/current"), WithUnmarshalBody(s))
}

func (receiver *ApiExecution) Get(ctx context.Context, execID string, s *Execution) (*http.Response, error) {
	return receiver.cl.Do(ctx, WithMethod(http.MethodGet), WithPath("/exec/%s", execID), WithUnmarshalBody(s))
}
