package ginrestaurantlike

import (
	"go-food-delivery/common"
	component "go-food-delivery/component/appctx"
	restaurantlikebusiness "go-food-delivery/module/restaurantlike/business"
	restaurantlikemodel "go-food-delivery/module/restaurantlike/model"
	restaurantlikestorage "go-food-delivery/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUserLikes(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fullfill()

		store := restaurantlikestorage.NewSQLStore(db)
		business := restaurantlikebusiness.NewListUserLikeRestaurantBusiness(store)

		result, err := business.ListUsersLikeRestaurant(ctx.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}
