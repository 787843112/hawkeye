package service

import (
	"web-server/model"
)

//AddCinema ...
func (s *Service) AddCinema(cinemaName, city, address string) (cinema *model.CinemaInfo, err error) {
	return s.dao.CreateCinema(cinemaName, city, address)
}

//DeleteCinemaInfo ...
func (s *Service) DeleteCinemaInfo(cinemaID int) (err error) {
	cinema, err := s.dao.FindCinemaByCinemaID(cinemaID)
	if err != nil {
		return
	}
	return s.dao.DeleteCinema(cinema)
}

//UpdateCinemaInfo ...
func (s *Service) UpdateCinemaInfo(cinemaID int, cinemaName, city, address string) (cinema *model.CinemaInfo, err error) {
	return s.dao.UpdateCinema(cinemaID, cinemaName, city, address)
}

//QueryCinemaByCinemaID ...
func (s *Service) QueryCinemaByCinemaID(cinemaID int) (cinema *model.CinemaInfo, err error) {
	return s.dao.FindCinemaByCinemaID(cinemaID)
}

//QueryCinemaInfo ...
func (s *Service) QueryCinemaInfo(city, cinemaName string, page, size int) (count int, cinemas []*model.CinemaInfo, err error) {
	return s.dao.FindCinema(city, cinemaName, page, size)
}

//QueryCinemaInfoByCity ...
func (s *Service) QueryCinemaInfoByCity(city string) (cinemas []*model.CinemaInfo, err error) {
	return s.dao.FindCinemaByCity(city)
}
