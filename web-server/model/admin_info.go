package model

//AdminInfo ...
type AdminInfo struct {
	AdminID     int    `gorm:"primary_key;AUTO_INCREMENT;comment:'管理员ID';not null" json:"AdminID" form:"AdminID"`
	AdminName   string `gorm:"type:varchar(20);unique;comment:'管理员名';not null" json:"AdminName"`
	Password    string `gorm:"type:varchar(32);comment:'管理员密码';not null" json:"-"`
	Avatar      string `gorm:"default:'defaultAvatar.png';comment:'头像'" json:"Avatar"`
	Phone       string `gorm:"type:varchar(20);comment:'手机号';not null" json:"Phone"`
	Permissions int    `gorm:"type:tinyint(1);default:2;comment:'权限等级';not null" json:"Permissions" form:"Permissions"`
	Time        `json:"-"`
	Cinemas     []*CinemaInfo `gorm:"many2many:manage_info;ForeignKey:AdminID;AssociationForeignKey:CinemaID;jointable_foreignkey:admin_id;association_jointable_foreignkey:cinema_id" json:"-"`
}
