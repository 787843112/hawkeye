package dao

import (
	"time"
	"web-server/model"
)

//CreateScreening ...
func (d *Dao) CreateScreening(hallID, movieID, price int, hallName string, startTime, endTime time.Time) (screening *model.ScreeningInfo, err error) {
	screening = &model.ScreeningInfo{HallID: hallID, HallName: hallName, MovieID: movieID, StartTime: startTime, EndTime: endTime, Price: price}
	return screening, d.db.Create(screening).Error
}

//DeleteScreening ...
func (d *Dao) DeleteScreening(screening *model.ScreeningInfo) error {
	return d.db.Delete(screening).Error
}

//FindScreeningsByHallID ...
func (d *Dao) FindScreeningsByHallID(hallID int) (screenings []*model.ScreeningInfo, err error) {
	return screenings, d.db.Where("hall_id = ?", hallID).Order("start_time").Find(&screenings).Error
}

//AdminFindScreenings ...
func (d *Dao) AdminFindScreenings(hallID, page, size int, startDate, endDate time.Time) (count int, screenings []*model.ScreeningInfo, err error) {
	d.db.Model(model.ScreeningInfo{}).Where("hall_id = ? AND start_time >= ? AND end_time <= ?", hallID, startDate, endDate).Count(&count)
	return count, screenings, d.db.Where("hall_id = ? AND start_time >= ? AND end_time <= ?", hallID, startDate, endDate).Limit(size).Offset((page - 1) * size).Order("start_time").Find(&screenings).Error
}

//AdminFindAllScreenings ...
func (d *Dao) AdminFindAllScreenings(hallID, page, size int) (count int, screenings []*model.ScreeningInfo, err error) {
	d.db.Model(model.ScreeningInfo{}).Where("hall_id = ?", hallID).Count(&count)
	return count, screenings, d.db.Where("hall_id = ?", hallID).Limit(size).Offset((page - 1) * size).Order("start_time").Find(&screenings).Error
}

//FindScreeningByScreeningID ...
func (d *Dao) FindScreeningByScreeningID(screeningID int) (screening *model.ScreeningInfo, err error) {
	screening = &model.ScreeningInfo{}
	return screening, d.db.Where("screening_id = ?", screeningID).First(screening).Error
}

//UserFindScreenings ...
func (d *Dao) UserFindScreenings(cinemaID, movieID, page, size int, startDate, endDate time.Time) (count int, screenings []*model.ScreeningInfo, err error) {
	d.db.Model(model.ScreeningInfo{}).Joins("JOIN hall_info ON screening_info.hall_id = hall_info.hall_id").Where("hall_info.cinema_id = ? AND movie_id = ? AND start_time >= ? AND end_time <= ?", cinemaID, movieID, startDate, endDate).Count(&count)
	return count, screenings, d.db.Joins("JOIN hall_info ON screening_info.hall_id = hall_info.hall_id").Where("hall_info.cinema_id = ? AND movie_id = ? AND start_time >= ? AND end_time <= ?", cinemaID, movieID, startDate, endDate).Limit(size).Offset((page - 1) * size).Order("start_time").Find(&screenings).Error
}
