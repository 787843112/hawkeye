package dao

import (
	"errors"
	"web-server/model"
)

//CreateHall ...
func (d *Dao) CreateHall(hallName string, cinemaID, hallRow, hallColumn int) (hall *model.HallInfo, err error) {
	hall = &model.HallInfo{CinemaID: cinemaID, HallName: hallName, HallRow: hallRow, HallColumn: hallColumn}
	return hall, d.db.Create(hall).Error
}

//DeleteHall ...
func (d *Dao) DeleteHall(hall *model.HallInfo) (err error) {
	return d.db.Delete(hall).Error
}

//UpdateHall ...
func (d *Dao) UpdateHall(hallName string, hallID, hallRow, hallColumn int) (oldRow, oldColumn int, hall *model.HallInfo, err error) {
	hall = &model.HallInfo{}
	if err = d.db.Where("hall_id = ?", hallID).First(hall).Error; err != nil {
		return 0, 0, nil, errors.New("影厅不存在")
	}
	oldRow = hall.HallRow
	oldColumn = hall.HallColumn
	hall.HallName = hallName
	hall.HallRow = hallRow
	hall.HallColumn = hallColumn
	if err = d.db.Save(hall).Error; err != nil {
		return 0, 0, nil, errors.New("修改失败")
	}
	return
}

//FindHallByHallID ...
func (d *Dao) FindHallByHallID(hallID int) (hall *model.HallInfo, err error) {
	hall = &model.HallInfo{}
	return hall, d.db.Where("hall_id = ?", hallID).First(hall).Error
}

//FindHall ...
func (d *Dao) FindHall(hallName string, cinemaID, page, size int) (count int, halls []*model.HallInfo, err error) {
	d.db.Model(model.HallInfo{}).Where("cinema_id = ? AND hall_name LIKE ?", cinemaID, "%"+hallName+"%").Count(&count)
	return count, halls, d.db.Where("cinema_id = ? AND hall_name LIKE ?", cinemaID, "%"+hallName+"%").Limit(size).Offset((page - 1) * size).Order("hall_id desc").Find(&halls).Error
}
