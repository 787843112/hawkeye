package model

//HallInfo ...
type HallInfo struct {
	HallID     int    `gorm:"primary_key;AUTO_INCREMENT;comment:'影厅ID';not null" json:"HallID"`
	CinemaID   int    `gorm:"comment:'影院ID';not null" json:"CinemaID" form:"CinemaID"`
	HallName   string `gorm:"type:varchar(20);comment:'影厅名称';not null" json:"HallName" form:"HallName"`
	HallRow    int    `gorm:"type:tinyint(1);comment:'影厅行数';not null" json:"HallRow" form:"HallRow"`
	HallColumn int    `gorm:"type:tinyint(1);comment:'影厅列数';not null" json:"HallColumn" form:"HallColumn"`
	Time       `json:"-"`
	//Seats []*SeatInfo `gorm:"ForeignKey:HallID" json:"-"`
}
