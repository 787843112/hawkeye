package service

import "web-server/model"

//AddSeat ...
func (s *Service) AddSeat(seat *model.SeatInfo) error {
	return s.dao.CreateSeat(seat)
}

//DeleteSeatInfo ...
func (s *Service) DeleteSeatInfo(seat *model.SeatInfo) error {
	return s.dao.DeleteSeat(seat)
}

//QuerySeat ...
func (s *Service) QuerySeat(seat *model.SeatInfo) error {
	return s.dao.FindSeat(seat)
}

//QuerySeatInfo ...
func (s *Service) QuerySeatInfo(hallID int) (seats []*model.SeatInfo, err error) {
	return s.dao.FindSeats(hallID)
}
