package panics_test

import (
	"errors"
	"runtime"
	"testing"

	"github.com/tenntenn/panics"
	"golang.org/x/sync/errgroup"
)

func TestRecover(t *testing.T) {
	cases := []struct {
		name       string
		panic      bool
		panicValue any
		goexit     bool
		err        error
	}{
		{"nopanic", false, nil, false, nil},
		{"with error", false, nil, false, errors.New("error")},
		{"with runtime.Goexit", false, nil, true, nil},
		{"panic", true, 100, false, nil},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var eg errgroup.Group
			// run Recover in another goroutine for runtime.Goexit test
			eg.Go(func() error {
				return panics.Recover(func() error {
					switch {
					case tt.panic:
						panic(tt.panicValue)
					case tt.goexit:
						runtime.Goexit()
					}
					return tt.err
				})
			})

			err := eg.Wait()
			got := panics.Value(err)

			t.Log("error:", err)
			t.Log("panic value:", got)

			switch {
			case tt.panic:
				if got != tt.panicValue {
					t.Errorf("panics.IsRecovered(%q) == %v want %v", err, got, tt.panicValue)
				}
			default:
				if err != tt.err {
					t.Errorf("got error: %q want %q", err, tt.err)
				}

				if panics.IsRecovered(err) {
					t.Errorf("panics.IsRecovered(%q) == true want false", err)
				}
			}
		})
	}
}
