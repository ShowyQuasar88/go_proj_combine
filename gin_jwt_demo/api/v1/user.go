package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name  string `form:"name" binding:"required"`
	Phone string `form:"phone" binding:"required,e164"`
}

func GetUser(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	c.JSON(http.StatusOK, gin.H{
		"message": "get user: " + id,
	})
}

func GetUserProto(c *gin.Context) {
	c.ProtoBuf(http.StatusOK, gin.H{
		"message": "get user: ",
	})
}

func AddUser(c *gin.Context) {
	req := User{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, req)
}
