package middleware

import (
	"go-food-delivery/common"
	"go-food-delivery/component/appctx"

	"github.com/gin-gonic/gin"
)

func Recover(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				if appError, ok := err.(*common.AppError); ok {
					ctx.AbortWithStatusJSON(appError.StatusCode, appError)
					return
				}

				appError := common.ErrInternal(err.(error))
				ctx.AbortWithStatusJSON(appError.StatusCode, appError)
				return
			}
		}()

		ctx.Next()
	}
}
