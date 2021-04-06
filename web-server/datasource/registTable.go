package datasource

import (
	"web-server/model"

	"github.com/jinzhu/gorm"
)

//注册数据库表
func registTable(db *gorm.DB) {
	//创建表
	if !db.HasTable(&model.UserWant{}) {
		db.CreateTable(&model.UserWant{})
		//创建触发器
		db.Exec("CREATE TRIGGER `add_want` AFTER INSERT ON `user_want` FOR EACH ROW BEGIN\n\tUPDATE movie_info as movie SET movie.want_number=movie.want_number+1 WHERE movie.movie_id=new.movie_id;\nEND")
		db.Exec("CREATE TRIGGER `delete_want` AFTER DELETE ON `user_want` FOR EACH ROW BEGIN\n\tUPDATE movie_info as movie SET movie.want_number=movie.want_number-1 WHERE movie.movie_id=old.movie_id;\nEND")
	}
	if !db.HasTable(&model.ManageInfo{}) {
		db.CreateTable(&model.ManageInfo{})
	}
	if !db.HasTable(&model.AdminInfo{}) {
		db.CreateTable(&model.AdminInfo{})
	}
	if !db.HasTable(&model.UserInfo{}) {
		db.CreateTable(&model.UserInfo{})
	}
	if !db.HasTable(&model.MovieInfo{}) {
		db.CreateTable(&model.MovieInfo{})
		//创建触发器
		db.Exec("CREATE TRIGGER `add_score_info` AFTER INSERT ON `movie_info` FOR EACH ROW BEGIN\n\tINSERT INTO score_info(movie_id) values(new.movie_id);\nEND")
	}
	if !db.HasTable(&model.MovieType{}) {
		db.CreateTable(&model.MovieType{})
	}
	if !db.HasTable(&model.CinemaInfo{}) {
		db.CreateTable(&model.CinemaInfo{})
	}
	if !db.HasTable(&model.HallInfo{}) {
		db.CreateTable(&model.HallInfo{})
	}
	if !db.HasTable(&model.ScreeningInfo{}) {
		db.CreateTable(&model.ScreeningInfo{})
	}
	if !db.HasTable(&model.SeatInfo{}) {
		db.CreateTable(&model.SeatInfo{})
	}
	if !db.HasTable(&model.StatusInfo{}) {
		db.CreateTable(&model.StatusInfo{})
	}
	if !db.HasTable(&model.UserSeen{}) {
		db.CreateTable(&model.UserSeen{})
	}
	if !db.HasTable(&model.UserOrder{}) {
		db.CreateTable(&model.UserOrder{})
		//创建触发器
		db.Exec("CREATE TRIGGER `add_user_seen` AFTER INSERT ON `user_order` FOR EACH ROW BEGIN\n\tINSERT INTO user_seen(user_id,movie_id) SELECT new.user_id,new.movie_id FROM DUAL WHERE NOT EXISTS(SELECT user_id,movie_id FROM user_seen WHERE user_id=new.user_id and movie_id=new.movie_id);\nEND")
		db.Exec("CREATE TRIGGER `add_box_office` AFTER INSERT ON `user_order` FOR EACH ROW BEGIN\n\tUPDATE movie_info as movie SET movie.box_office=movie.box_office+new.price WHERE movie.movie_id=new.movie_id;\nEND")
	}
	if !db.HasTable(&model.ScoreInfo{}) {
		db.CreateTable(&model.ScoreInfo{})
		//创建触发器
		db.Exec("CREATE TRIGGER `update_total` AFTER UPDATE ON `score_info` FOR EACH ROW BEGIN\n\tUPDATE movie_info as movie SET movie.total=new.score_one+new.score_two+new.score_three+new.score_four+new.score_five+new.score_six+new.score_seven+new.score_eight+new.score_nine+new.score_ten WHERE movie.movie_id=new.movie_id;\nEND")
		db.Exec("CREATE TRIGGER `update_score` AFTER UPDATE ON `score_info` FOR EACH ROW BEGIN\n\tUPDATE movie_info as movie SET score=(new.score_one+new.score_two*2+new.score_three*3+new.score_four*4+new.score_five*5+new.score_six*6+new.score_seven*7+new.score_eight*8+new.score_nine*9+new.score_ten*10)/movie.total WHERE movie.movie_id=new.movie_id;\nEND")
	}
	//给各表添加外键及约束
	// 1st: 外键字段 2nd: 外键表(字段) 3rd: ONDELETE 4th: ONUPDATE
	db.Model(&model.ManageInfo{}).AddForeignKey("admin_id", "admin_info(admin_id)", "CASCADE", "CASCADE")
	db.Model(&model.ManageInfo{}).AddForeignKey("cinema_id", "cinema_info(cinema_id)", "CASCADE", "CASCADE")
	db.Model(&model.UserWant{}).AddForeignKey("user_id", "user_info(user_id)", "CASCADE", "CASCADE")
	db.Model(&model.UserWant{}).AddForeignKey("movie_id", "movie_info(movie_id)", "CASCADE", "CASCADE")
	db.Model(&model.UserSeen{}).AddForeignKey("user_id", "user_info(user_id)", "CASCADE", "CASCADE")
	db.Model(&model.UserSeen{}).AddForeignKey("movie_id", "movie_info(movie_id)", "CASCADE", "CASCADE")
	db.Model(&model.HallInfo{}).AddForeignKey("cinema_id", "cinema_info(cinema_id)", "CASCADE", "CASCADE")
	db.Model(&model.ScreeningInfo{}).AddForeignKey("hall_id", "hall_info(hall_id)", "CASCADE", "CASCADE")
	db.Model(&model.ScreeningInfo{}).AddForeignKey("movie_id", "movie_info(movie_id)", "CASCADE", "CASCADE")
	db.Model(&model.SeatInfo{}).AddForeignKey("hall_id", "hall_info(hall_id)", "CASCADE", "CASCADE")
	db.Model(&model.StatusInfo{}).AddForeignKey("Screening_id", "Screening_info(Screening_id)", "CASCADE", "CASCADE")
	db.Model(&model.UserOrder{}).AddForeignKey("user_id", "user_info(user_id)", "CASCADE", "CASCADE")
	// db.Model(&model.UserOrder{}).AddForeignKey("movie_id", "movie_info(movie_id)", "RESTRICT", "CASCADE")
	db.Model(&model.ScoreInfo{}).AddForeignKey("movie_id", "movie_info(movie_id)", "CASCADE", "CASCADE")
	//创建索引
	db.Model(&model.CinemaInfo{}).AddIndex("city", "city")
	db.Model(&model.MovieInfo{}).AddIndex("status", "status")
	db.Model(&model.MovieInfo{}).AddIndex("region", "region")
}
