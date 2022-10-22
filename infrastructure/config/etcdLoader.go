package config

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type EtcdLoader struct {
	client *clientv3.Client
}

func (e EtcdLoader) getEtcDConnection() *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	return cli
}
func (e EtcdLoader) GetVal(key string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	r, err := e.client.Get(ctx, key)
	if err != nil {
		cancel()
		return nil, err
	}
	cancel()
	return r, nil
}
func (e EtcdLoader) ImportData() map[string]any {
	e.client = e.getEtcDConnection()
	keysToSet := []string{
		"DB_HOST",
		"DB_USER",
		"DB_PASS",
		"DB_NAME",
		"DB_PORT",
	}
	m := map[string]any{}
	for _, key := range keysToSet {
		if v, err := e.GetVal(key); err == nil {
			m[key] = v
		} else {
			panic(err)
		}
	}
	return m
}
