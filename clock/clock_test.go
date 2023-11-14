package clock

import (
	"testing"
	"time"
)

// Бенчмарк для time.Now().UnixNano()
func BenchmarkTimeNowUnixNano(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = time.Now().UnixNano()
	}
}

// Бенчмарк для time.Now()
func BenchmarkClockTimeUnixNano(b *testing.B) {
	// prepare an efficient Clock
	c := NewTickerClock(time.Second)
	defer c.Stop()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.UnixNano()
	}
}

// Бенчмарк для time.Now()
func BenchmarkClockTimeNow(b *testing.B) {
	// prepare an efficient Clock
	c := NewTickerClock(time.Millisecond)
	defer c.Stop()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.UnixNano()
	}
}
