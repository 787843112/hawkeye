package model

//ScoreInfo ...
type ScoreInfo struct {
	MovieID    int `gorm:"primary_key;comment:'电影ID';not null"`
	ScoreOne   int `gorm:"default:0;comment:'打1分人数';not null"`
	ScoreTwo   int `gorm:"default:0;comment:'打2分人数';not null"`
	ScoreThree int `gorm:"default:0;comment:'打3分人数';not null" `
	ScoreFour  int `gorm:"default:0;comment:'打4分人数';not null"`
	ScoreFive  int `gorm:"default:0;comment:'打5分人数';not null"`
	ScoreSix   int `gorm:"default:0;comment:'打6分人数';not null"`
	ScoreSeven int `gorm:"default:0;comment:'打7分人数';not null"`
	ScoreEight int `gorm:"default:0;comment:'打8分人数';not null"`
	ScoreNine  int `gorm:"default:0;comment:'打9分人数';not null"`
	ScoreTen   int `gorm:"default:0;comment:'打10分人数';not null"`
	Time
}
