package routes

import (
	"douyin-lite/app/controllers"
	"github.com/labstack/echo/v4"
)

func PublicRoutes(a *echo.Echo) {
	feedGroup := a.Group("/feed")
	feedGroup.GET("/", controllers.GetFeed)

	userGroup := a.Group("/user")
	userGroup.GET("/", controllers.GetUserInfo)
	userGroup.POST("/register", controllers.Register)
	userGroup.POST("/login", controllers.Login)

	publishGroup := a.Group("/publish")
	publishGroup.POST("/action", controllers.GetHello)
	publishGroup.GET("/list", controllers.GetHello)

	favoriteGroup := a.Group("/favorite")
	favoriteGroup.POST("/action", controllers.GetHello)
	favoriteGroup.GET("/list", controllers.GetHello)

	commentGroup := a.Group("/comment")
	commentGroup.POST("/action", controllers.GetHello)
	commentGroup.GET("/list", controllers.GetHello)

	relationGroup := a.Group("/relation")
	relationGroup.POST("/action", controllers.GetHello)

	followGroup := relationGroup.Group("/follow")
	followGroup.GET("/list", controllers.GetHello)

	followerGroup := relationGroup.Group("/follower")
	followerGroup.GET("/list", controllers.GetHello)
}
