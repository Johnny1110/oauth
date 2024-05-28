package respMsg

type ResponseType struct {
	MsgCode  string `json:"msgCode"`
	MsgLevel string `json:"msgLevel"`
	Desc     string `json:"desc"`
}

var (
	SUCCESS      = ResponseType{MsgCode: "0000000", MsgLevel: "SUCCESS", Desc: "access success."}
	INFO         = ResponseType{MsgCode: "0000000", MsgLevel: "INFO", Desc: "notice info."}
	UNAUTHORIZED = ResponseType{MsgCode: "9999992", MsgLevel: "UNAUTHORIZED", Desc: "access denied."}
	WARNING      = ResponseType{MsgCode: "0000001", MsgLevel: "WARNING", Desc: "warning msg."}
	ERROR        = ResponseType{MsgCode: "9999999", MsgLevel: "ERROR", Desc: "system error."}

	INCORRECT_INPUT    = ResponseType{MsgCode: "9000001", MsgLevel: "WARNING", Desc: "input incorrect."}
	OPERATION_PROHIBIT = ResponseType{MsgCode: "9000002", MsgLevel: "WARNING", Desc: "not allow to do this."}
	RESOURCE_NOT_FOUND = ResponseType{MsgCode: "9000003", MsgLevel: "WARNING", Desc: "resource not found."}
	SYSTEM_ERROR       = ResponseType{MsgCode: "9999998", MsgLevel: "WARNING", Desc: "system error."}

	PASSWORD_INCORRECT    = ResponseType{MsgCode: "8000001", MsgLevel: "WARNING", Desc: "password incorrect."}
	ACCOUNT_ALREADY_EXIST = ResponseType{MsgCode: "8000002", MsgLevel: "WARNING", Desc: "account already exist."}
	ACCOUNT_NOT_EXIST     = ResponseType{MsgCode: "8000003", MsgLevel: "WARNING", Desc: "account not exist."}
)
