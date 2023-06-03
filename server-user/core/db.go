package core

import (
	"context"
	"errors"
	_leveldb "github.com/jeffcail/ginframe/server-common/pkg/leveldb"
	leveldb1 "github.com/jeffcail/leveldb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/redis.v5"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	Rd  *redis.Client
	mg  *mongo.Client
	Ldb *leveldb1.LevelDB
)

func SetMysql(e *gorm.DB) {
	Db = e
}

func SetRedis(_rd *redis.Client) {
	Rd = _rd
}

func SetMongo(_m *mongo.Client) {
	mg = _m
}

func SetLevelDB(ldb *leveldb1.LevelDB) {
	Ldb = ldb
}

type MongoClient struct{}

// InsertCollection 写入 db 写入的数据库  collection 写入的文档（表） data 写入的数据
func (mc *MongoClient) InsertCollection(db string, collection string, data interface{}) (
	res *mongo.InsertOneResult, err error) {
	c := mg.Database(db).Collection(collection)
	return c.InsertOne(context.TODO(), data)
}

// BatchInsertCollection 批量写入
func (mc *MongoClient) BatchInsertCollection(db string, collection string, data []interface{}) (
	res *mongo.InsertManyResult, err error) {
	c := mg.Database(db).Collection(collection)
	return c.InsertMany(context.TODO(), data)
}

// UpdateOneRecord 修改单条记录
func (mc *MongoClient) UpdateOneRecord(db string, collection string, id string, data bson.D) (
	res *mongo.UpdateResult, err error) {
	c := mg.Database(db).Collection(collection)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objId}}
	d := bson.D{{"$set", data}}
	res, err = c.UpdateOne(context.TODO(), filter, d)
	return
}

// CountCollection 统计
func (mc *MongoClient) CountCollection(db string, collection string, filter map[string]interface{}) (int64, int64, error) {
	c := mg.Database(db).Collection(collection)

	esCount, err := c.EstimatedDocumentCount(context.TODO())
	if err != nil {
		return -1, -1, err
	}

	if len(filter) == 0 {
		count, err := c.CountDocuments(context.TODO(), bson.D{{}})
		if err != nil {
			return -1, -1, err
		}
		return esCount, count, nil
	}

	if len(filter) > 0 {
		for k, v := range filter {
			count, err := c.CountDocuments(context.TODO(), bson.D{{k, v}})
			if err != nil {
				return -1, -1, err
			}
			return esCount, count, nil
		}
	}

	return -1, -1, errors.New("查询条件异常")
}

// DeleteOneRecord 删除单条记录
func (mc *MongoClient) DeleteOneRecord(db string, collection string, filter map[string]interface{}) (
	res *mongo.DeleteResult, err error) {
	c := mg.Database(db).Collection(collection)
	for k, v := range filter {
		filter := bson.D{{k, v}}
		return c.DeleteOne(context.TODO(), filter)
	}
	return
}

// InitLevelDb init level db
func (mc *MongoClient) InitLevelDb(path string) *leveldb1.LevelDB {
	return _leveldb.InitLevelDb(path)
}
