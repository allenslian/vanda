package db

import (
	"errors"
)

var (
	//ErrMissingOption indicates system didn't find database setting.
	errMissingOption = errors.New("Missing database option")
)
