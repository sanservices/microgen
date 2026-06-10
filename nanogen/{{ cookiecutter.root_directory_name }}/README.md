# Go HTTP Function

Welcome to your new Go Function! The boilerplate function code can be found in
[`handle.go`](handle.go). This Function responds to HTTP requests.

## Development

Develop new features by adding a test to [`handle_test.go`](handle_test.go) for
each feature, and confirm it works with `go test`.

Update the running analog of the function using the `func` CLI or client
library, and it can be invoked from your browser or from the command line:

```console
curl http://myfunction.example.com/
```

For more, see [the complete documentation]('https://github.com/knative/func/tree/main/docs')

## Observability

This function emits [Datadog](https://docs.datadoghq.com/tracing/) APM traces. The tracer is started when the function
package loads and is configured entirely through the standard Datadog environment variables — `DD_ENV`, `DD_SERVICE`,
`DD_VERSION`, `DD_AGENT_HOST` (default `localhost`) and `DD_TRACE_AGENT_PORT` (default `8126`). If no Datadog agent is
reachable the tracer simply no-ops, so it is safe to leave enabled in local development.


