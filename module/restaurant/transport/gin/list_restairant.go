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

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		var pagingData common.Paging

		if err := ctx.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		pagingData.Fullfill()

		var filter restaurantmodel.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter.Status = []int{1}

		store := restaurantstore.NewSQLStore(db)
		business := restaurantbusiness.NewListRestaurantBusiness(store)
		result, err := business.ListRestaurant(ctx.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessRespone(result, pagingData, filter))
	}
}
