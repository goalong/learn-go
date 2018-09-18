package jwt

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/goalong/learn-go/gin_blog/pkg/util"
	"github.com/goalong/learn-go/gin_blog/pkg/err"
)


func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = err.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = err.INVALID_PARAMS
		} else {
			claims, _err := util.ParseToken(token)
			if _err != nil {
				code = err.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = err.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != err.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg": err.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
