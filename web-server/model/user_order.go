package model

//UserOrder ...
type UserOrder struct {
	OrderID    int    `gorm:"primary_key;AUTO_INCREMENT;comment:'订单ID';not null" json:"OrderID"`
	UserID     int    `gorm:"comment:'用户ID';not null" json:"-"`
	MovieID    int    `gorm:"comment:'电影ID';not null" json:"MovieID" form:"MovieID"`
	MovieName  string `gorm:"type:varchar(20);comment:'电影名称';not null" json:"MovieName" form:"MovieName"`
	Date       string `gorm:"type:varchar(10);comment:'观影日期';not null" json:"Date" form:"Date"`
	StartTime  string `gorm:"type:varchar(5);comment:'开始时间';not null" json:"StartTime" form:"StartTime"`
	EndTime    string `gorm:"type:varchar(5);comment:'结束时间';not null" json:"EndTime" form:"EndTime"`
	CinemaName string `gorm:"type:varchar(20);comment:'影院名称';not null" json:"CinemaName" form:"CinemaName"`
	Address    string `gorm:"type:varchar(50);comment:'影院详细地址';not null" json:"Address" form:"Address"`
	HallName   string `gorm:"type:varchar(20);comment:'影厅名称';not null" json:"HallName" form:"HallName"`
	Number     int    `gorm:"type:tinyint(1);comment:'座位数量';not null" json:"Number" form:"Number"`
	Seat       string `gorm:"type:varchar(50);comment:'座位';not null" json:"Seat" form:"Seat"`
	Price      int    `gorm:"type:smallint(2);comment:'总价';not null" json:"Price" form:"Price"`
	Time
}
