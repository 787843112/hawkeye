package service

import (
	"errors"
	"web-server/model"
)

//AddHall ...
func (s *Service) AddHall(hallName string, cinemaID, hallRow, hallColumn int) (hall *model.HallInfo, err error) {
	hall, err = s.dao.CreateHall(hallName, cinemaID, hallRow, hallColumn)
	if  err != nil {
		return nil, errors.New("添加失败")
	}
	for r := 1; r <= hallRow; r++ {
		for c := 1; c <= hallColumn; c++ {
			go func(r, c int) {
				seat := &model.SeatInfo{HallID: hall.HallID, SeatRow: r, SeatColumn: c}
				s.dao.CreateSeat(seat)
			}(r, c)
		}
	}
	return
}

//DeleteHallInfo ...
func (s *Service) DeleteHallInfo(hallID int) (err error) {
	hall, err := s.dao.FindHallByHallID(hallID)
	if err != nil {
		return
	}
	//存在场次则无法删除影厅
	screenings, _ := s.dao.FindScreeningsByHallID(hallID)
	if len(screenings) != 0 {
		return errors.New("此影厅存在场次无法删除")
	}
	return s.dao.DeleteHall(hall)
}

//UpdateHallInfo ...
func (s *Service) UpdateHallInfo(hallName string, hallID, hallRow, hallColumn int) (hall *model.HallInfo, err error) {
	//该影厅存在场次则无法修改
	screenings, _ := s.dao.FindScreeningsByHallID(hallID)
	if len(screenings) != 0 {
		return nil, errors.New("此影厅存在场次无法修改")
	}
	oldRow, oldColumn, hall, err := s.dao.UpdateHall(hallName, hallID, hallRow, hallColumn)
	if err != nil {
		return 
	}
	if oldRow > hallRow && oldColumn > hallColumn {
		if err = s.dao.DeleteSeatRow(hallID, hallRow); err != nil {
			err = errors.New("删除行失败")
		}
		if err = s.dao.DeleteSeatColumn(hallID, hallColumn); err != nil {
			err = errors.New("删除列失败")
		}
	} else if oldRow < hallRow && oldColumn < hallColumn {
		for r := oldRow + 1; r <= hallRow; r++ {
			for c := oldColumn + 1; c <= hallColumn; c++ {
				go func(r, c int) {
					seat := &model.SeatInfo{HallID: hallID, SeatRow: r, SeatColumn: c}
					s.dao.CreateSeat(seat)
				}(r, c)
			}
		}
		for r := 1; r <= oldRow; r++ {
			for c := oldColumn + 1; c <= hallColumn; c++ {
				go func(r, c int) {
					seat := &model.SeatInfo{HallID: hallID, SeatRow: r, SeatColumn: c}
					s.dao.CreateSeat(seat)
				}(r, c)
			}
		}
		for r := oldRow + 1; r <= hallRow; r++ {
			for c := 1; c <= oldColumn; c++ {
				go func(r, c int) {
					seat := &model.SeatInfo{HallID: hallID, SeatRow: r, SeatColumn: c}
					s.dao.CreateSeat(seat)
				}(r, c)
			}
		}
	} else if oldRow < hallRow && oldColumn >= hallColumn {
		if err = s.dao.DeleteSeatColumn(hallID, hallColumn); err != nil {
			err = errors.New("删除列失败")
		}
		for r := oldRow + 1; r <= hallRow; r++ {
			for c := 1; c <= hallColumn; c++ {
				go func(r, c int) {
					seat := &model.SeatInfo{HallID: hallID, SeatRow: r, SeatColumn: c}
					s.dao.CreateSeat(seat)
				}(r, c)
			}
		}
	} else if oldRow >= hallRow && oldColumn < hallColumn {
		if err = s.dao.DeleteSeatRow(hallID, hallRow); err != nil {
			err = errors.New("删除行失败")
		}
		for r := 1; r <= hallRow; r++ {
			for c := oldColumn + 1; c <= hallColumn; c++ {
				go func(r, c int) {
					seat := &model.SeatInfo{HallID: hallID, SeatRow: r, SeatColumn: c}
					s.dao.CreateSeat(seat)
				}(r, c)
			}
		}
	}
	if err == nil {
		return hall, errors.New("修改成功")
	}
	return
}

//QueryHallInfo ...
func (s *Service) QueryHallInfo(hallName string, cinemaID, page, size int) (count int, halls []*model.HallInfo, err error) {
	return s.dao.FindHall(hallName, cinemaID, page, size)
}

//QueryHallByHallID ...
func (s *Service) QueryHallByHallID(hallID int) (halls *model.HallInfo, err error) {
	return s.dao.FindHallByHallID(hallID)
}
