package main

import (
	component "go-food-delivery/component/appctx"

	"go-food-delivery/middleware"
	ginuser "go-food-delivery/module/user/transport/gin"

	"github.com/gin-gonic/gin"
)

func setupAdminRoutes(appCtx component.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin", middleware.RequiredAuth(appCtx), middleware.RoleRequired(appCtx, "admin", "mod"))
	{
		admin.GET("/profile", ginuser.Profile(appCtx))
	}
}
