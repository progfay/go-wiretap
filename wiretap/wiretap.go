package wiretap

import (
	"fmt"
	"reflect"
	"sync/atomic"
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
	vDst := reflect.ValueOf(dst)
	vSrc := reflect.ValueOf(src)

	if vDst.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("First argument of WiretapWithArgs must be pointer of function")
	}

	vDst = vDst.Elem()
	if vDst.Kind() != reflect.Func {
		return nil, fmt.Errorf("First argument of WiretapWithArgs must be pointer of function")
	}

	if vSrc.Kind() != reflect.Func {
		return nil, fmt.Errorf("Second argument of WiretapWithArgs must be function")
	}

	if vDst.Type() != vSrc.Type() {
		return nil, fmt.Errorf("Value of first argument must be same type function as second argument")
	}

	r := Report{callCount: 0}
	f := func(values []reflect.Value) []reflect.Value {
		r.IncCount()
		return vSrc.Call(values)
	}

	vDst.Set(reflect.MakeFunc(vDst.Type(), f))

	return &r, nil
}
