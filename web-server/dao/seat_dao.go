package dao

import "web-server/model"

//CreateSeat ...
func (d *Dao) CreateSeat(seat *model.SeatInfo) error {
	return d.db.Create(seat).Error
}

//DeleteSeat ...
func (d *Dao) DeleteSeat(seat *model.SeatInfo) error {
	return d.db.Delete(seat).Error
}

//DeleteSeatRow ...
func (d *Dao) DeleteSeatRow(hallID, hallRow int) error {
	return d.db.Where("hall_id = ? AND seat_row > ?", hallID, hallRow).Delete(model.SeatInfo{}).Error
}

//DeleteSeatColumn ...
func (d *Dao) DeleteSeatColumn(hallID, hallColumn int) error {
	return d.db.Where("hall_id = ? AND seat_column > ?", hallID, hallColumn).Delete(model.SeatInfo{}).Error
}

//FindSeats ...
func (d *Dao) FindSeats(hallID int) (seats []*model.SeatInfo, err error) {
	return seats, d.db.Where("hall_id = ?", hallID).Find(&seats).Error
}

//FindSeat ...
func (d *Dao) FindSeat(seat *model.SeatInfo) error {
	return d.db.Find(&seat).Error
}
