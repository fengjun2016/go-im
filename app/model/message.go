package model

import (
	"go-im/app/appInit"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

//消息的持久化存储
type Message struct {
	Id       string    `json:"id,omitempty" form:"id"`           //消息ID
	Userid   string    `json:"userid,omitempty" form:"userid"`   //谁发的
	Cmd      int       `json:"cmd,omitempty" form:"cmd"`         //群聊还是私聊
	Dstid    string    `json:"dstid,omitempty" form:"dstid"`     //对端用户ID/群ID
	Media    int       `json:"media,omitempty" form:"media"`     //消息按照什么样式展示
	Content  string    `json:"content,omitempty" form:"content"` //消息的内容
	Pic      string    `json:"pic,omitempty" form:"pic"`         //预览图片
	Url      string    `json:"url,omitempty" form:"url"`         //服务的URL
	Memo     string    `json:"memo,omitempty" form:"memo"`       //简单描述
	Amount   int       `json:"amount,omitempty" form:"amount"`   //其他和数字相关的
	Createat time.Time `json:"createat"`                         //创建时间
}

func (m *Message) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", xid.New().String())
}

func (m *Message) Create() (err error) {
	err = appInit.DB.
		Model(m).
		Create(m).
		Error

	return
}

func (m *Message) SearchPersonalMessage() ([]*Message, error) {
	msgs := make([]*Message, 0)

	err := appInit.DB.Model(&Message{}).
		Where("cmd = ?", m.Cmd).
		Where("dstid = ?", m.Dstid).
		Order("createat desc").
		Find(&msgs).
		Error

	return msgs, err
}

func (m *Message) List(rawQuery string, rawOrder string, offset int, limit int) ([]*Message, int, error) {
	msgs := make([]*Message, 0)
	total := 0

	db := appInit.DB.Model(&Message{})

	db, err := buildWhere(rawQuery, db)
	if err != nil {
		return msgs, total, err
	}

	db, err = buildOrder(rawOrder, db)
	if err != nil {
		return msgs, total, err
	}

	db.Order("createat desc").
		Count(&total).
		Offset(offset).
		Limit(limit).
		Find(&msgs)

	err = db.Error

	return msgs, total, err
}
