package v1

import (
	"context"
	"io"
	"net/http"
)

type Client interface {
	Do(ctx context.Context, method, path string, b io.Reader) (body []byte, res *http.Response, err error)
}
