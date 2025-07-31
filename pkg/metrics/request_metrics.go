package metrics

import (
	"sync"
	"time"
)

// RequestMetrics tracks request rates per second and per minute.
type RequestMetrics struct {
	mu              sync.Mutex
	requestsPerSec  int
	requestsPerMin  int
	lastSec         time.Time
	lastMin         time.Time
}

// NewRequestMetrics creates a new RequestMetrics instance.
func NewRequestMetrics() *RequestMetrics {
	now := time.Now()
	return &RequestMetrics{
		lastSec: now,
		lastMin: now,
	}
}

// TrackRequest increments request counters.
func (rm *RequestMetrics) TrackRequest() {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	now := time.Now()

	// Reset per second counter if second has changed
	if now.Sub(rm.lastSec) >= time.Second {
		rm.requestsPerSec = 0
		rm.lastSec = now
	}
	// Reset per minute counter if minute has changed
	if now.Sub(rm.lastMin) >= time.Minute {
		rm.requestsPerMin = 0
		rm.lastMin = now
	}
	rm.requestsPerSec++
	rm.requestsPerMin++
}

// GetRates returns current request rates.
func (rm *RequestMetrics) GetRates() (int, int) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	return rm.requestsPerSec, rm.requestsPerMin
}