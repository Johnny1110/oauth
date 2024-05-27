package respMsg

type ResponseType struct {
	MsgCode  string `json:"msgCode"`
	MsgLevel string `json:"msgLevel"`
	Desc     string `json:"desc"`
}

var (
	SUCCESS      = ResponseType{MsgCode: "0000000", MsgLevel: "SUCCESS", Desc: "access success."}
	INFO         = ResponseType{MsgCode: "0000001", MsgLevel: "INFO", Desc: "notice info."}
	UNAUTHORIZED = ResponseType{MsgCode: "9000001", MsgLevel: "UNAUTHORIZED", Desc: "access denied."}
	WARNING      = ResponseType{MsgCode: "9000002", MsgLevel: "WARNING", Desc: "warning msg."}
	ERROR        = ResponseType{MsgCode: "9999999", MsgLevel: "ERROR", Desc: "system error."}

	ACCOUNT_ALREADY_EXIST = ResponseType{MsgCode: "9000003", MsgLevel: "ERROR", Desc: "account already exist."}
)
