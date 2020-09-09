package model

import (
	"go-im/app/appInit"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

//好友和群都存在这个表里面
//可根据具体业务做拆分
type Contact struct {
	Id       string    `form:"id" json:"id"`
	Ownerid  string    `form:"ownerid" json:"ownerid"`   // 记录是谁的
	Dstobj   string    `form:"dstobj" json:"dstobj"`     // 对端信息
	Cate     int       `form:"cate" json:"cate"`         // 什么类型
	Memo     string    `form:"memo" json:"memo"`         // 备注
	Createat time.Time `form:"createat" json:"createat"` // 创建时间
}

func (c *Contact) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", xid.New().String())
}

func (c *Contact) Create() (err error) {
	err = appInit.DB.
		Create(c).
		Error

	return
}

func (c *Contact) Get() (err error) {
	err = appInit.DB.
		Model(c).
		Where("id = ?", c.Id).
		First(c).
		Error

	return
}

func (c *Contact) Update(where, update map[string]interface{}) (err error) {
	err = appInit.DB.
		Model(c).
		Where("id = ?", c.Id).
		Where(where).
		Updates(update).
		Error

	return
}

func CheckFriendRelationShip(ownId, dstId string) (friend *Contact, err error) {
	friend = new(Contact)
	err = appInit.DB.
		Where("ownerid = ?", ownId).
		Where("dstobj = ?", dstId).
		Where("cate = ?", ConcatCateUser).
		First(friend).
		Error

	return
}

func CheckCommunityRelationShip(ownId, comId string) (com *Contact, err error) {
	err = appInit.DB.
		Where("ownerid = ?", ownId).
		Where("dstobj = ?", comId).
		Where("cate = ?", ConcatCateComunity).
		First(com).
		Error

	return
}

func CreatePersonalFriendRelationShip(userId, dstId string) (err error) {
	tx := appInit.DB.Begin()
	//插入两条好友关系数据
	contact := Contact{
		Ownerid:  userId,
		Dstobj:   dstId,
		Cate:     ConcatCateUser,
		Createat: time.Now(),
	}

	otherContact := Contact{
		Ownerid:  dstId,
		Dstobj:   userId,
		Cate:     ConcatCateUser,
		Createat: time.Now(),
	}

	err = tx.Create(&contact).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Create(&otherContact).Error
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (c *Contact) SearchFriends() (friends []*Contact, err error) {
	friends = make([]*Contact, 0)
	err = appInit.DB.
		Where("ownerid = ? and cate = ?", c.Ownerid, ConcatCateUser).
		Find(&friends).
		Error

	return
}

func (c *Contact) SearchCommunitys() (coms []*Contact, err error) {
	coms = make([]*Contact, 0)

	err = appInit.DB.
		Where("ownerid = ? and cate = ?", c.Ownerid, ConcatCateComunity).
		Find(&coms).
		Error

	return
}

func (c *Contact) List(rawQuery string, rawOrder string, offset int, limit int) ([]*Contact, int, error) {
	contatcs := make([]*Contact, 0)
	total := 0

	db := appInit.DB.Model(&Contact{})

	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return contatcs, total, err
	}

	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return contatcs, total, err
	}

	db.Order("createat desc").
		Count(&total).
		Offset(offset).
		Limit(limit).
		Find(&contatcs)

	err = db.Error

	return contatcs, total, err
}
