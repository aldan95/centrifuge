package timers

import (
	"sync"
	"time"
)

var timerPool sync.Pool

// AcquireTimer from pool.
// Don't forget to mark timer as read when the value from t.C was received.
func AcquireTimer(d time.Duration) *PoolTimer {
	v := timerPool.Get()
	if v == nil {
		return newTimer(d)
	}
	tm := v.(*PoolTimer)
	tm.Reset(d)
	return tm
}

// ReleaseTimer to pool.
func ReleaseTimer(tm *PoolTimer) {
	timerPool.Put(tm)
}

// PoolTimer ...
type PoolTimer struct {
	*time.Timer
	read bool
}

// MarkRead must be called after receiving value from timer chan.
func (t *PoolTimer) MarkRead() {
	t.read = true
}

// Reset timer safely.
func (t *PoolTimer) Reset(d time.Duration) {
	stopped := t.Stop()
	if !stopped && !t.read {
		<-t.C
	}
	t.Timer.Reset(d)
	t.read = false
}

func newTimer(d time.Duration) *PoolTimer {
	return &PoolTimer{
		Timer: time.NewTimer(d),
	}
}
