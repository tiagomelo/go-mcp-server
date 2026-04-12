# go-mcp-server

[![CI](https://github.com/tiagomelo/go-mcp-server/actions/workflows/ci.yml/badge.svg)](https://github.com/tiagomelo/go-mcp-server/actions/workflows/ci.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/tiagomelo/go-mcp-server.svg)](https://pkg.go.dev/github.com/tiagomelo/go-mcp-server)

A sample [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server implemented in [Go](https://go.dev/) from scratch, using [JSON-RPC 2.0](https://www.jsonrpc.org/specification) over stdio.

No third-party MCP libraries -- just the standard library and the MCP specification.

Article: https://tiagomelo.info/golang/mcp/2026/04/10/go-mcp-server.html

## Available tools

| Tool | Description |
|------|-------------|
| `hello_world` | Returns a greeting message for the provided name |
| `health_check` | Performs an HTTP GET against a URL and returns status code and latency |
| `latency_percentiles` | Computes min, p50, p95, p99, max and average for a list of numeric values |

## Running

```bash
make run
```

This keeps stdin open so you can paste JSON-RPC messages interactively.

## Manual testing

### 1. Initialize

```json
{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-06-18","capabilities":{},"clientInfo":{"name":"manual-test","version":"0.1.0"}}}
```

### 2. Confirm initialization

```json
{"jsonrpc":"2.0","method":"notifications/initialized"}
```

### 3. Ping

```json
{"jsonrpc":"2.0","id":2,"method":"ping"}
```

### 4. List tools

```json
{"jsonrpc":"2.0","id":3,"method":"tools/list"}
```

### 5. Call tools

**hello_world**

```json
{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"hello_world","arguments":{"name":"Tiago"}}}
```

**health_check**

```json
{"jsonrpc":"2.0","id":5,"method":"tools/call","params":{"name":"health_check","arguments":{"url":"https://httpbin.org/get","timeout_ms":5000}}}
```

**latency_percentiles**

```json
{"jsonrpc":"2.0","id":6,"method":"tools/call","params":{"name":"latency_percentiles","arguments":{"values":[12.5,45.3,67.8,23.1,89.4,34.6,56.7,78.9,11.2,99.0]}}}
```

## Tests

```bash
make test
```

## Coverage

```bash
make coverage
```

This generates `coverage.out` and `coverage.html`.
