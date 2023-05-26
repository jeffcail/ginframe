package _leveldb

import (
	"github.com/jeffcail/ginframe/common/global"
	leveldb1 "github.com/jeffcail/leveldb"
)

func InitLevelDb() *leveldb1.LevelDB {
	db, err := leveldb1.CreateLevelDB(global.Config.LevelDb.Path)
	if err != nil {
		panic(err)
	}
	return db
}
