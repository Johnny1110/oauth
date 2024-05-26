package model

type RESP struct {
	MsgCode  string `json:"msgCode"`
	MsgLevel string `json:"msgLevel"`
	Desc     string `json:"desc"`
	Data     any    `json:"data"`
}
