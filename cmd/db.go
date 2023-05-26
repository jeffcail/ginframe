package cmd

import "github.com/jeffcail/ginframe/core/db"

// InitDb init gorm
func InitDb() {
	db.InitDb()
	db.InitRedisClient()
	//db.InitMongo()
	db.InitLevelDb()
}
