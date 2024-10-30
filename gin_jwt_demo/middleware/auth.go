package middleware

import (
	"github.com/gin-gonic/gin"
	proj_jwt "github.com/showyquasar88/proj-combine/gin_jwt_demo/jwt"
	"net/http"
)

var token = "1234567"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("access_token")
		if accessToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "无权限",
			})
			c.Abort()
			return
		}
		data := &proj_jwt.Data{}
		err := proj_jwt.Verify(accessToken, data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("auth_info", data)
		c.Next()
	}
}
