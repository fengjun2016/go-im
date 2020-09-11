package args

type MessageArg struct {
	PageArg
	Userid string `json:"userid" form:"userid"`
	Dstid  string `json:"dstid" form:"dstid"`
}
