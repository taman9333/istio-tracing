package customhttpclient

import (
	"context"
	"io"
	"log"
	"net/http"
	"sync"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

var once sync.Once
var client http.Client

// Req make a custom http request
func Req(ctx context.Context, method string, url string, body io.Reader) (*http.Response, error) {
	once.Do(func() {
		client = http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	})

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return client.Do(req)
}
