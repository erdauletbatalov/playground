package clock

import (
	"sync/atomic"
	"time"
)

// Clock tells current time as a Unix time, the number of nanoseconds
// elapsed since January 1, 1970 UTC.
type Clock interface {
	UnixNano() int64
}

// Now returns current time.Time told by c.
func Now(c Clock) time.Time {
	return time.Unix(0, c.UnixNano())
}

// GoClock is the clock backed by Go's "time.Now".
type GoClock struct{}

// UnixNano implements Clock interface.
func (x *GoClock) UnixNano() int64 {
	return time.Now().UnixNano()
}

// TickerClock is used to read current time in a delayed fashion releasing
// pressure on time package.
type TickerClock struct {
	nsec   int64
	stopCh chan struct{}
}

// NewTickerClock returns new clock with specified resolution. The
// time is read from time.Ticker every period specified with delay and
// then atomically set to variable.
func NewTickerClock(delay time.Duration) *TickerClock {
	stopCh := make(chan struct{}, 1)
	c := &TickerClock{time.Now().UnixNano(), stopCh}

	go func() {
		ticker := time.NewTicker(delay)
		defer ticker.Stop()
		for {
			select {
			case t := <-ticker.C:
				atomic.StoreInt64(&c.nsec, t.UnixNano())
			case <-stopCh:
				return
			}
		}
	}()

	return c
}

// UnixNano implements Clock interface. The value set by ticker is
// atomically read as the current time.
func (c *TickerClock) UnixNano() int64 {
	return atomic.LoadInt64(&c.nsec)
}

// Stop turns off a clock. After Stop, time may be read but is no
// longer updated.
func (c *TickerClock) Stop() {
	close(c.stopCh)
}
