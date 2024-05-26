package controller

import (
	"encoding/json"
	"net/http"
	"oauth/model"
	"oauth/respMsg"
	"oauth/sys"
)

func ParseRequestBody(w http.ResponseWriter, r *http.Request, input any) {
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func HandleSuccess(w http.ResponseWriter, data any) {
	var resp model.RESP
	resp.MsgCode = respMsg.SUCCESS.MsgCode
	resp.MsgLevel = respMsg.SUCCESS.MsgLevel
	resp.Desc = respMsg.SUCCESS.Desc
	resp.Data = data

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		sys.Logger().Error(err)
	}
}

func HandleError(w http.ResponseWriter, msgType respMsg.ResponseType, data any) {
	var resp model.RESP

	resp.MsgCode = msgType.MsgCode
	resp.MsgLevel = msgType.MsgLevel
	resp.Desc = msgType.Desc
	resp.Data = data

	var statusCode int
	switch msgType.MsgLevel {
	case "SUCCESS":
		statusCode = http.StatusOK
		break
	case "ERROR":
		statusCode = http.StatusInternalServerError
		break
	case "WARNING":
		statusCode = http.StatusBadRequest
		break
	case "INFO":
		statusCode = http.StatusOK
		break
	case "UNAUTHORIZED":
		statusCode = http.StatusUnauthorized
		break
	default:
		panic("[HandleError] unknown msg level")
	}

	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		sys.Logger().Error(err)
	}
}
