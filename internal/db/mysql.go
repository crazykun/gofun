package db

import (
	"fmt"
	"gofun/conf"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlPool struct {
	Instance map[string]*gorm.DB
}

func (this *MysqlPool) addInstance(dbname string, db *gorm.DB) {
	fmt.Println("=============  初始化 mysql " + dbname + "  =============")
	this.Instance[dbname] = db
}

func (this *MysqlPool) GetInstance(dbname string) *gorm.DB {
	return this.Instance[dbname]
}

var pool *MysqlPool

func RegisterMysqlPool(clientName string, dbConfig conf.MySQLConfig) {

	if pool == nil {
		pool = new(MysqlPool)
		pool.Instance = make(map[string]*gorm.DB)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		sqlDB, _ := db.DB()

		sqlDB.SetMaxIdleConns(dbConfig.MinNum) // SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxOpenConns(dbConfig.MaxNum) // SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetConnMaxLifetime(time.Hour)    // SetConnMaxLifetime 设置了连接可复用的最大时间。
		pool.addInstance(clientName, db)
	}
}

func GetMysqlPool() *MysqlPool {
	return pool
}
