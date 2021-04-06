package service

import (
	"errors"
	"web-server/model"
)

//QueryAdminInfoByAdminID ...
func (s *Service) QueryAdminInfoByAdminID(adminID int) (admin *model.AdminInfo, err error) {
	return s.dao.FindAdminByAdminID(adminID)

}

//QueryAdminInfoByAdminName ...
func (s *Service) QueryAdminInfoByAdminName(adminName string) (admin *model.AdminInfo, err error) {
	return s.dao.FindAdminByAdminName(adminName)
}

//QueryAllCinemaAdmin ...
func (s *Service) QueryAllCinemaAdmin(page, size int) (count int, admins []*model.AdminInfo, err error) {
	return s.dao.FindAllCinemaAdmin(page, size)
}

//RegisterAdmin ...
func (s *Service) RegisterAdmin(adminName, password, phone string) (admin *model.AdminInfo, err error) {
	return s.dao.CreateAdmin(adminName, password, phone)
}

//UpdateAdminInfo ...
func (s *Service) UpdateAdminInfo(adminID int, password, phone string) (admin *model.AdminInfo, err error) {
	return s.dao.UpdateAdmin(adminID, password, phone)
}

//UpdateAdminPermissions ...
func (s *Service) UpdateAdminPermissions(adminID int) error {
	return s.dao.UpdatePermissions(adminID)
}

//DeleteAdminInfo ...
func (s *Service) DeleteAdminInfo(adminID int) (err error) {
	admin, err := s.dao.FindAdminByAdminID(adminID)
	if err != nil || admin.Permissions == 1 {
		return
	}
	return s.dao.DeleteAdmin(admin)
}

//QueryManageCinema ...
func (s *Service) QueryManageCinema(adminID, page, size int, city, CinemaName string) (count int, cinemas []*model.CinemaInfo, err error) {
	return s.dao.FindAminManageCinemas(adminID, page, size, city, CinemaName)
}

//AddManage ...
func (s *Service) AddManage(adminID, cinemaID int) (cienma *model.CinemaInfo, err error) {
	if _, err = s.dao.FindAdminByAdminID(adminID); err != nil {
		return nil, errors.New("管理员不存在")
	}
	if _, err = s.dao.FindCinemaByCinemaID(cinemaID); err != nil {
		return nil, errors.New("影院不存在")
	}
	manage := &model.ManageInfo{AdminID: adminID, CinemaID: cinemaID}
	if err = s.dao.FindManage(manage); err == nil {
		return nil, errors.New("此影院已在管理")
	}
	if err = s.dao.AddManageInfo(manage); err != nil {
		return nil, errors.New("写入数据库失败")
	}
	return s.dao.FindCinemaByCinemaID(cinemaID)
}

//FindManageInfo ...
func (s *Service) FindManageInfo(adminID, cinemaID int) error {
	manage := &model.ManageInfo{AdminID: adminID, CinemaID: cinemaID}
	return s.dao.FindManage(manage)
}

//DeleteManage ...
func (s *Service) DeleteManage(adminID, cinemaID int) (err error) {
	manage := &model.ManageInfo{AdminID: adminID, CinemaID: cinemaID}
	if s.dao.FindManage(manage) != nil {
		return
	}
	return s.dao.DeleteManageInfo(manage)
}

//UpdateAdminAvatar ...
func (s *Service) UpdateAdminAvatar(adminID int, avatar string) (admin *model.AdminInfo, err error) {
	return s.dao.UpdateAvatarByAdmin(adminID, avatar)
}
