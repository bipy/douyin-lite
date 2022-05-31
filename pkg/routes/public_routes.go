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
	publishGroup.POST("/action", controllers.PublishAction)
	publishGroup.GET("/list", controllers.GetPublishList)

	favoriteGroup := a.Group("/favorite")
	favoriteGroup.POST("/action", controllers.FavoriteAction)
	favoriteGroup.GET("/list", controllers.GetFavoriteList)

	commentGroup := a.Group("/comment")
	commentGroup.POST("/action", controllers.CommentAction)
	commentGroup.GET("/list", controllers.GetCommentList)

	relationGroup := a.Group("/relation")
	relationGroup.POST("/action", controllers.RelationAction)

	followGroup := relationGroup.Group("/follow")
	followGroup.GET("/list", controllers.GetFollowList)

	followerGroup := relationGroup.Group("/follower")
	followerGroup.GET("/list", controllers.GetFollowerList)
}
