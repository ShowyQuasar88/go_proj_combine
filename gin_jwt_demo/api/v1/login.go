package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	proj_jwt "github.com/showyquasar88/proj-combine/gin_jwt_demo/jwt"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	data := proj_jwt.Data{
		Name:   "showyquasar88",
		Age:    18,
		Gender: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                    // 不可处理时间
		},
	}
	sign, err := proj_jwt.Sign(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"access_token": sign,
	})
}
