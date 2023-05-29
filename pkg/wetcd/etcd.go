package wetcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func InitSingleEtcd(addr string) (*clientv3.Client, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {

	}
	return client, err
}

func InitClusterEtcd(addr []string) (*clientv3.Client, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {

	}
	return client, err
}

type Wetcd struct{}

func (w *Wetcd) Put(client *clientv3.Client, key, value string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = client.Put(ctx, key, value)
	return
}

func (w *Wetcd) Get(client *clientv3.Client, key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := client.Get(ctx, key)
	if err != nil {
		return "", err
	}
	var result string
	for _, v := range res.Kvs {
		result = string(v.Value)
	}
	return result, nil
}

func (w *Wetcd) Update(client *clientv3.Client, key, value string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = client.Put(ctx, key, value)
	return
}

func (w *Wetcd) Delete(client *clientv3.Client, key string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = client.Delete(ctx, key)
	return
}
