package driver

import (
	_gorm "github.com/jeffcail/ginframe/server-common/pkg/gorm"
	_leveldb "github.com/jeffcail/ginframe/server-common/pkg/leveldb"
	_mongo "github.com/jeffcail/ginframe/server-common/pkg/mongo"
	_redis "github.com/jeffcail/ginframe/server-common/pkg/redis"
	leveldb1 "github.com/jeffcail/leveldb"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/redis.v5"
	"gorm.io/gorm"
)

// CreateDb create mysql db
func CreateDb(dbDsn string, maxOpenConn, maxIdleConn int) (*gorm.DB, error) {
	arg := []int{maxOpenConn, maxIdleConn}
	db, err := _gorm.InitGormMysql(dbDsn, arg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateRedis 初始化redis 连接
func CreateRedis(addr, pass string, rdb int) (*redis.Client, error) {
	rd, err := _redis.InitRedis(addr, pass, rdb)
	if err != nil {
		return nil, err
	}
	return rd, nil
}

// InitMongo 初始化mongo
func InitMongo(addr string) (*mongo.Client, error) {
	mg, err := _mongo.InitMongoDb(addr)
	if err != nil {
		return nil, err
	}
	return mg, nil
}

// InitLevelDb init level db
func InitLevelDb(path string) *leveldb1.LevelDB {
	return _leveldb.InitLevelDb(path)
}
