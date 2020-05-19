package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"notes/response"
)

func RecoverMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		defer func(){
			if err := recover(); err != nil{
				response.Faild(c,nil,fmt.Sprint(err))
			}
		}()
	}
}
