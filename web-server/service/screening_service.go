package service

import (
	"errors"
	"time"
	"web-server/model"
)

var oneday, _ = time.ParseDuration("24h")

//AddScreening ...
func (s *Service) AddScreening(hallID, movieID, price int, startTime, endTime time.Time) (screening *model.ScreeningInfo, err error) {
	hall, _ := s.dao.FindHallByHallID(hallID)
	screenings, err := s.dao.FindScreeningsByHallID(hallID)
	if err != nil {
		return nil, errors.New("查询场次失败")
	}
	if len(screenings) == 0 {
		if screening, err = s.dao.CreateScreening(hallID, movieID, price, hall.HallName, startTime, endTime); err != nil {
			return nil, errors.New("添加失败")
		}
		return
	}
	if !endTime.After(screenings[0].StartTime) {
		if screening, err = s.dao.CreateScreening(hallID, movieID, price, hall.HallName, startTime, endTime); err != nil {
			return nil, errors.New("添加失败")
		}
		return
	} else if !startTime.Before(screenings[len(screenings)-1].EndTime) {
		if screening, err = s.dao.CreateScreening(hallID, movieID, price, hall.HallName, startTime, endTime); err != nil {
			return nil, errors.New("添加失败")
		}
		return
	} else {
		for i := len(screenings) - 1; i > 0; i-- {
			if !startTime.Before(screenings[i-1].EndTime) && !endTime.After(screenings[i].StartTime) {
				if screening, err = s.dao.CreateScreening(hallID, movieID, price, hall.HallName, startTime, endTime); err != nil {
					return nil, errors.New("添加失败")
				}
				return
			}
		}
		return nil, errors.New("此时间段已被占用")
	}
}

//DeleteScreeningInfo ...
func (s *Service) DeleteScreeningInfo(screeningID int) error {
	screening := &model.ScreeningInfo{ScreeningID: screeningID}
	return s.dao.DeleteScreening(screening)
}

//AdminQueryScreenings ...
func (s *Service) AdminQueryScreenings(hallID, page, size int, startDate, endDate time.Time) (count int, screenings []*model.ScreeningInfo, err error) {
	// if startDate == endDate {
	endDate = endDate.Add(oneday)
	// }
	return s.dao.AdminFindScreenings(hallID, page, size, startDate, endDate)
}

//AdminQueryAllScreenings ...
func (s *Service) AdminQueryAllScreenings(hallID, page, size int) (count int, screenings []*model.ScreeningInfo, err error) {
	return s.dao.AdminFindAllScreenings(hallID, page, size)
}

//UserQueryScreenings ...
func (s *Service) UserQueryScreenings(cinemaID, movieID, page, size int, selectDate time.Time) (count int, screenings []*model.ScreeningInfo, err error) {
	var (
		endDate    = selectDate.Add(oneday)
		timeLayout = "2006-01-02"
		today, _   = time.ParseInLocation(timeLayout, time.Now().Format(timeLayout), time.Local)
	)
	if selectDate == today {
		selectDate = time.Now()
	}
	return s.dao.UserFindScreenings(cinemaID, movieID, page, size, selectDate, endDate)
}

//QueryScreeningByHallID ...
func (s *Service) QueryScreeningByHallID(hallID int) error {
	screening, _ := s.dao.FindScreeningsByHallID(hallID)
	if len(screening) == 0 {
		return nil
	}
	return errors.New("影厅存在场次无法修改")
}
