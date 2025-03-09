package router

import "github.com/gin-gonic/gin"

func initializeRoutes(route *gin.Engine) {
	v1 := route.Group("/api/v1")
	{
		v1.GET("/opening")
	}
}
