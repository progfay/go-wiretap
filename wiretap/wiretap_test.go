package wiretap_test

import (
	"testing"

	"github.com/progfay/go-wiretap/wiretap"
)

func globalFn()                                         {}
func globalFnWithArgs(a, b, c int)                      {}
func globalFnWithReturns() (string, string)             { return "first", "second" }
func globalFnWithArgsAndReturns(a, b, c int) (int, int) { return a + b + c, a * b * c }

func Test_Wiretap_global(t *testing.T) {
	t.Run("no args and no return values", func(t *testing.T) {
		var fn func()
		r, err := wiretap.Wiretap(&fn, globalFn)
		if err != nil {
			t.Error(err)
		}

		if c := r.GetCount(); c != 0 {
			t.Errorf("count is incorrect, want 0, got %v", c)
		}

		fn()
		if c := r.GetCount(); c != 1 {
			t.Errorf("count is incorrect, want 1, got %v", c)
		}
	})

	t.Run("many args and no return values", func(t *testing.T) {
		var fn func(a, b, c int)
		r, err := wiretap.Wiretap(&fn, globalFnWithArgs)
		if err != nil {
			t.Error(err)
		}

		if c := r.GetCount(); c != 0 {
			t.Errorf("count is incorrect, want 0, got %v", c)
		}

		fn(1, 2, 3)
		if c := r.GetCount(); c != 1 {
			t.Errorf("count is incorrect, want 1, got %v", c)
		}
	})

	t.Run("no args and many return values", func(t *testing.T) {
		var fn func() (string, string)
		r, err := wiretap.Wiretap(&fn, globalFnWithReturns)
		if err != nil {
			t.Error(err)
		}

		ret1, ret2 := fn()
		if c := r.GetCount(); c != 1 {
			t.Errorf("count is incorrect, want 1, got %v", c)
		}

		if ret1 != "first" || ret2 != "second" {
			t.Error("return value was incorrect")
		}
	})

	t.Run("many args and many return values", func(t *testing.T) {
		var fn func(a, b, c int) (int, int)
		r, err := wiretap.Wiretap(&fn, globalFnWithArgsAndReturns)
		if err != nil {
			t.Error(err)
		}

		ret1, ret2 := fn(1, 2, 3)
		if c := r.GetCount(); c != 1 {
			t.Errorf("count is incorrect, want 1, got %v", c)
		}

		if ret1 != 6 || ret2 != 6 {
			t.Error("return value was incorrect")
		}
	})
}

func Test_Wiretap_local(t *testing.T) {
	t.Run("no args and no return values", func(t *testing.T) {
		fn := func() {}
		r, err := wiretap.Wiretap(&fn, fn)
		if err != nil {
			t.Error(err)
		}

		if c := r.GetCount(); c != 0 {
			t.Errorf("count is incorrect, want 0, got %v", c)
		}

		fn()
		if c := r.GetCount(); c != 1 {
			t.Errorf("count is incorrect, want 1, got %v", c)
		}
	})

	t.Run("many args and no return values", func(t *testing.T) {
		fn := func(a, b, c int) {}
		r, err := wiretap.Wiretap(&fn, fn)
		if err != nil {
			t.Error(err)
		}

		if c := r.GetCount(); c != 0 {
			t.Errorf("count is incorrect, want 0, got %v", c)
		}

		fn(1, 2, 3)
		if c := r.GetCount(); c != 1 {
			t.Errorf("count is incorrect, want 1, got %v", c)
		}
	})

	t.Run("no args and many return values", func(t *testing.T) {
		fn := func() (string, string) { return "first", "second" }
		r, err := wiretap.Wiretap(&fn, fn)
		if err != nil {
			t.Error(err)
		}

		ret1, ret2 := fn()
		if c := r.GetCount(); c != 1 {
			t.Errorf("count is incorrect, want 1, got %v", c)
		}

		if ret1 != "first" || ret2 != "second" {
			t.Error("return value was incorrect")
		}
	})

	t.Run("many args and many return values", func(t *testing.T) {
		fn := func(a, b, c int) (int, int) { return a + b + c, a * b * c }
		r, err := wiretap.Wiretap(&fn, fn)
		if err != nil {
			t.Error(err)
		}

		ret1, ret2 := fn(1, 2, 3)
		if c := r.GetCount(); c != 1 {
			t.Errorf("count is incorrect, want 1, got %v", c)
		}

		if ret1 != 6 || ret2 != 6 {
			t.Error("return value was incorrect")
		}
	})
}

func Test_Wiretap_invalid(t *testing.T) {
	t.Run("first argument must be function pointer", func(t *testing.T) {
		fn := func() {}
		_, err := wiretap.Wiretap(fn, globalFn)
		if err == nil {
			t.Error("cannot set other than function pointer for first argument")
		}

		intValue := 0
		_, err = wiretap.Wiretap(&intValue, globalFn)
		if err == nil {
			t.Error("cannot set other than function pointer for first argument")
		}
	})

	t.Run("second argument must be function", func(t *testing.T) {
		fn := func() {}
		_, err := wiretap.Wiretap(fn, 1)
		if err == nil {
			t.Error("cannot set other than function for second argument")
		}
	})

	t.Run("value of first argument must be same type as second argument", func(t *testing.T) {
		fn1 := func(a int32) {}
		fn2 := func(a int64) {}
		_, err := wiretap.Wiretap(&fn1, fn2)
		if err == nil {
			t.Error("value of first argument must be same type as second argument")
		}
	})
}
