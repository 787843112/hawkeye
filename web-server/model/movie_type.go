package model

//MovieType ...
type MovieType struct {
	TypeID   int    `gorm:"primary_key;AUTO_INCREMENT;comment:'类型ID';not null" json:"-"`
	TypeName string `gorm:"type:varchar(2);unique;comment:'类型名称';not null" json:"TypeName" form:"TypeName"`
	Time     `json:"-"`
}
