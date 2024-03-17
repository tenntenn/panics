package panics

import (
	"errors"
	"fmt"
)

// Recover calls f and recovering a panic.
func Recover(f func() error) error {
	return withRecover(f)
}

// Value returns a value if err has recovered value with [Recover].
func Value(err error) any {
	var rerr *recoveredError
	if errors.As(err, &rerr) {
		return rerr.value
	}
	return nil
}

// IsRecovered returns whether err has recovered value with [Recover] or not.
func IsRecovered(err error) bool {
	return errors.As(err, new(*recoveredError))
}

func withRecover(f func() error) (rerr error) {
	// double defer sandwitch
	// more details see https://go.dev/cl/134395
	var (
		normalReturn bool
		recovered    bool
		panicValue   any
	)

	defer func() {
		if recovered {
			rerr = &recoveredError{
				value: panicValue,
			}
		}
	}()

	func() {
		defer func() {
			panicValue = recover()
		}()
		rerr = f()
		normalReturn = true
	}()

	if !normalReturn {
		recovered = true
	}

	return rerr
}

type recoveredError struct {
	value any
}

var _ error = (*recoveredError)(nil)

// implements error intercace
func (err *recoveredError) Error() string {
	return fmt.Sprintf("recovered with: %v", err.value)
}
