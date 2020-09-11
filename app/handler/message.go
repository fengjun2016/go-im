package handler

import (
	"go-im/app/args"
	"go-im/app/util"
	"net/http"
)

//加载私聊消息列表
func LoadPersonalMessage(writer http.ResponseWriter, request *http.Request) {
	var arg args.MessageArg
	util.Bind(request, &arg)
	messages, err := messageService.SearchPersonalMessage(CmdSingleMsg, arg.Userid, arg.Dstid)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	util.RespOkList(writer, messages, len(messages))
}

//加载群聊消息列表
func LoadCommunityMessage(writer http.ResponseWriter, request *http.Request) {
	var arg args.MessageArg
	util.Bind(request, &arg)
	messages, err := messageService.SearchCommunityMessage(CmdRoomMsg, arg.Dstid)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	util.RespOkList(writer, messages, len(messages))
}
