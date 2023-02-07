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

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		var data restaurantmodel.RestaurantCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data.UserId = requester.GetUserId()

		store := restaurantstore.NewSQLStore(db)
		business := restaurantbusiness.NewCreateRestaurantBusiness(store)

		if err := business.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(data.FakeId))
	}
}
