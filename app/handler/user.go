package handler

import (
	"go-im/app/model"
	"go-im/app/service"
	"go-im/app/util"
	"net/http"

	"github.com/sirupsen/logrus"
)

var userService service.UserService

func UserRegister(rw http.ResponseWriter, req *http.Request) {
	var (
		u   model.User
		err error
	)

	util.Bind(req, &u)
	u, err = userService.UserRegister(u.Mobile, u.Passwd, u.Nickname, u.Avatar, u.Sex)
	if err != nil {
		util.RespFail(rw, err.Error())
		return
	}

	util.RespOk(rw, u, "")
}

func UserLogin(rw http.ResponseWriter, req *http.Request) {
	var (
		mobile   string
		plainPwd string
		u        model.User
		err      error
	)
	req.ParseForm()
	mobile = req.PostForm.Get("mobile")
	plainPwd = req.PostForm.Get("passwd")
	logrus.Println("mobile: ", mobile)
	logrus.Println("passwd: ", plainPwd)

	//校验参数
	if len(mobile) == 0 || len(plainPwd) == 0 {
		util.RespFail(rw, "用户名或者密码不正确")
		return
	}

	u, err = userService.Login(mobile, plainPwd)
	if err != nil {
		util.RespFail(rw, err.Error())
		return
	}

	logrus.Println("login user", u)
	util.RespOk(rw, u, "")
}
