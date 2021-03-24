package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONResponse json response
type JSONResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	// CodeSuccess default success code
	CodeSuccess = 0
	// CodeError default error code
	CodeError = -1
	// MsgSuccess default success message
	MsgSuccess = "success"
	// MsgError default error message
	MsgError = "error"
)

// Result base json json response
func Result(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, JSONResponse{
		code,
		msg,
		data,
	})
}

// Ok base ok json response
func Ok(c *gin.Context) {
	Result(c, CodeSuccess, map[string]interface{}{}, MsgSuccess)
}

// OkMessage ok json response with customize message
func OkMessage(c *gin.Context, message string) {
	Result(c, CodeSuccess, map[string]interface{}{}, message)
}

// OkData ok json response with data
func OkData(c *gin.Context, data interface{}) {
	Result(c, CodeSuccess, data, MsgSuccess)
}

// OkDetailes ok json response with all details
func OkDetailes(c *gin.Context, data interface{}, message string) {
	Result(c, CodeSuccess, data, message)
}

// Fail failed json response
func Fail(c *gin.Context) {
	Result(c, CodeError, map[string]interface{}{}, MsgError)
}

// FailMessage failed json response with customized message
func FailMessage(c *gin.Context, message string) {
	Result(c, CodeError, map[string]interface{}{}, message)
}

// FailCodeMessage failed json response with customized code and message
func FailCodeMessage(c *gin.Context, code int, message string) {
	Result(c, code, map[string]interface{}{}, message)
}

// FailDetails failed json response with all details
func FailDetails(c *gin.Context, code int, data interface{}, message string) {
	Result(c, code, data, message)
}
