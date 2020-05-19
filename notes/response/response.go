package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, statusCode int, code int, msg string, data interface{}){
	c.JSON(statusCode, gin.H{
		"code" : code,
		"message":msg,
		"data" : data,
	})
}

// 成功返回
func Success(c *gin.Context, data gin.H, msg string){
	Response(c, http.StatusOK, 200, msg, data)
}

// 失败返回
func Faild(c *gin.Context, data gin.H, msg string){
	Response(c, http.StatusOK, 400, msg, data)
}