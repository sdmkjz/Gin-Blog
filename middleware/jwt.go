package middleware

import (
	"gin-blog/pkg/e"
	"gin-blog/pkg/utils"
	"gin-blog/serializer"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCSE
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.ERROR_TOKEN_EXIST
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = e.ERROR_TOKEN_TYPE_WRONG
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = e.ERROR_TOKEN_RUNTIME
			}
		}
		if code != e.SUCCSE {
			c.JSON(e.SUCCSE, serializer.Response{
				Status: code,
				Msg:    e.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
