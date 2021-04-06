package model

//CinemaInfo ...
type CinemaInfo struct {
	CinemaID   int    `gorm:"primary_key;AUTO_INCREMENT;comment:'影院ID';not null" json:"CinemaID" form:"CinemaID"`
	CinemaName string `gorm:"type:varchar(32);comment:'影院名称';not null" json:"CinemaName" form:"CinemaName"`
	City       string `gorm:"type:varchar(20);comment:'影院所在城市';not null" json:"City" form:"City"`
	Address    string `gorm:"type:varchar(50);comment:'影院详细地址';not null" json:"Address" form:"Address"`
	Time       `json:"-"`
	//Halls []*HallInfo `gorm:"ForeignKey:CinemaID" json:"-"`
	//Admins []*AdminInfo `gorm:"many2many:manage_info;ForeignKey:CinemaID" json:"-"`
}
