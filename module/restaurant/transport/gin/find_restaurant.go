package ginrestaurant

import (
	"go-food-delivery/common"
	component "go-food-delivery/component/appctx"
	restaurantbusiness "go-food-delivery/module/restaurant/business"
	restaurantstore "go-food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstore.NewSQLStore(db)
		business := restaurantbusiness.NewFindRestaurantBusiness(store)
		data, err := business.FindRestaurant(ctx.Request.Context(), map[string]interface{}{"id": int(uid.GetLocalID())})

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(data))
	}
}
