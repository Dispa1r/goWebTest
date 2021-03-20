package middleware

import (
	"github.com/gin-gonic/gin"
	"goWeb/common"
	"goWeb/model"
	"log"
	"net/http"
	"strings"
)
func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenString :=c.GetHeader("Authorization")
		log.Println(tokenString)
		if tokenString=="" || !strings.HasPrefix(tokenString,"Bearer "){//Bearer来源于jwt规定的文档
			log.Println("授权失败1")
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"未授权",
			})
			c.Abort()
			return

		}
		tokenString =tokenString[7:]
		token,claims,err :=common.ParseToken(tokenString)
		if err!=nil ||!token.Valid{
			log.Println("授权失败2")
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"未授权",
			})
			c.Abort()
			return
		}
		userId :=claims.UserId
		DB :=common.GetDB()
		var user model.User
		DB.First(&user,userId)
		if userId ==0{
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"用户不存在 ",
			})
			c.Abort()
			return
		}
		c.Set("user",user)

	}
}