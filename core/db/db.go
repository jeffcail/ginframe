package db

import (
	"fmt"
	"github.com/jeffcail/ginframe/common/global"
	_gorm "github.com/jeffcail/ginframe/pkg/gorm"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func InitDb() {
	arg := []int{global.Config.Mysql.MaxOpenConns, global.Config.Mysql.MaxIdleConns}
	Db, err = _gorm.InitGormMysql(global.Config.Mysql.DbDsn, arg)
	if err != nil {
		panic(err)
	}
	fmt.Println("mysql connection success...")
}
