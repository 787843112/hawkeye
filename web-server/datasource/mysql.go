package datasource

import (
	"fmt"
	"web-server/config"

	"github.com/jinzhu/gorm"
)

//Mysql ...
var db *gorm.DB

//NewMySQL ...
func NewMySQL(c *config.Mysql) (db *gorm.DB) {
	var err error
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Dbname,
		c.Config,
	)
	db, err = gorm.Open("mysql", url)
	if err != nil {
		panic("mysql数据库连接失败!")
	}
	db.SingularTable(true)
	registTable(db)
	return
}
