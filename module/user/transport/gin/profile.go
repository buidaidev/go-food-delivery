package ginuser

import (
	"go-food-delivery/common"
	component "go-food-delivery/component/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := ctx.MustGet(common.CurrentUser).(common.Requester)

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(u))
	}
}
