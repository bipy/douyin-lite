package routes

import (
	"douyin-lite/app/controllers"
	"github.com/labstack/echo/v4"
)

func PublicRoutes(a *echo.Echo) {
	douyin := a.Group("/douyin")

	douyin.GET("/feed", controllers.GetFeed)

	userGroup := douyin.Group("/user")
	userGroup.GET("/", controllers.GetUserInfo)
	userGroup.POST("/register/", controllers.Register)
	userGroup.POST("/login/", controllers.Login)

	publishGroup := douyin.Group("/publish")
	publishGroup.POST("/action/", controllers.PublishAction)
	publishGroup.GET("/list/", controllers.GetPublishList)

	favoriteGroup := douyin.Group("/favorite")
	favoriteGroup.POST("/action/", controllers.FavoriteAction)
	favoriteGroup.GET("/list/", controllers.GetFavoriteList)

	commentGroup := douyin.Group("/comment")
	commentGroup.POST("/action/", controllers.CommentAction)
	commentGroup.GET("/list/", controllers.GetCommentList)

	relationGroup := douyin.Group("/relation")
	relationGroup.POST("/action/", controllers.RelationAction)

	followGroup := relationGroup.Group("/follow")
	followGroup.GET("/list/", controllers.GetFollowList)

	followerGroup := relationGroup.Group("/follower")
	followerGroup.GET("/list/", controllers.GetFollowerList)

	fileGroup := a.Group("/file")
	fileGroup.GET("/:type/:uuid", controllers.GetFile)
}
