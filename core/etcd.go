package core

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type EtcdX struct {
	Endpoints   []string `json:"endpoints"`
	DialTimeout int      `json:"dialTimeout"`
	Client      *clientv3.Client
}

func BuildEtcdX(endpoints []string) (*EtcdX, error) {
	var client *clientv3.Client
	var err error
	if client, err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	}); err != nil {
		return nil, err
	}
	return &EtcdX{
		Endpoints:   endpoints,
		DialTimeout: 5,
		Client:      client,
	}, nil
}

func (e *EtcdX) Close() {
	e.Client.Close()
}

func (e *EtcdX) Put(key, value string) error {
	_, err := e.Client.Put(e.Client.Ctx(), key, value)
	return err
}

func (e *EtcdX) Get(key string) (string, error) {
	resp, err := e.Client.Get(e.Client.Ctx(), key)
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) == 0 {
		return "", nil
	}
	return string(resp.Kvs[0].Value), nil
}
