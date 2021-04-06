package model

//UserInfo ...
type UserInfo struct {
	UserID   int    `gorm:"primary_key;AUTO_INCREMENT;comment:'用户ID';not null" json:"UserID"`
	UserName string `gorm:"type:varchar(20);unique;comment:'用户名';not null" json:"UserName"`
	Nickname string `gorm:"type:varchar(10);comment:'用户昵称'" json:"Nickname"`
	Password string `gorm:"type:varchar(32);comment:'密码';not null" json:"-"`
	Phone    string `gorm:"type:varchar(20);comment:'手机号';not null" json:"Phone"`
	Avatar   string `gorm:"default:'defaultAvatar.png';comment:'用户头像'" json:"Avatar"`
	Time
	UserOrders []*UserOrder `gorm:"ForeignKey:UserID" json:"-"`
	WantMovies []*MovieInfo `gorm:"many2many:user_want;ForeignKey:UserID;AssociationForeignKey:MovieID;jointable_foreignkey:user_id;association_jointable_foreignkey:movie_id" json:"-"`
	UserSeens  []*UserSeen  `gorm:"ForeignKey:UserID" json:"-"`
	// SeenMovies []*MovieInfo `gorm:"many2many:user_seen;ForeignKey:UserID;AssociationForeignKey:MovieID;jointable_foreignkey:user_id;association_jointable_foreignkey:movie_id" json:"-"`
}
