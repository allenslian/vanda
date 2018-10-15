package cache

import (
	"errors"
)

var (
	//ErrMissingOption means cache wasn't initialized yet.
	errMissingOption = errors.New("Missing cache option")
)
