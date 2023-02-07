package ginrestaurant

import (
	"go-food-delivery/common"
	component "go-food-delivery/component/appctx"
	restaurantbusiness "go-food-delivery/module/restaurant/business"
	restaurantmodel "go-food-delivery/module/restaurant/model"
	restaurantstore "go-food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.RestaurantUpdate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstore.NewSQLStore(db)
		business := restaurantbusiness.NewUpdateRestaurantBusiness(store, requester)

		if err := business.UpdateRestaurant(ctx.Request.Context(), &data, int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}
