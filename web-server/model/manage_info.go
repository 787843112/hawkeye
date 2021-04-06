package model

//ManageInfo ...
type ManageInfo struct {
	AdminID  int `gorm:"primary_key;type:int;comment:'管理员ID';not null" json:"AdminID" form:"AdminID"`
	CinemaID int `gorm:"primary_key;type:int;comment:'影院ID';not null" json:"CinemaID" form:"CinemaID"`
	Time     `json:"-"`
}
