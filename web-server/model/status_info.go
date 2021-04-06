package model

//StatusInfo ...
type StatusInfo struct {
	ScreeningID int `gorm:"primary_key;type:int;comment:'场次ID';not null" json:"ScreeningID"`
	SeatRow     int `gorm:"primary_key;type:tinyint(1);comment:'座位行号';not null" json:"SeatRow"`
	SeatColumn  int `gorm:"primary_key;type:tinyint(1);comment:'座位列号';not null" json:"SeatColumn"`
	Status      int `gorm:"type:tinyint(1);default:0;comment:'座位状态';not null" json:"Status"`
	Time        `json:"-"`
}
