package _mongo

import (
	"context"
	"github.com/jeffcail/ginframe/common/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongoDb mongodb 连接
func InitMongoDb() (*mongo.Client, error) {
	addr := global.Config.Mongo.Addr
	clientOptions := options.Client().ApplyURI(addr)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	return client, nil
}
