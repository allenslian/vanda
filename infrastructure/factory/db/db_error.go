package db

import (
	"errors"
)

var (
	//ErrMissingDBOption indicates system didn't find database setting.
	ErrMissingDBOption = errors.New("Missing database option")
)
