package args

type MessageArg struct {
	PageArg
	Userid string `json:"userid" form:"userid"`
}
