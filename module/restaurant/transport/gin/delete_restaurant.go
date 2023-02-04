package ginrestaurant

import (
	"go-food-delivery/common"
	"go-food-delivery/component/appctx"
	restaurantbusiness "go-food-delivery/module/restaurant/business"
	restaurantstore "go-food-delivery/module/restaurant/storage"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstore.NewSQLStore(db)
		business := restaurantbusiness.NewDeleteRestaurantBusiness(store)

		if err := business.DeleteRestaurant(ctx.Request.Context(), id); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}
