package application

import "errors"

var ErrKeyNotFound = errors.New("configuration key not found")

type Configuration interface {
	Get(key string) (any, error)
	GetString(key string) (string, error)
	GetInt(key string) (int, error)
	GetBool(key string) (bool, error)
}
type ConfigLoader interface {
	ImportData() map[string]any
}
