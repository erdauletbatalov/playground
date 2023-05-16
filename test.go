package main

import "github.com/pkg/errors"

func main() {
	// Using errors.Wrap constructs a stack of errors, adding context to the
	// preceding error. Depending on the nature of the error it may be necessary
	// to reverse the operation of errors.Wrap to retrieve the original error for
	// inspection. Any error value which implements this interface

	type causer interface {
		Cause() error
	}
	// can be inspected by errors.Cause. errors.Cause will recursively retrieve
	// the topmost error that does not implement causer, which is assumed to be
	// the original cause. For example:

	switch err := errors.Cause(err).(type) {
	case *MyError:
		// handle specifically
	default:
		// unknown error
	}
	// Although the causer interface is not exported by this package, it is
	// considered a part of its stable public interface.
}
