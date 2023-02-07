package ginuser

import (
	"go-food-delivery/common"
	component "go-food-delivery/component/appctx"
	"go-food-delivery/component/hasher"
	usersbusiness "go-food-delivery/module/user/business"
	usermodel "go-food-delivery/module/user/model"
	userstorage "go-food-delivery/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		var data usermodel.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		business := usersbusiness.NewRegisterBusiness(store, md5)

		if err := business.Register(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(data.FakeId.String()))
	}
}
