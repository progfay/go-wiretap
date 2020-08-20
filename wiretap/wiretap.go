package wiretap

import "sync/atomic"

type Report struct {
	callCount int64
}

func (r *Report) IncCount() {
	atomic.AddInt64(&r.callCount, 1)
}

func (r *Report) GetCount() int64 {
	return atomic.LoadInt64(&r.callCount)
}

func Wiretap(f func()) (func(), *Report) {
	r := Report{callCount: 0}
	return func() {
		r.IncCount()
		f()
	}, &r
}
