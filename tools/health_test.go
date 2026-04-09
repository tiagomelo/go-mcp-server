// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package tools

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {

	originalRequestBuilderProvider := requestBuilderProvider
	originalNewHTTPClient := newHTTPClient

	t.Run("valid URL", func(t *testing.T) {
		defer func() {
			requestBuilderProvider = originalRequestBuilderProvider
			newHTTPClient = originalNewHTTPClient
		}()

		requestBuilderProvider = &mockRequestBuilder{}
		newHTTPClient = func(timeout time.Duration) httpClient {
			return &mockHTTPClient{
				resp: &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader("")),
				},
				err: nil,
			}
		}

		expectedResult := HealthCheckResult{
			URL:        "https://www.google.com",
			StatusCode: http.StatusOK,
			LatencyMS:  0, // latency is not tested here
			OK:         true,
		}

		result, err := HealthCheck(context.TODO(), HealthCheckArgs{URL: "https://www.google.com", TimeoutMS: 1000})
		require.NoError(t, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("missing url", func(t *testing.T) {
		result, err := HealthCheck(context.TODO(), HealthCheckArgs{})
		require.Error(t, err)
		require.Equal(t, HealthCheckResult{}, result)
	})

	t.Run("create request error", func(t *testing.T) {
		defer func() {
			requestBuilderProvider = originalRequestBuilderProvider
		}()

		requestBuilderProvider = &mockRequestBuilder{err: errors.New("request creation error")}

		result, err := HealthCheck(context.TODO(), HealthCheckArgs{URL: "https://www.google.com"})
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to create request")
		require.Equal(t, HealthCheckResult{}, result)
	})

	t.Run("perform request error", func(t *testing.T) {
		defer func() {
			newHTTPClient = originalNewHTTPClient
		}()

		newHTTPClient = func(timeout time.Duration) httpClient {
			return &mockHTTPClient{
				resp: nil,
				err:  errors.New("request error"),
			}
		}

		result, err := HealthCheck(context.TODO(), HealthCheckArgs{URL: "https://www.google.com"})
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to perform request")
		require.Equal(t, HealthCheckResult{}, result)
	})
}

type mockRequestBuilder struct {
	err error
}

func (m *mockRequestBuilder) NewRequestWithContext(ctx context.Context, method string, url string, body io.Reader) (*http.Request, error) {
	return new(http.Request), m.err
}

type mockHTTPClient struct {
	resp *http.Response
	err  error
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.resp, m.err
}
