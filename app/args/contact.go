package args

type ContactArg struct {
	PageArg
	Userid string `json:"userid" form:"userid"`
	Dstid  string `json:"dstid" form:"dstid"`
}

//添加新的成员
type AddNewMember struct {
	PageArg
	Userid  string `json:"userid" form:"userid"`
	DstName string `json:"dstname" form:"dstname"`
}
