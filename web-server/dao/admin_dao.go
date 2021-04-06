package dao

import (
	"errors"
	"web-server/model"
	"web-server/util"
)

//FindAdminByAdminID ...
func (d *Dao) FindAdminByAdminID(adminID int) (admin *model.AdminInfo, err error) {
	admin = &model.AdminInfo{}
	return admin, d.db.Where("admin_id = ?", adminID).First(admin).Error
}

//FindAdminByAdminName ...
func (d *Dao) FindAdminByAdminName(adminName string) (admin *model.AdminInfo, err error) {
	admin = &model.AdminInfo{}
	return admin, d.db.Where("admin_name = ?", adminName).First(admin).Error
}

//FindAllCinemaAdmin ...
func (d *Dao) FindAllCinemaAdmin(page, size int) (count int, admins []*model.AdminInfo, err error) {
	d.db.Model(model.AdminInfo{}).Where("permissions = 2").Count(&count)
	return count, admins, d.db.Where("permissions = 2").Limit(size).Offset((page - 1) * size).Find(&admins).Error
}

//CreateAdmin ...
func (d *Dao) CreateAdmin(adminName, password, phone string) (admin *model.AdminInfo, err error) {
	admin = &model.AdminInfo{AdminName: adminName, Password: password, Phone: phone, Permissions: 2}
	return admin, d.db.Create(admin).Error
}

//UpdateAdmin ...
func (d *Dao) UpdateAdmin(adminID int, password, phone string) (admin *model.AdminInfo, err error) {
	admin = &model.AdminInfo{}
	if err = d.db.Where("admin_id = ?", adminID).First(admin).Error; err != nil {
		return nil, errors.New("管理员不存在")
	}
	if password != "" {
		password = util.MD5V([]byte(password))
		admin.Password = password
	}
	if phone != "" {
		admin.Phone = phone
	}
	if err = d.db.Save(admin).Error; err != nil {
		return nil, errors.New("修改失败")
	}
	return
}

//UpdatePermissions ...
func (d *Dao) UpdatePermissions(adminID int) error {
	admin := &model.AdminInfo{}
	if d.db.Where("admin_id = ?", adminID).First(admin).Error != nil {
		return errors.New("管理员不存在")
	}
	if admin.Permissions == 1 {
		return errors.New("已是系统管理员")
	}
	admin.Permissions = 1
	if d.db.Save(admin).Error != nil {
		return errors.New("修改失败")
	}
	d.db.Model(model.ManageInfo{}).Delete("admin_id = ?", adminID)
	return nil
}

//DeleteAdmin ...
func (d *Dao) DeleteAdmin(admin *model.AdminInfo) error {
	return d.db.Delete(admin).Error
}

//FindManage ...
func (d *Dao) FindManage(manage *model.ManageInfo) error {
	return d.db.First(manage).Error
}

//AddManageInfo ...
func (d *Dao) AddManageInfo(manage *model.ManageInfo) error {
	return d.db.Create(manage).Error
}

//DeleteManageInfo ...
func (d *Dao) DeleteManageInfo(manage *model.ManageInfo) error {
	return d.db.Delete(manage).Error
}

//UpdateAvatarByAdmin ...
func (d *Dao) UpdateAvatarByAdmin(adminID int, avatar string) (admin *model.AdminInfo, err error) {
	admin = &model.AdminInfo{}
	if err = d.db.Where("admin_id = ?", adminID).First(admin).Error; err != nil {
		return nil, errors.New("管理员不存在")
	}
	admin.Avatar = avatar
	return admin, d.db.Save(admin).Error
}
