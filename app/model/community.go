package model

import (
	"go-im/app/appInit"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

type Community struct {
	Id       string    `form:"id" json:"id"`
	Name     string    `form:"name" json:"name"`       //名称
	Ownerid  string    `form:"ownerid" json:"ownerid"` //群主ID
	Icon     string    `form:"icon" json:"icon"`       //群logo
	Cate     int       `form:"cate" json:"cate"`       //群的类型
	Memo     string    `form:"memo" json:"memo"`
	Createat time.Time `form:"createat" json:"createat"`
}

func (c *Community) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", xid.New().String())
}

func (c *Community) Create() (err error) {
	err = appInit.DB.
		Create(c).
		Error

	return
}

func (c *Community) GetByName() (err error) {
	err = appInit.DB.
		Where("name = ?", c.Name).
		First(c).
		Error

	return
}

func (c *Community) CountCommunitys() (count int, err error) {
	err = appInit.DB.
		Where("ownerid = ?", c.Ownerid).
		Count(&count).
		Error

	return
}

func FindInCommunitys(incommunitys []string) (coms []*Community, err error) {
	err = appInit.DB.
		Where("id in (?)", incommunitys).
		Find(&coms).
		Error

	return
}

func (c *Community) UserCreateCommunity() (err error) {
	tx := appInit.DB.Begin()

	err = tx.Create(c).
		Error

	if err != nil {
		tx.Rollback()
		return
	}

	//将自己加入到该联系人群关系当中去
	contact := Contact{
		Ownerid:  c.Ownerid,
		Dstobj:   c.Id,
		Cate:     ConcatCateComunity,
		Createat: time.Now(),
	}

	err = tx.Create(&contact).Error
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return nil
}
