package ginuser

import (
	"go-food-delivery/common"
	component "go-food-delivery/component/appctx"
	"go-food-delivery/component/hasher"
	"go-food-delivery/component/tokenprovider/jwt"
	userbusiness "go-food-delivery/module/user/business"
	usermodel "go-food-delivery/module/user/model"
	userstorage "go-food-delivery/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := ctx.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMaiDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		business := userbusiness.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(ctx.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(account))
	}
}
