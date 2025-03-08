package router

import "github.com/gin-gonic/gin"

func Initialize() {
	route := gin.Default()

	route.Run()
}
