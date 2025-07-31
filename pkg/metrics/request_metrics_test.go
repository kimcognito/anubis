package metrics

import (
	"testing"
	"time"
)

func TestRequestMetrics(t *testing.T) {
	rm := NewRequestMetrics()

	// Simulate 5 requests in the same second
	for i := 0; i < 5; i++ {
		rm.TrackRequest()
	}
	sec, min := rm.GetRates()
	if sec != 5 || min != 5 {
		t.Fatalf("expected 5/sec and 5/min, got %d/sec and %d/min", sec, min)
	}

	// Wait for a second to pass
	time.Sleep(time.Second)
	rm.TrackRequest()
	sec, min = rm.GetRates()
	if sec != 1 {
		t.Fatalf("expected 1/sec after reset, got %d", sec)
	}

	// Wait for a minute to pass
	time.Sleep(time.Minute)
	rm.TrackRequest()
	_, min = rm.GetRates()
	if min != 1 {
		t.Fatalf("expected 1/min after reset, got %d", min)
	}
}