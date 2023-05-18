package model

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		log.Printf("tokenStringr: %v\n", tokenString)
		if len(ctx.Request.Header) > 0 {
			for k, v := range ctx.Request.Header {
				fmt.Printf("%s=%s\n", k, v[0])
			}
		}
		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusOK, gin.H{"code": 401, "msg": "权限不足,token为空"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusOK, gin.H{"code": 401, "msg": "权限不足,token解析失败"})
			ctx.Abort()
			return

		}
		//验证通过后获取claim 中的userid
		userId := claims.UserId

		var user User

		db.Where("id = ?", userId).First(&user)
		if user.ID == 0 {
			ctx.JSON(http.StatusOK, gin.H{"code": 401, "msg": "权限不足,token验证失败"})
			ctx.Abort()
			return
		}
		//用户存在 将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
func XcxAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		log.Printf("tokenStringr: %v\n", tokenString)
		if len(ctx.Request.Header) > 0 {
			for k, v := range ctx.Request.Header {
				fmt.Printf("%s=%s\n", k, v[0])
			}
		}
		log.Printf("tokenStringr: 1 \n")

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusOK, gin.H{"code": 401, "msg": "权限不足,token为空"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		log.Printf("tokenStringr: 2 \n")

		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusOK, gin.H{"code": 401, "msg": "权限不足,token解析失败"})
			ctx.Abort()
			return

		}
		log.Printf("tokenStringr: 3 \n")

		//验证通过后获取claim 中的userid
		userId := claims.UserId

		var user XcxUser

		db.Where("id = ?", userId).First(&user)
		if user.ID == 0 {
			ctx.JSON(http.StatusOK, gin.H{"code": 401, "msg": "权限不足,token验证失败"})
			ctx.Abort()
			return
		}
		//用户存在 将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
