package errorx

import (
	"errors"
)

// Errorx is a struct for wrap raw err with message
type Errorx struct {
	message string
	err     error
}

// Error return the error message
func (e Errorx) Error() string {
	return e.message
}

// Is reports whether any error in err's chain matches target.
func (e Errorx) Is(target error) bool {
	if target == nil || e.err == nil {
		return e.err == target
	}

	return errors.Is(e.err, target)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
func (e *Errorx) Unwrap() error {
	u, ok := e.err.(interface {
		Unwrap() error
	})
	if !ok {
		return e.err
	}

	return u.Unwrap()
}
