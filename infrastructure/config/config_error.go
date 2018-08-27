package config

import (
	"errors"
)

var (
	//ErrConfigFileNotFound indicates Vanda can't find configuration file.
	ErrConfigFileNotFound = errors.New("Vanda::configuration file not found")
	// ErrConfigFileFormat indicates wrong configuration file format is being used.
	ErrConfigFileFormat = errors.New("Vanda::configuration file format isn't toml file")
)
