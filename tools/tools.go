// Copyright (c) 2026 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

// Package tools provides a set of example tools that can be registered with the MCP server.
// These tools demonstrate how to define tool metadata, input schemas, and handlers.
// You can use these as a starting point for creating your own custom tools.
package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tiagomelo/go-mcp-server/server"
)

// RegisterDefaultTools registers the default tools
// (hello_world, health_check, latency_percentiles) with the provided server.
func RegisterDefaultTools(s *server.Server) {
	s.RegisterTool(
		server.ToolDefinition{
			Name:        "hello_world",
			Description: "Returns a hello message for the provided name.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"name": map[string]any{
						"type":        "string",
						"description": "Optional name to greet.",
					},
				},
			},
		},
		func(ctx context.Context, raw json.RawMessage) (any, error) {
			var args HelloArgs
			if len(raw) > 0 {
				if err := json.Unmarshal(raw, &args); err != nil {
					return nil, fmt.Errorf("decoding arguments: %w", err)
				}
			}
			return Hello(args)
		},
	)

	s.RegisterTool(
		server.ToolDefinition{
			Name:        "health_check",
			Description: "Performs an HTTP GET against a URL and returns status code and latency.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"url": map[string]any{
						"type":        "string",
						"description": "The URL to check.",
					},
					"timeout_ms": map[string]any{
						"type":        "integer",
						"description": "Optional timeout in milliseconds.",
					},
				},
				"required": []string{"url"},
			},
		},
		func(ctx context.Context, raw json.RawMessage) (any, error) {
			var args HealthCheckArgs
			if err := json.Unmarshal(raw, &args); err != nil {
				return nil, fmt.Errorf("decoding arguments: %w", err)
			}
			return HealthCheck(ctx, args)
		},
	)

	s.RegisterTool(
		server.ToolDefinition{
			Name:        "latency_percentiles",
			Description: "Computes min, p50, p95, p99, max and average for a list of numeric values.",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"values": map[string]any{
						"type":        "array",
						"description": "Numeric values, typically latency measurements in milliseconds.",
						"items": map[string]any{
							"type": "number",
						},
					},
				},
				"required": []string{"values"},
			},
		},
		func(ctx context.Context, raw json.RawMessage) (any, error) {
			var args PercentilesArgs
			if err := json.Unmarshal(raw, &args); err != nil {
				return nil, fmt.Errorf("decoding arguments: %w", err)
			}
			return Percentiles(args)
		},
	)
}
