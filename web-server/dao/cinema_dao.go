package dao

import (
	"errors"
	"web-server/model"
)

//CreateCinema ...
func (d *Dao) CreateCinema(cinemaName, city, address string) (cinema *model.CinemaInfo, err error) {
	cinema = &model.CinemaInfo{CinemaName: cinemaName, City: city, Address: address}
	return cinema, d.db.Create(cinema).Error
}

//DeleteCinema ...
func (d *Dao) DeleteCinema(cinema *model.CinemaInfo) error {
	return d.db.Delete(cinema).Error
}

//UpdateCinema ...
func (d *Dao) UpdateCinema(cinemaID int, cinemaName, city, address string) (cinema *model.CinemaInfo, err error) {
	cinema = &model.CinemaInfo{}
	if err = d.db.Where("cinema_id = ?", cinemaID).First(cinema).Error; err != nil {
		return nil, errors.New("影院不存在")
	}
	cinema.CinemaName = cinemaName
	cinema.City = city
	cinema.Address = address
	if err = d.db.Save(cinema).Error; err != nil {
		return nil, errors.New("修改失败")
	}
	return
}

//FindCinemaByCinemaID ...
func (d *Dao) FindCinemaByCinemaID(cinemaID int) (cinema *model.CinemaInfo, err error) {
	cinema = &model.CinemaInfo{}
	return cinema, d.db.Where("cinema_id = ?", cinemaID).First(cinema).Error
}

//FindCinema ...
func (d *Dao) FindCinema(city, cinemaName string, page, size int) (count int, cinemas []*model.CinemaInfo, err error) {
	if city == "" {
		d.db.Model(model.CinemaInfo{}).Where("cinema_name LIKE ?", "%"+cinemaName+"%").Count(&count)
		return count, cinemas, d.db.Where("cinema_name LIKE ?", "%"+cinemaName+"%").Limit(size).Offset((page - 1) * size).Order("cinema_id desc").Find(&cinemas).Error
	}
	d.db.Model(model.CinemaInfo{}).Where("cinema_name LIKE ? AND city = ?", "%"+cinemaName+"%", city).Count(&count)
	return count, cinemas, d.db.Where("cinema_name LIKE ? AND city = ?", "%"+cinemaName+"%", city).Limit(size).Offset((page - 1) * size).Order("cinema_id desc").Find(&cinemas).Error
}

//FindCinemaByCity ...
func (d *Dao) FindCinemaByCity(city string) (cinemas []*model.CinemaInfo, err error) {
	return cinemas, d.db.Where("city = ?", city).Find(&cinemas).Error
}

//FindAminManageCinemas ...
func (d *Dao) FindAminManageCinemas(adminID, page, size int, city, cinemaName string) (count int, cinemas []*model.CinemaInfo, err error) {
	admin := &model.AdminInfo{}
	if err = d.db.Where("admin_id = ?", adminID).First(admin).Error; err != nil {
		return 0, nil, errors.New("管理员不存在")
	}
	if city == "" {
		d.db.Model(&admin).Where("cinema_name LIKE ?", "%"+cinemaName+"%").Related(&cinemas, "Cinemas")
		count = len(cinemas)
		if d.db.Model(&admin).Where("cinema_name LIKE ?", "%"+cinemaName+"%").Limit(size).Offset((page-1)*size).Order("create_time desc").Related(&cinemas, "Cinemas").Error != nil {
			return 0, nil, errors.New("获取失败")
		}
	} else {
		d.db.Model(&admin).Where("cinema_name LIKE ? AND city = ?", "%"+cinemaName+"%", city).Related(&cinemas, "Cinemas")
		count = len(cinemas)
		if d.db.Model(&admin).Where("cinema_name LIKE ? AND city = ?", "%"+cinemaName+"%", city).Limit(size).Offset((page-1)*size).Order("create_time desc").Related(&cinemas, "Cinemas").Error != nil {
			return 0, nil, errors.New("获取失败")
		}
	}
	return
}
