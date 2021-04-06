package model

//UserSeen ...
type UserSeen struct {
	UserID  int    `gorm:"primary_key;type:int;comment:'用户ID';not null" json:"UserID" form:"UserID"`
	MovieID int    `gorm:"primary_key;type:int;comment:'电影ID';not null" json:"MovieID" form:"MovieID"`
	Score   int    `gorm:"type:tinyint(1);comment:'所评分数';" json:"Score" form:"Score"`
	Review  string `gorm:"type:varchar(1024);comment:'评价';" json:"Review" form:"Review"`
	Time
	MovieInfo *MovieInfo `gorm:"ForeignKey:MovieID" json:"MovieInfo"`
}
