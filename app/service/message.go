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
