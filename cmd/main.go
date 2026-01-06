package main

import (
	"github.com/ghifar1327/koda-b5-backend/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	router.Init(app)
	app.Run(":3000")
}