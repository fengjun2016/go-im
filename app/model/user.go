package model

import (
	"errors"
	"go-im/app/appInit"
)

type User struct {
	Id       int64  `gorm:"pk autoincr bigint(64)" form:"id" json:"id"`
	Mobile   string `gorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd   string `gorm:"varchar(40)" form:"passwd" json:"-"` // 用户密码 md5(passwd + salt)
	Avatar   string `gorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex      string `gorm:"varchar(2)" form:"sex" json:"sex"`
	Nickname string `gorm:"varchar(20)" form:"nickname" json:"nickname"`
	Salt     string `gorm:"varchar(10)" form:"salt" json:"-"`
	Online   int    `gorm:"int(10)" form:"online" json:"online"`   //是否在线
	Token    string `gorm:"varchar(40)" form:"token" json:"token"` //用户鉴权
	Memo     string `gorm:"varchar(140)" form:"memo" json:"memo"`
}

//检查是否已注册过该手机号
func (u *User) ExistUserByMobile() (exists bool, err error) {
	if u.Mobile == "" {
		err = errors.New("mobile can not empty")
		return
	}
	err = appInit.DB.
		Model(&User{}).
		Where("mobile = ?", u.Mobile).
		First(u).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		exists = false
		return
	}

	if u.ID != "" {
		exists = true
		err = nil
		return
	}

	return false, nil
}

func (u *User) Create() (err error) {
	err = appInit.DB.
		Create(u).
		Error

	return
}

func (u *User) Get() (err error) {
	err = appInit.DB.
		Model(u).
		Where("id = ?", u.Id)
	First(u).
		Error
	return
}
