package main

import (
	component "go-food-delivery/component/appctx"

	"go-food-delivery/middleware"

	ginrestaurant "go-food-delivery/module/restaurant/transport/gin"
	ginupload "go-food-delivery/module/upload/transport/gin"
	ginuser "go-food-delivery/module/user/transport/gin"

	"github.com/gin-gonic/gin"
)

func setupRoutes(appCtx component.AppContext, v1 *gin.RouterGroup) {
	// upload
	v1.POST("/upload", ginupload.Upload(appCtx))

	// auth
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/authenticate", ginuser.Login(appCtx))
	v1.POST("/profile", middleware.RequiredAuth(appCtx), ginuser.Profile(appCtx))

	// restaurants
	restaurants := v1.Group("restaurants", middleware.RequiredAuth(appCtx))
	{
		restaurants.POST("/", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.FindRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

}
