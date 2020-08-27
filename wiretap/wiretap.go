package wiretap

import (
	"sync/atomic"

	"github.com/progfay/go-wrap/wrap"
)

type Report struct {
	callCount int64
	fn        interface{}
}

func (r *Report) IncCount() {
	atomic.AddInt64(&r.callCount, 1)
}

func (r *Report) GetCount() int64 {
	return atomic.LoadInt64(&r.callCount)
}

func Wiretap(dst, src interface{}) (*Report, error) {
	r := Report{callCount: 0}
	err := wrap.Before(dst, src, func() {
		r.IncCount()
	})
	if err != nil {
		return nil, err
	}

	return &r, nil
}
