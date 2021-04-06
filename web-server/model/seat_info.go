package model

//SeatInfo ...
type SeatInfo struct {
	HallID     int `gorm:"primary_key;type:int;comment:'影厅ID';not null" json:"HallID" form:"HallID"`
	SeatRow    int `gorm:"primary_key;type:tinyint(1);comment:'座位行号';not null" json:"SeatRow" form:"SeatRow"`
	SeatColumn int `gorm:"primary_key;type:tinyint(1);comment:'座位列号';not null" json:"SeatColumn" form:"SeatColumn"`
	Time       `json:"-"`
}
