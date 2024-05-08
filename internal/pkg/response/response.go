package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}



func Success (c *gin.Context){
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Msg: "ok",
		Data: nil,
	})
}


func SuccessMessage(c *gin.Context, message string){
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Msg: message,
		Data: nil,
	})
}

func SuccessData()

func FailWithStatus(c *gin.Context, status int, err error) {
	c.JSON(status, Response{
		Code:    status,
		Msg: err.Error(),
		Data: nil,
	})
}

