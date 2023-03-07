package model

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 400, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 400, "msg": "权限不足"})
			ctx.Abort()
			return

		}
		//验证通过后获取claim 中的userid
		userId := claims.UserId

		var user User

		db.Where("id = ?", userId).First(&user)
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 400, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//用户存在 将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
