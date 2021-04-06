package dao

import (
	"web-server/config"
	"web-server/datasource"

	"github.com/jinzhu/gorm"
)

//Dao ...
type Dao struct {
	c  *config.Config
	db *gorm.DB
}

//New new dao
func New(c *config.Config) (dao *Dao) {
	dao = &Dao{
		c:  c,
		db: datasource.NewMySQL(&c.Mysql),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.db.Close()
}
