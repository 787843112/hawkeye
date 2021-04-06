package dao

import (
	"errors"
	"web-server/model"
)

//CreateStatus ...
func (d *Dao) CreateStatus(status *model.StatusInfo) error {
	return d.db.Create(status).Error
}

//UpdateStatus ...
func (d *Dao) UpdateStatus(status *model.StatusInfo) (err error) {
	if err = d.db.First(&status).Error; err != nil {
		return
	}
	if status.Status == 1 {
		return errors.New("此座位已出售")
	}
	status.Status = 1
	d.db.Save(status)
	return
}

//RestoreStatus ...
func (d *Dao) RestoreStatus(status *model.StatusInfo) {
	d.db.First(&status)
	status.Status = 0
	d.db.Save(status)
}

//FindStatus ...
func (d *Dao) FindStatus(screeningID int) (status []*model.StatusInfo, err error) {
	return status, d.db.Where("screening_id = ?", screeningID).Find(&status).Error
}
