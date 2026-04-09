// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package tools

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// HealthCheckArgs defines the arguments for the HealthCheck tool.
type HealthCheckArgs struct {
	URL       string `json:"url"`
	TimeoutMS int    `json:"timeout_ms,omitempty"`
}

// HealthCheckResult defines the result of the HealthCheck tool.
type HealthCheckResult struct {
	URL        string `json:"url"`
	StatusCode int    `json:"status_code"`
	LatencyMS  int64  `json:"latency_ms"`
	OK         bool   `json:"ok"`
}

// HealthCheck performs a health check on the specified URL and returns
// the status code, latency, and whether the check was successful.
func HealthCheck(ctx context.Context, args HealthCheckArgs) (HealthCheckResult, error) {
	if args.URL == "" {
		return HealthCheckResult{}, errors.New("url is required")
	}

	timeout := 3 * time.Second
	if args.TimeoutMS > 0 {
		timeout = time.Duration(args.TimeoutMS) * time.Millisecond
	}

	client := newHTTPClient(timeout)

	req, err := requestBuilderProvider.NewRequestWithContext(ctx, http.MethodGet, args.URL, nil)
	if err != nil {
		return HealthCheckResult{}, errors.WithMessage(err, "failed to create request")
	}

	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return HealthCheckResult{}, errors.WithMessage(err, "failed to perform request")
	}
	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	latency := time.Since(start).Milliseconds()

	return HealthCheckResult{
		URL:        args.URL,
		StatusCode: resp.StatusCode,
		LatencyMS:  latency,
		OK:         resp.StatusCode >= 200 && resp.StatusCode < 300,
	}, nil
}
