package handler

import (
	"errors"
)

var (
	//ErrEmptyRouteKeyNotAllowed describes one error about empty route key.
	ErrEmptyRouteKeyNotAllowed = errors.New("Not allow to use empty route key")
	//ErrSlashPrefixNotAllowed describes one error about empty route key.
	ErrSlashPrefixNotAllowed = errors.New("Not allow to use slash prefix as route key")
)
