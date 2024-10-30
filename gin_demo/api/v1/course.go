package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Course struct {
	Name    string  `form:"name" binding:"required"`
	Teacher string  `form:"teacher" binding:"required"`
	Price   float64 `form:"price" binding:"number"`
}

func GetCourse(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	c.JSON(http.StatusOK, gin.H{
		"message": "get course: " + id,
	})
}

func AddCourse(c *gin.Context) {
	req := Course{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, req)
}
