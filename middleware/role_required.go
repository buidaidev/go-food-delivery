package middleware

import (
	"errors"
	"go-food-delivery/common"
	component "go-food-delivery/component/appctx"

	"github.com/gin-gonic/gin"
)

func RoleRequired(appCtx component.AppContext, allowRoles ...string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		u := ctx.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false

		for _, item := range allowRoles {
			if u.GetRole() == item {
				hasFound = true
				break
			}
		}

		if !hasFound {
			panic(common.ErrNoPermission(errors.New("invalid role user")))
		}

		ctx.Next()
	}
}
