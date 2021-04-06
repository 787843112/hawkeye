package model

import "time"

//ScreeningInfo ...
type ScreeningInfo struct {
	ScreeningID int       `gorm:"primary_key;AUTO_INCREMENT;comment:'场次ID';not null" json:"ScreeningID"`
	HallID      int       `gorm:"comment:'影厅ID';not null" json:"HallID" form:"HallID"`
	HallName    string    `gorm:"type:varchar(20);comment:'影厅名称';not null" json:"HallName" form:"HallName"`
	MovieID     int       `gorm:"comment:'电影ID';not null" json:"MovieID" form:"MovieID"`
	StartTime   time.Time `gorm:"type:TIMESTAMP;comment:'开始时间';not null" json:"StartTime" form:"StartTime"`
	EndTime     time.Time `gorm:"type:TIMESTAMP;comment:'结束时间';not null" json:"EndTime" form:"EndTime"`
	Price       int       `gorm:"type:tinyint(1);comment:'单价';not null" json:"Price" form:"Price"`
	Time        `json:"-"`
	//SeatsStatus []*StatusInfo `gorm:"ForeignKey:ScreeningID" json:"-"`
}
