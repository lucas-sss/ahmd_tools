package db

import (
	"ahmd_tools/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ahmdDB *gorm.DB
var proxyServerDB *gorm.DB

func Init(mysqlConf config.Mysql) {
	ahmdDbInit(mysqlConf)
	// proxyServerDbInit(mysqlConf)
}

func ahmdDbInit(mysqlConf config.Mysql) {
	dsn := mysqlConf.Username + ":" + mysqlConf.Password + "@tcp(" + mysqlConf.Host + ")/" + mysqlConf.AhmdDB + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("open ahmd db fail")
	}
	ahmdDB = db
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(2)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(5)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}

func proxyServerDbInit(mysqlConf config.Mysql) {
	dsn := mysqlConf.Username + ":" + mysqlConf.Password + "@tcp(" + mysqlConf.Host + ")/" + mysqlConf.ProxyDB + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("open proxy server db fail")
	}
	proxyServerDB = db
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(2)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(5)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
