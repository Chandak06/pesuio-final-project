package auth

import (
	"github.com/anuragrao04/pesuio-final-project/database"
	"github.com/anuragrao04/pesuio-final-project/models"
	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {
	var request models.SignInRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid input",
		})
	}

	// implement
	value, data := database.CheckPassword(request.Username, request.Password)

	c.JSON(200, gin.H{
		"success": true,
		"message": value,
		"record":  data,
	})
}
