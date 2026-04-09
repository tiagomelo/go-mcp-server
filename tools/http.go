package tools

import (
	"context"
	"io"
	"net/http"
	"time"
)

// requestBuilder defines an interface for building HTTP requests,
// allowing for easier testing and abstraction.
type requestBuilder interface {
	NewRequestWithContext(ctx context.Context, method string, url string, body io.Reader) (*http.Request, error)
}

// defaultRequestBuilderProvider is the default implementation of the requestBuilder interface,
// using the standard library's http.NewRequestWithContext function.
type defaultRequestBuilderProvider struct{}

// NewRequestWithContext creates a new HTTP request with the given context, method, URL, and body.
func (p *defaultRequestBuilderProvider) NewRequestWithContext(ctx context.Context, method string, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, body)
}

// requestBuilderProvider is the global variable that holds
// the current request builder implementation.
var requestBuilderProvider requestBuilder = &defaultRequestBuilderProvider{}

// httpClient defines an interface for making HTTP requests,
// allowing for easier testing and abstraction.
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// httpClientFactory defines a function type for creating new HTTP clients with a specified timeout.
type httpClientFactory func(timeout time.Duration) httpClient

// newHTTPClient is the default implementation of the httpClientFactory,
// creating a new http.Client with the specified timeout.
var newHTTPClient httpClientFactory = func(timeout time.Duration) httpClient {
	return &http.Client{Timeout: timeout}
}
