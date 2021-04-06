package dao

import (
	"errors"
	"web-server/model"
	"web-server/util"
)

//FindUserByUserID ...
func (d *Dao) FindUserByUserID(userID int) (user *model.UserInfo, err error) {
	user = &model.UserInfo{}
	return user, d.db.Where("user_id = ?", userID).First(user).Error
}

//FindUserByUserName ...
func (d *Dao) FindUserByUserName(userName string) (user *model.UserInfo, err error) {
	user = &model.UserInfo{}
	return user, d.db.Where("user_name = ?", userName).First(user).Error
}

//CreateUser ...
func (d *Dao) CreateUser(userName, password, phone string) (user *model.UserInfo, err error) {
	user = &model.UserInfo{UserName: userName, Password: password, Phone: phone}
	return user, d.db.Create(user).Error
}

//UpdateUser ...
func (d *Dao) UpdateUser(userID int, nickname, password, phone string) (user *model.UserInfo, err error) {
	user = &model.UserInfo{}
	if err = d.db.Where("user_id = ?", userID).First(user).Error; err != nil {
		return nil, errors.New("用户不存在")
	}
	if nickname != "" {
		user.Nickname = nickname
	}
	if password != "" {
		password = util.MD5V([]byte(password))
		user.Password = password
	}
	if phone != "" {
		user.Phone = phone
	}
	if err = d.db.Save(user).Error; err != nil {
		return nil, errors.New("修改失败")
	}
	return
}

//UpdateAvatarByUser ...
func (d *Dao) UpdateAvatarByUser(userID int, avatar string) (user *model.UserInfo, err error) {
	user = &model.UserInfo{}
	if err = d.db.Where("user_id = ?", userID).First(user).Error; err != nil {
		return nil, errors.New("用户不存在")
	}
	user.Avatar = avatar
	return user, d.db.Save(user).Error
}

//FindUserOrder ...
func (d *Dao) FindUserOrder(userID, page, size int) (count int, orders []*model.UserOrder, err error) {
	user := &model.UserInfo{}
	if err = d.db.Where("user_id = ?", userID).First(user).Error; err != nil {
		return 0, nil, errors.New("用户不存在")
	}
	d.db.Model(model.UserOrder{}).Where("user_id = ?", userID).Count(&count)
	if d.db.Where("user_id = ?", userID).Limit(size).Offset((page-1)*size).Order("create_time desc").Find(&orders).Error != nil {
		return 0, nil, errors.New("获取失败")
	}
	return
}

//FindUserWant ...
func (d *Dao) FindUserWant(userID, page, size int) (count int, movies []*model.MovieInfo, err error) {
	user := &model.UserInfo{}
	if err = d.db.Where("user_id = ?", userID).First(user).Error; err != nil {
		return 0, nil, errors.New("用户不存在")
	}
	d.db.Model(&user).Related(&movies, "WantMovies")
	count = len(movies)
	if d.db.Model(&user).Limit(size).Offset((page-1)*size).Order("create_time desc").Related(&movies, "WantMovies").Error != nil {
		return 0, nil, errors.New("获取失败")
	}
	return
}

//FindUserWantMovieID ...
func (d *Dao) FindUserWantMovieID(userID int) (movieIDs []int, err error) {
	user := &model.UserInfo{}
	if d.db.Where("user_id = ?", userID).First(user).Error != nil {
		return nil, errors.New("用户不存在")
	}
	movies := []*model.MovieInfo{}
	if d.db.Model(&user).Related(&movies, "WantMovies").Error != nil {
		return nil, errors.New("获取失败")
	}
	for _, v := range movies {
		movieIDs = append(movieIDs, v.MovieID)
	}
	return
}

//FindWant ...
func (d *Dao) FindWant(want *model.UserWant) error {
	return d.db.First(&want).Error
}

//FindUserSeen ...
func (d *Dao) FindUserSeen(userID, page, size int) (count int, seens []*model.UserSeen, err error) {
	user := &model.UserInfo{}
	if err = d.db.Where("user_id = ?", userID).First(user).Error; err != nil {
		return 0, nil, errors.New("用户不存在")
	}
	d.db.Model(&user).Related(&seens, "UserSeens")
	count = len(seens)
	if d.db.Model(&user).Limit(size).Offset((page-1)*size).Order("create_time desc").Related(&seens, "UserSeens").Error != nil {
		return 0, nil, errors.New("获取失败")
	}
	return
}

