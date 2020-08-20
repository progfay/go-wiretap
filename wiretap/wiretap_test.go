package wiretap_test

import (
	"testing"

	"github.com/progfay/go-wiretap/wiretap"
)

func Test_Wiretap(t *testing.T) {
	t.Run("0 count with no call", func(t *testing.T) {
		f := func() {}
		f, r := wiretap.Wiretap(f)
		if r.GetCount() != 0 {
			t.Error("")
		}
	})

	t.Run("count up how many called", func(t *testing.T) {
		f := func() {}
		f, r := wiretap.Wiretap(f)

		f()
		if r.GetCount() != 1 {
			t.Error("")
		}

		f()
		if r.GetCount() != 2 {
			t.Error("")
		}

		f()
		if r.GetCount() != 3 {
			t.Error("")
		}
	})
}
