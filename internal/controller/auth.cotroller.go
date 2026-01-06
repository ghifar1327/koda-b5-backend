package controller

import (
	"net/http"
	"github.com/ghifar1327/koda-b5-backend/internal/dto"
	"github.com/ghifar1327/koda-b5-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}
var Users = make(map[string]string)

func NewAuthConroller() *AuthController {
	return &AuthController{}
}
func (a AuthController) Register(c *gin.Context) {
	var body *dto.Auth
	if err := c.ShouldBindWith(&body, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid Request body",
		})
		return
	}
	if !service.IsEmailValid(body.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid email format",
		})
		return
	}
	if len(body.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "password must be at latest 6 characters",
		})
		return
	}
	if _, exist := Users[body.Email]; exist {
		c.JSON(http.StatusConflict, gin.H{
			"msg": "Email Already Registered",
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(body.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to hash password",
		})
		return
	}
	Users[body.Email] = string(hashedPassword)
	c.JSON(http.StatusCreated, gin.H{
		"msg":  "register success",
		"email": body.Email,
	})
}

func (a AuthController)Login(c *gin.Context){
 	var body *dto.Auth
	
	if err := c.ShouldBindWith(&body, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid requst body",
		})
	}
	if !service.IsEmailValid(body.Email){
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid email format",
		})
	}
	hashedPassword, exists := Users[body.Email]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "email or password wrong",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(body.Password),
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "email or password wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "selamat datang " + body.Email ,
	})
}