//FindSeen ...
func (d *Dao) FindSeen(userID, movieID int) (seen *model.UserSeen, err error) {
	seen = &model.UserSeen{UserID: userID, MovieID: movieID}
	return seen, d.db.First(seen).Error
}

//CreateUserWant ...
func (d *Dao) CreateUserWant(want *model.UserWant) error {
	return d.db.Create(want).Error
}

//DeleteUserWant ...
func (d *Dao) DeleteUserWant(want *model.UserWant) error {
	return d.db.Delete(want).Error
}

//数据库触发器完成
// func (d *Dao) CreateUserSeen(seen *model.UserSeen) error {
// 	return d.db.Create(seen).Error
// }

//DeleteUserSeen ...
func (d *Dao) DeleteUserSeen(seen *model.UserSeen) error {
	if d.db.Delete(seen).Error != nil {
		return errors.New("删除失败")
	}
	return nil
}

//CreateOrder ...
func (d *Dao) CreateOrder(userID, movieID, number, price int, movieName, date, startTime, endTime, cinemaName, address, hallName, seat string) (order *model.UserOrder, err error) {
	order = &model.UserOrder{
		UserID:     userID,
		MovieID:    movieID,
		MovieName:  movieName,
		Date:       date,
		StartTime:  startTime,
		EndTime:    endTime,
		CinemaName: cinemaName,
		Address:    address,
		HallName:   hallName,
		Number:     number,
		Seat:       seat,
		Price:      price,
	}
	return order, d.db.Create(order).Error
}

//DeleteOrder ...
func (d *Dao) DeleteOrder(orderID, userID int) (err error) {
	order := &model.UserOrder{OrderID: orderID, UserID: userID}
	if err = d.db.First(&order).Error; err != nil {
		return errors.New("未找到订单")
	}
	// if err = d.db.Where("order_id = ?", orderID).First(order).Error; err != nil {
	// 	return
	// }
	return d.db.Delete(order).Error
}

//UpdateEvaluation ...
func (d *Dao) UpdateEvaluation(userID, movieID, score int, review string) (oldScore int, seen *model.UserSeen, err error) {
	seen = &model.UserSeen{}
	if err = d.db.Where("user_id = ? and movie_id = ?", userID, movieID).First(seen).Error; err != nil {
		return 0, nil, errors.New("此电影未看过")
	}
	oldScore = seen.Score
	if err = d.db.Model(seen).Update("review", review).Error; err != nil {
		err = errors.New("更新评论失败")
	}
	if err = d.db.Model(seen).Update("score", score).Error; err != nil {
		err = errors.New("更新评分失败")
	}
	return
}

//UpdateScore ...
func (d *Dao) UpdateScore(movieID, score, flag int) error {
	scoreInfo := &model.ScoreInfo{}
	if d.db.Where("movie_id = ?", movieID).First(scoreInfo).Error != nil {
		return errors.New("此电影没有评分表")
	}
	switch score {
	case 1:
		d.db.Model(&scoreInfo).Update("score_one", scoreInfo.ScoreOne+flag)
		return nil
	case 2:
		d.db.Model(&scoreInfo).Update("score_two", scoreInfo.ScoreTwo+flag)
		return nil
	case 3:
		d.db.Model(&scoreInfo).Update("score_three", scoreInfo.ScoreThree+flag)
		return nil
	case 4:
		d.db.Model(&scoreInfo).Update("score_four", scoreInfo.ScoreFour+flag)
		return nil
	case 5:
		d.db.Model(&scoreInfo).Update("score_five", scoreInfo.ScoreFive+flag)
		return nil
	case 6:
		d.db.Model(&scoreInfo).Update("score_six", scoreInfo.ScoreSix+flag)
		return nil
	case 7:
		d.db.Model(&scoreInfo).Update("score_seven", scoreInfo.ScoreSeven+flag)
		return nil
	case 8:
		d.db.Model(&scoreInfo).Update("score_eight", scoreInfo.ScoreEight+flag)
		return nil
	case 9:
		d.db.Model(&scoreInfo).Update("score_nine", scoreInfo.ScoreNine+flag)
		return nil
	case 10:
		d.db.Model(&scoreInfo).Update("score_ten", scoreInfo.ScoreTen+flag)
		return nil
	default:
		return errors.New("分值不合法")
	}
}
