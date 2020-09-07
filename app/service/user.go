package service

import (
	"errors"
	"fastIM/app/util"
	"fmt"
	"go-im/app/model"
	"math/rand"
	"time"
)

type UserService struct{}

//用户注册
func (us *UserService) UserRegister(mobile, plainPwd, nickname, avatar, sex string) (u model.User, err error) {
	var (
		exists bool
	)

	u.Mobile = mobile

	exists, err = u.ExistUserByMobile()
	if err != nil {
		return
	}

	//代表已有相同手机号已注册
	if exists {
		err = errors.New("该手机号已注册")
		return
	}

	u.Mobile = mobile
	u.Avatar = avatar
	u.Nickname = nickname
	u.Sex = sex
	u.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	u.Passwd = util.MakePasswd(plainPwd, u.Salt)
	u.Createat = time.Now()

	err = u.Create()

	return
}

//用户登录
func (us *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	//数据库操作
	var (
		exists bool
	)

	u.Mobile = mobile

	exists, err = u.ExistUserByMobile()
	if err != nil {
		return
	}

	if !exists {
		err = errors.New("用户不存在")
		return
	}

	//判断密码是否正确
	if !util.ValidatePasswd(plainpwd, loginUser.Salt, loginUser.Passwd) {
		return loginUser, errors.New("密码不正确")
	}
	//刷新用户登录的token值
	token := util.GenRandomStr(32)
	loginUser.Token = token
	model.DbEngine.ID(loginUser.Id).Cols("token").Update(&loginUser)

	//返回新用户信息
	return loginUser, nil
}

//查找某个用户
func (s *UserService) Find(userId int64) (user model.User, err error) {
	user.Id = userId
	err = user.Get()
	return
}
