package _gorm

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// InitGormMysql gorm 初始化mysql连接
// eg: root:123456@tcp(127.0.0.1:3306)/jiaxiao?charset=utf8mb4&parseTime=True&loc=Local
func InitGormMysql(dsn string, args ...[]int) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名加s
		},
		Logger:                                   logger.Default.LogMode(logger.Info), // 打印sql语句
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
	})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	if len(args) > 0 {
		if len(args[0]) >= 3 {
			return nil, errors.New("参数错误")
		}
		if len(args[0]) == 1 {
			// 设置数据库连接池最大连接数
			sqlDB.SetMaxOpenConns(args[0][0])
		}
		if len(args[0]) == 2 {
			// 连接池最大允许的空闲连接数
			sqlDB.SetMaxIdleConns(args[0][1])
		}
	} else {
		// 设置数据库连接池最大连接数 100
		sqlDB.SetMaxOpenConns(100)
		// 连接池最大允许的空闲连接数 20
		sqlDB.SetMaxIdleConns(20)
	}

	go func() {
		for {
			sqlDB.Ping()
			time.Sleep(1 * time.Hour)
		}
	}()

	return db, nil
}
