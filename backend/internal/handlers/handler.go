package handlers

import "github.com/gin-gonic/gin"

const StatusSuccess = "success"
const StatusError = "error"

func NewResponseJsonStruct(status, message string, error error, data any) gin.H {
	resp := gin.H{
		"message": message,
	}
	if error != nil {
		resp["status"] = StatusError
		resp["error"] = error.Error()
	} else {
		resp["status"] = StatusSuccess
		resp["data"] = data
	}
	return resp
}
