package config

import "github.com/d3n972/mavint/application"

type Config struct {
	data map[string]any
}

func (c Config) keyExists(key string) bool {
	if _, ok := c.data[key]; !ok {
		return false
	}
	return true
}

func NewConfig() application.Configuration {
	return Config{
		data: EnvLoader{}.ImportData(),
	}
}

func (c Config) Get(key string) (any, error) {
	if v, ok := c.data[key]; ok {
		return v, nil
	}
	return nil, application.ErrKeyNotFound
}

func (c Config) GetString(key string) (string, error) {
	v, err := c.Get(key)
	if err != nil {
		return "", err
	}
	return v.(string), nil
}

func (c Config) GetInt(key string) (int, error) {
	v, err := c.Get(key)
	if err != nil {
		return 0, err
	}
	return v.(int), nil
}

func (c Config) GetBool(key string) (bool, error) {
	v, err := c.Get(key)
	if err != nil {
		return false, err
	}
	return v.(bool), nil
}
