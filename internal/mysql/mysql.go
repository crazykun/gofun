package mysql

import (
	"gofun/conf"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var DbMap map[string]*gorm.DB

func Dsn(c conf.MySQLConfig) string {
	return c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Database + "?" + c.Charset
}

type DB struct {
	Name string
	Ctx  *gin.Context
}

func (d *DB) GetConn() *gorm.DB {
	return DbMap[d.Name]
}

func NewDb(ctx *gin.Context, name string) DB {
	return DB{
		Name: name,
		Ctx:  ctx,
	}
}
