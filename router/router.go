package router

import (
	"github.com/GustavoAFur/gogogo/router"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	route := gin.Default()

	router.Initialize(route)

	route.Run()
}
