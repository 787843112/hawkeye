package service

import (
	"errors"
	"web-server/model"
)

//QueryUserInfoByUserID ...
func (s *Service) QueryUserInfoByUserID(userID int) (user *model.UserInfo, err error) {
	return s.dao.FindUserByUserID(userID)
}

//QueryUserInfoByUserName ...
func (s *Service) QueryUserInfoByUserName(username string) (user *model.UserInfo, err error) {
	return s.dao.FindUserByUserName(username)
}

//RegisterUser ...
func (s *Service) RegisterUser(userName, password, phone string) (user *model.UserInfo, err error) {
	return s.dao.CreateUser(userName, password, phone)

}

//UpdateUserInfo ...
func (s *Service) UpdateUserInfo(userID int, nickname, password, phone string) (user *model.UserInfo, err error) {
	return s.dao.UpdateUser(userID, nickname, password, phone)
}

//UpdateUserAvatar ...
func (s *Service) UpdateUserAvatar(userID int, avatar string) (user *model.UserInfo, err error) {
	return s.dao.UpdateAvatarByUser(userID, avatar)
}

//QueryUserOrder ...
func (s *Service) QueryUserOrder(userID, page, size int) (count int, orders []*model.UserOrder, err error) {
	return s.dao.FindUserOrder(userID, page, size)
}

//QueryUserWant ...
func (s *Service) QueryUserWant(userID, page, size int) (count int, movies []*model.MovieInfo, err error) {
	return s.dao.FindUserWant(userID, page, size)
}

//QueryUserWantMovieID ...
func (s *Service) QueryUserWantMovieID(userID int) (movieIDS []int, err error) {
	return s.dao.FindUserWantMovieID(userID)
}

//QueryUserSeen ...
func (s *Service) QueryUserSeen(userID, page, size int) (count int, seens []*model.UserSeen, err error) {
	count, seens, err = s.dao.FindUserSeen(userID, page, size)
	for _, v := range seens {
		v.MovieInfo, _ = s.dao.FindMovieByMovieID(v.MovieID)
	}
	return count, seens, err
}

//QueryUserSeenByMovieID ...
func (s *Service) QueryUserSeenByMovieID(userID, movieID int) (seen *model.UserSeen, err error) {
	seen, err = s.dao.FindSeen(userID, movieID)
	if err != nil {
		return
	}
	seen.MovieInfo, _ = s.dao.FindMovieByMovieID(movieID)
	return
}

//AddWantMovie ...
func (s *Service) AddWantMovie(userID, movieID int) (err error) {
	want := &model.UserWant{UserID: userID, MovieID: movieID}
	if _, err = s.dao.FindUserByUserID(userID); err != nil {
		return errors.New("用户不存在")
	}
	if _, err = s.dao.FindMovieByMovieID(movieID); err != nil {
		return errors.New("电影不存在")
	}
	if err = s.dao.FindWant(want); err == nil {
		return errors.New("此电影已想看")
	}
	if err = s.dao.CreateUserWant(want); err != nil {
		return errors.New("添加失败")
	}
	return
}

//DeleteWantMovie ...
func (s *Service) DeleteWantMovie(userID, movieID int) (err error) {
	want := &model.UserWant{UserID: userID, MovieID: movieID}
	if err = s.dao.FindWant(want); err != nil {
		return errors.New("电影未在想看")
	}
	if err = s.dao.DeleteUserWant(want); err != nil {
		err = errors.New("删除失败")
	}
	return
}

//DeleteSeenMovie ...
func (s *Service) DeleteSeenMovie(userID, movieID int) (err error) {
	flag := -1
	seen := &model.UserSeen{}
	if seen, err = s.dao.FindSeen(userID, movieID); err != nil {
		return errors.New("此电影未看过")
	}
	if seen.Score != 0 {
		s.dao.UpdateScore(movieID, seen.Score, flag)
	}
	return s.dao.DeleteUserSeen(seen)
}

//AddUserOrder ...
func (s *Service) AddUserOrder(userID, movieID, number, price int, movieName, date, startTime, endTime, cinemaName, address, hallName, seat string) (order *model.UserOrder, err error) {
	return s.dao.CreateOrder(userID, movieID, number, price, movieName, date, startTime, endTime, cinemaName, address, hallName, seat)
}

//DeleteUserOrder ...
func (s *Service) DeleteUserOrder(orderID, userID int) error {
	return s.dao.DeleteOrder(orderID, userID)
}

//UpdateMovieEvaluation ...
func (s *Service) UpdateMovieEvaluation(userID, movieID, score int, review string) (seen *model.UserSeen, err error) {
	oldScore, seen, err := s.dao.UpdateEvaluation(userID, movieID, score, review)
	if seen == nil {
		return
	}
	flag := -1
	//判断是否有过评分
	if oldScore != 0 {
		s.dao.UpdateScore(movieID, oldScore, flag)
	}
	flag = 1
	s.dao.UpdateScore(movieID, score, flag)
	return seen, errors.New("更新成功")
}

//DeleteMovieEvaluation ...
func (s *Service) DeleteMovieEvaluation(userID, movieID int) (seen *model.UserSeen, err error) {
	score, flag := 0, -1
	review := ""
	oldScore, seen, err := s.dao.UpdateEvaluation(userID, movieID, score, review)
	if seen == nil {
		return
	}
	err = s.dao.UpdateScore(movieID, oldScore, flag)
	if err == nil {
		return seen, errors.New("删除成功")
	}
	return
}
