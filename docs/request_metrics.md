# Request Rate Metrics

This module provides request rate metrics tracking for per-second and per-minute intervals.

## Usage

```go
rm := metrics.NewRequestMetrics()
rm.TrackRequest()
sec, min := rm.GetRates()
fmt.Printf("Requests this second: %d, this minute: %d\n", sec, min)
```

## CEL Integration

You can expose `GetRates()` metrics for use in CEL expressions to enforce request limits, detect abuse, or trigger alerts.

Example CEL usage:
```
request_metrics.GetRates().per_sec > 100
request_metrics.GetRates().per_min > 1000
```