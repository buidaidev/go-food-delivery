package subscriber

import (
	"context"
	component "go-food-delivery/component/appctx"
)

func Setup(appCtx component.AppContext, ctx context.Context) {
	IncreaseLikeCountAfterUserUnlikeRestaurant(appCtx, ctx)
	DecreaseLikeCountAfterUserUnlikeRestaurant(appCtx, ctx)
}
