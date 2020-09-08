package handler

import (
	"go-im/app/args"
	"go-im/app/model"
	"go-im/app/service"
	"go-im/app/util"
	"log"
	"net/http"
)

var concatService service.ContactService

//添加朋友
func AddFriend(writer http.ResponseWriter, request *http.Request) {
	var arg args.AddNewMember
	util.Bind(request, &arg)
	friend, err := concatService.SearchFriendByName(arg.DstName)
	if friend.Id == "" || err != nil {
		util.RespFail(writer, "您要添加的好友不存在")
	} else {
		//调用service
		err := concatService.AddFriend(arg.Userid, friend.Id)
		if err != nil {
			util.RespFail(writer, err.Error())
		} else {
			util.RespOk(writer, nil, "好友添加成功!")
		}
	}
}

//加载好友列表
func LoadFriend(writer http.ResponseWriter, request *http.Request) {
	var arg args.ContactArg
	util.Bind(request, &arg)
	users, err := concatService.SearchFriend(arg.Userid)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	util.RespOkList(writer, users, len(users))
}

//创建群
func CreateCommunity(w http.ResponseWriter, req *http.Request) {
	var arg model.Community
	//如果这个用的上,那么可以直接
	util.Bind(req, &arg)
	com, err := concatService.CreateCommunity(arg)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, com, "")
	}
}

//用户加群
func JoinCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.AddNewMember
	//如果这个用的上,那么可以直接
	util.Bind(req, &arg)
	//查看群是否存在
	com, err := concatService.SearchCommunityByName(arg.DstName)
	if com.Id == "" || err != nil {
		util.RespFail(w, "您要加入的群不存在")
	} else {
		log.Printf("community id:%d", com.Id)
		err := concatService.JoinCommunity(arg.Userid, com.Id)
		//刷新用户的群组信息 todo
		// AddGroupId(arg.Userid, com.Id)
		if err != nil {
			util.RespFail(w, err.Error())
		} else {
			util.RespOk(w, nil, "")
		}
	}
}

//加载群列表
func LoadCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(req, &arg)
	comunitys, err := concatService.SearchComunity(arg.Userid)
	if err != nil {
		util.RespFail(w, err.Error())
		return
	}

	util.RespOkList(w, comunitys, len(comunitys))
}
