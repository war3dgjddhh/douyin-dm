package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	r.Use(gin.Logger())

	apiRouter := r.Group("/douyin")

	authRouter := apiRouter.Group("").Use(middleware.AuthMiddleware())

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	authRouter.GET("/user/", controller.UserInfo) // auth
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	authRouter.POST("/publish/action/", controller.Publish)  // auth
	authRouter.GET("/publish/list/", controller.PublishList) // auth

	// extra apis - I
	authRouter.POST("/favorite/action/", controller.FavoriteAction) // auth
	authRouter.GET("/favorite/list/", controller.FavoriteList)      // auth
	authRouter.POST("/comment/action/", controller.CommentAction)   // auth
	authRouter.GET("/comment/list/", controller.CommentList)        // auth

	// extra apis - II
	authRouter.POST("/relation/action/", controller.RelationAction)     // auth
	authRouter.GET("/relation/follow/list/", controller.FollowList)     // auth
	authRouter.GET("/relation/follower/list/", controller.FollowerList) // auth
	authRouter.GET("/relation/friend/list/", controller.FriendList)     // auth
	authRouter.GET("/message/chat/", controller.MessageChat)            // auth looks difficulty
	authRouter.POST("/message/action/", controller.MessageAction)       // auth
}
