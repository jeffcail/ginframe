package boot

import (
	"github.com/jeffcail/ginframe/server-common/driver"
	"github.com/jeffcail/ginframe/server-user/core"
	"github.com/jeffcail/ginframe/server-user/ulogger"
	"log"
)

// InitDb init gorm
func InitDb(dbDsn string, maxOpenConn, maxIdleConn int) {
	gdb, err := driver.CreateDb(dbDsn, maxOpenConn, maxIdleConn)
	if err != nil {
		log.Fatalf("Create mysql connection err: %v\n", err)
	}
	core.SetMysql(gdb)

	ulogger.UserLogger.Info("🚀🚀🚀🚀🚀🚀mysql success...🚀🚀🚀🚀🚀🚀")
}

// InitRedis init redis
func InitRedis(addr, pass string, rdb int) {
	rd, err := driver.CreateRedis(addr, pass, rdb)
	if err != nil {
		log.Fatalf("Create redis client err: %v\n", err)
	}
	core.SetRedis(rd)

	ulogger.UserLogger.Info("🚀🚀🚀🚀🚀🚀redis success...🚀🚀🚀🚀🚀🚀")
}

// InitMongo init mongodb
func InitMongo(addr string) {
	mongo, err := driver.InitMongo(addr)
	if err != nil {
		log.Fatalf("Create mongodb client err: %v\n", err)
	}
	core.SetMongo(mongo)
	ulogger.UserLogger.Info("🚀🚀🚀🚀🚀🚀 mongodb success...🚀🚀🚀🚀🚀🚀")
}

// InitLevelDb init leveldb
func InitLevelDb(path string) {
	ldb := driver.InitLevelDb(path)
	core.SetLevelDB(ldb)

	ulogger.UserLogger.Info("🚀🚀🚀🚀🚀🚀levedb success...🚀🚀🚀🚀🚀🚀")
}
