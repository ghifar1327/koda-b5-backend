package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.POST("/", func(c *gin.Context) {
		var body Auth
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Internal Server Error",
		})
		return
	}
	isValid :=  isEmailValid(body.Email)
	if isValid {
		c.JSON(http.StatusOK, gin.H{
			"msg": "OK",
			"auth": body,
		})
		}
	})
	app.Run("localhouse:8080")
}

type Auth struct {
	Email string 	`form:"email"`
	Passwoed string `form:"password"`
}


func isEmailValid(e string) bool {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return emailRegex.MatchString(e)
}