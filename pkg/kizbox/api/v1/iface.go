package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

type Client interface {
	Do(ctx context.Context, method, path string, b io.Reader) (res *http.Response, err error)
	DoParams(ctx context.Context, params ...WithParam) (res *http.Response, err error)
}

type ResponseHandlerFunc func(response *http.Response) error

type Params struct {
	Method string
	Path   string
	Body   io.Reader

	ResponseHandlers []ResponseHandlerFunc
}

type WithParam func(params *Params)

func (p WithParam) Apply(params *Params) {
	p(params)
}

func WithMethod(method string) WithParam {
	return func(params *Params) {
		params.Method = method
	}
}

func WithPath(path string, a ...any) WithParam {
	return func(params *Params) {
		params.Path = fmt.Sprintf(path, a...)
	}
}

func WithBody(b io.Reader) WithParam {
	return func(params *Params) {
		params.Body = b
	}
}

func WithResponseHandler(f ResponseHandlerFunc) WithParam {
	return func(params *Params) {
		params.ResponseHandlers = append(params.ResponseHandlers, f)
	}
}

func WithUnmarshalBody(v any) WithParam {
	return func(params *Params) {
		params.ResponseHandlers = append(params.ResponseHandlers, func(response *http.Response) error {
			if v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
				return nil
			}

			body, err := io.ReadAll(response.Body)
			if err != nil {
				return err
			}

			err = json.Unmarshal(body, v)
			if err != nil {
				return err
			}

			response.Body = io.NopCloser(bytes.NewBuffer(body))
			return nil
		})
	}
}
