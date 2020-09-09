package model

import (
	"errors"
	"go-im/app/appInit"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

type User struct {
	Id       string    `form:"id" json:"id"`
	Mobile   string    `form:"mobile" json:"mobile"`
	Passwd   string    `form:"passwd" json:"-"` // 用户密码 md5(passwd + salt)
	Avatar   string    `form:"avatar" json:"avatar"`
	Sex      string    `form:"sex" json:"sex"`
	Nickname string    `form:"nickname" json:"nickname"`
	Salt     string    `form:"salt" json:"-"`
	Online   int       `form:"online" json:"online"` //是否在线
	Token    string    `form:"token" json:"token"`   //用户鉴权
	Memo     string    `form:"memo" json:"memo"`
	Createat time.Time `form:"createat" json:"createat"` // 创建时间
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", xid.New().String())
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

	if u.Id != "" {
		exists = true
		err = nil
		return
	}

	return false, nil
}

func (u *User) Create() (err error) {
	err = appInit.DB.
		Model(&User{}).
		Create(u).
		Error

	return
}

func (u *User) Get() (err error) {
	err = appInit.DB.
		Model(&User{}).
		Where("id = ?", u.Id).
		First(u).
		Error
	return
}

func (u *User) GetByName() (err error) {
	err = appInit.DB.
		Model(&User{}).
		Where("mobile = ?", u.Mobile).
		First(u).
		Error

	return
}

func (u *User) Update(where, update map[string]interface{}) (err error) {
	err = appInit.DB.
		Model(&User{}).
		Where("id = ?", u.Id).
		Where(where).
		Updates(update).
		Error

	return
}

func FindInUsers(inUsers []string) (users []*User, err error) {
	users = make([]*User, 0)
	err = appInit.DB.
		Model(&User{}).
		Where("id in (?)", inUsers).
		Find(&users).
		Error

	return
}

func (u *User) List(rawQuery string, rawOrder string, offset int, limit int) ([]*User, int, error) {
	users := make([]*User, 0)
	total := 0

	db := appInit.DB.Model(&User{})

	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return users, total, err
	}

	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return users, total, err
	}

	db.Order("createat desc").
		Count(&total).
		Offset(offset).
		Limit(limit).
		Find(&users)

	err = db.Error

	return users, total, err
}
