package ginrestaurant

import (
	"go-food-delivery/common"
	component "go-food-delivery/component/appctx"
	restaurantbusiness "go-food-delivery/module/restaurant/business"
	restaurantstore "go-food-delivery/module/restaurant/storage"

	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstore.NewSQLStore(db)
		business := restaurantbusiness.NewDeleteRestaurantBusiness(store, requester)

		if err := business.DeleteRestaurant(ctx.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}
