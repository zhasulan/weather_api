package utils

import (
	"fmt"
	"github.com/pkg/errors"
)

// builds error from panic, which can be of any type
func AnyError(r any) error {
	var ex error
	switch x := r.(type) {
	case string:
		ex = errors.New(x)
	case error:
		ex = x
	default:
		ex = fmt.Errorf("Unknown panic type: %v", r)
	}
	return ex
}
