package utils

import "fmt"

func CatchError(err *error) {
	if r := recover(); r != nil {
		switch r.(type) {
		case error:
			*err = r.(error)
		default:
			*err = fmt.Errorf("%v", r)
		}
	}
}
