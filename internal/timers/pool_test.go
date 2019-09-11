package timers

import (
	"testing"
	"time"
)

func BenchmarkPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tm := AcquireTimer(time.Microsecond)
		<-tm.C
		tm.MarkRead()
		ReleaseTimer(tm)
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tm := time.NewTimer(time.Microsecond)
		<-tm.C
	}
	b.StopTimer()
	b.ReportAllocs()
}
