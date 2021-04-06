package model

import "time"

//Time ...
type Time struct {
	CreateAt  time.Time `gorm:"column:create_time;type:TIMESTAMP;default:CURRENT_TIMESTAMP;comment:'创建时间';not null" json:"CreateTime"`
	UpdatedAt time.Time `gorm:"column:update_time;type:TIMESTAMP;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间';not null" json:"-"`
}
