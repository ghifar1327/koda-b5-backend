package router

import (
	"github.com/ghifar1327/koda-b5-backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	authControler := controller.NewAuthConroller()
	app.POST("/auth/register", authControler.Register)
	app.POST("/auth/login", authControler.Login)
}
