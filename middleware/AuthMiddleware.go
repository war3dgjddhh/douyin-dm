package middleware

import (
	"net/http"
	"strings"

	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		//	tokenString := c.Request.Header.Get("Authorization")
		tokenString := c.Query("token")
		if tokenString == "" {
			tokenString = c.PostForm("token")
		}
		// token为空 非法token
		if tokenString == "" || len(tokenString) < 7 || !strings.HasPrefix(tokenString, common.ToeknPrefix) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		// 提取token的有效部分
		logrus.Infof("go run to this , token =%s", tokenString)
		tokenString = tokenString[len(common.ToeknPrefix):]
		// 解析token
		token, claims, err := common.ParseToken(tokenString)
		logrus.Infof("go run resolve token , token =%s", tokenString)

		// 非法token
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		// 获取claims中的userId
		userId := claims.UserId
		user := repository.GetUserById(userId)
		// 将用户信息写入上下文便于读取
		c.Set("user", user)
		c.Next()
	}
}
