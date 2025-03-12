package routes

import (
	"garavel/internal/controllers"
	"garavel/internal/middlewares"

	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func HandleRequest() {
	R = gin.Default()
	R.Use(gin.Recovery())
	R.GET("/health", controllers.Health)
	R.POST("/login", controllers.Login)
	crud(controllers.UserClass, true)
	crud(controllers.CustomerClass, true)
}

func crud(collect string, isProtected bool) {
	crl := controllers.Make(collect)
	group := R.Group("/" + collect)
	if isProtected {
		group.Use(middlewares.Auth())
	}
	group.GET("", crl.Index)
	group.POST("", crl.Create)
	group.GET("/:id", crl.Show)
	group.PUT("/:id", crl.Update)
	group.DELETE("/:id", crl.Destroy)
}
