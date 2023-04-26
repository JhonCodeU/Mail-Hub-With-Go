package profiling

import (
	"net/http"
)

// WithStats wraps a handler with stats reporting. It tracks the metrics such as
// as the number of requests, the number of errors, and the latency of requests.

type StatsHandler struct {
	handler http.Handler
}
