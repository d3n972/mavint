package config

import "os"

type EnvLoader struct {
}

func (e EnvLoader) ImportData() map[string]any {
	m := map[string]any{}
	m["DB_HOST"] = os.Getenv("DB_HOST")
	m["DB_USER"] = os.Getenv("DB_USER")
	m["DB_PASS"] = os.Getenv("DB_PASS")
	m["DB_NAME"] = os.Getenv("DB_NAME")
	m["DB_PORT"] = os.Getenv("DB_PORT")
	return m
}
