package _leveldb

import (
	leveldb1 "github.com/jeffcail/leveldb"
)

func InitLevelDb(path string) *leveldb1.LevelDB {
	db, err := leveldb1.CreateLevelDB(path)
	if err != nil {
		panic(err)
	}
	return db
}
