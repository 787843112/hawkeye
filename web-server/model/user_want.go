package model

//UserWant ...
type UserWant struct {
	UserID  int `gorm:"primary_key;type:int;comment:'用户ID';not null" json:"UserID" form:"UserID"`
	MovieID int `gorm:"primary_key;type:int;comment:'电影ID';not null" json:"MovieID" form:"MovieID"`
	Time    `json:"-"`
}
