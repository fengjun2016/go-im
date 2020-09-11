package service

import (
	"go-im/app/model"

	"github.com/sirupsen/logrus"
)

type MessageService struct{}

func (m *MessageService) LoadToDb(msg model.Message) {
	err := msg.Create()
	if err != nil {
		logrus.Println("load msg to db failed.", err.Error())
	}
}

func (m *MessageService) SearchPersonalMessage(cmd int, userid, dstId string) (msgs []*model.Message, err error) {
	msg := model.Message{}
	msg.Dstid = dstId
	msg.Userid = userid
	msg.Cmd = cmd

	msgs = make([]*model.Message, 0)
	msgs, err = msg.SearchPersonalMessage()
	if err != nil {
		return
	}

	return
}

func (m *MessageService) SearchCommunityMessage(cmd int, dstid string) (msgs []*model.Message, err error) {
	msg := model.Message{}
	msg.Dstid = dstid
	msg.Cmd = cmd

	msgs = make([]*model.Message, 0)
	msgs, err = msg.SearchCommunityMessage()
	if err != nil {
		return
	}

	return
}
