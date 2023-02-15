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

func UserLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		ps := appCtx.GetPubSub()
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(db)
		business := restaurantlikebusiness.NewUserLikeRestaurantBusiness(store, ps)

		if err := business.LikeRestaurant(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}
