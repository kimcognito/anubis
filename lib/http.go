// lib/http.go
// ... other imports ...
import (
    // existing imports ...
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

// Add request rate metrics
var requestsHandled = promauto.NewCounterVec(
    prometheus.CounterOpts{
        Name: "anubis_requests_handled_total",
        Help: "Total number of HTTP requests handled",
    },
    []string{"method", "path"},
)

// Inside your Server methods:
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    requestsHandled.WithLabelValues(r.Method, r.URL.Path).Inc()
    s.mux.ServeHTTP(w, r)
}

// If you handle requests in ServeHTTPNext, add instrumentation similarly:
func (s *Server) ServeHTTPNext(w http.ResponseWriter, r *http.Request) {
    requestsHandled.WithLabelValues(r.Method, r.URL.Path).Inc()
    // rest of function as before...
}