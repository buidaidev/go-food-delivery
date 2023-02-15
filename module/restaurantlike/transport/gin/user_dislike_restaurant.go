package ginrestaurantlike

import (
	"go-food-delivery/common"
	component "go-food-delivery/component/appctx"
	restaurantlikebusiness "go-food-delivery/module/restaurantlike/business"
	restaurantlikestorage "go-food-delivery/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserDislikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		ps := appCtx.GetPubSub()
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(db)		
		business := restaurantlikebusiness.NewUserDislikeRestaurantBusiness(store, ps)

		if err := business.DislikeRestaurant(ctx.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}
