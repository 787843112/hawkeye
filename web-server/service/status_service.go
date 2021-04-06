package service

import "web-server/model"

//InitializeSeatStatus ...初始化座位状态
func (s *Service) InitializeSeatStatus(hallID, screeningID int) (err error) {
	seats, err := s.dao.FindSeats(hallID)
	if err != nil {
		return
	}
	for _, v := range seats {
		status := &model.StatusInfo{ScreeningID: screeningID, SeatRow: v.SeatRow, SeatColumn: v.SeatColumn}
		if err = s.dao.CreateStatus(status); err != nil {
			return
		}
	}
	return
}

//UpdateStatusInfo ...
func (s *Service) UpdateStatusInfo(status *model.StatusInfo) error {
	return s.dao.UpdateStatus(status)
}

//RestoreStatusInfo ...
func (s *Service) RestoreStatusInfo(status []*model.StatusInfo) {
	for _, v := range status {
		s.dao.RestoreStatus(v)
	}
}

//QueryStatus ...
func (s *Service) QueryStatus(screeningID int) (hall *model.HallInfo, status []*model.StatusInfo, err error) {
	screening, _ := s.dao.FindScreeningByScreeningID(screeningID)
	hall, _ = s.dao.FindHallByHallID(screening.HallID)
	status, err = s.dao.FindStatus(screeningID)
	return hall, status, err
}
