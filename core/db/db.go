package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jeffcail/ginframe/common/global"
	_gorm "github.com/jeffcail/ginframe/pkg/gorm"
	_mongo "github.com/jeffcail/ginframe/pkg/mongo"
	_redis "github.com/jeffcail/ginframe/pkg/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/redis.v5"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
	Rd  *redis.Client
	mg  *mongo.Client
)

func InitDb() {
	arg := []int{global.Config.Mysql.MaxOpenConns, global.Config.Mysql.MaxIdleConns}
	Db, err = _gorm.InitGormMysql(global.Config.Mysql.DbDsn, arg)
	if err != nil {
		panic(err)
	}
	fmt.Println("mysql connection success...")
}

// InitRedisClient 初始化redis 连接
func InitRedisClient() {
	Rd, err = _redis.InitRedis()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis connection success...")
}

// InitMongo 初始化mongo
func InitMongo() {
	mg, err = _mongo.InitMongoDb()
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected and pinged.")
}

type mongoClient struct{}

// NewMongoClient mongo实例
func NewMongoClient() *mongoClient {
	return &mongoClient{}
}

// InsertCollection 写入 db 写入的数据库  collection 写入的文档（表） data 写入的数据
func (mc *mongoClient) InsertCollection(db string, collection string, data interface{}) (
	res *mongo.InsertOneResult, err error) {
	c := mg.Database(db).Collection(collection)
	return c.InsertOne(context.TODO(), data)
}

// BatchInsertCollection 批量写入
func (mc *mongoClient) BatchInsertCollection(db string, collection string, data []interface{}) (
	res *mongo.InsertManyResult, err error) {
	c := mg.Database(db).Collection(collection)
	return c.InsertMany(context.TODO(), data)
}

// UpdateOneRecord 修改单条记录
func (mc *mongoClient) UpdateOneRecord(db string, collection string, id string, data bson.D) (
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
func (mc *mongoClient) CountCollection(db string, collection string, filter map[string]interface{}) (int64, int64, error) {
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
func (mc *mongoClient) DeleteOneRecord(db string, collection string, filter map[string]interface{}) (
	res *mongo.DeleteResult, err error) {
	c := mg.Database(db).Collection(collection)
	for k, v := range filter {
		filter := bson.D{{k, v}}
		return c.DeleteOne(context.TODO(), filter)
	}
	return
}
