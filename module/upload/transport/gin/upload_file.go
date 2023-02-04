package ginupload

import (
	"go-food-delivery/common"
	"go-food-delivery/component/appctx"
	uploadbusiness "go-food-delivery/module/upload/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := ctx.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		business := uploadbusiness.NewUploadBusiness(appCtx.UploadProvider(), nil)
		img, err := business.Upload(ctx.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRespone(img))
	}
}
