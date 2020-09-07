package handler

import (
	"go-im/app/model"
	"go-im/app/util"
	"net/http"
)

func UserRegister(rw http.ResponseWriter, req *http.Request) {
	var (
		u   model.User
		err error
	)

	util.Bind(request, &u)
	u, err = UserService.UserRegister(u.Mobile, u.Passwd, u.Nickname, u.Avatar, u.Sex)
	if err != nil {
		util.RespFail(rw, err.Error())
	} else {
		util.RespOk(rw, u, "")
	}
}

func UserLogin(rw http.ResponseWriter, req *http.Request) {
	var (
		mobile   string
		plainPwd string
		u        model.User
		err      error
	)

	mobile = request.PostForm.Get("mobile")
	plainPwd = request.PostForm.Get("passwd")

	//校验参数
	if len(mobile) == 0 || len(plainPwd) == 0 {
		util.RespFail(rw, "用户名或者密码不正确")
	}

	u, err = UserService.Login(mobile, plainPwd)
	if err != nil {
		util.RespFail(rw, err.Error())
	} else {
		util.RespOk(rw, u, "")
	}
}
