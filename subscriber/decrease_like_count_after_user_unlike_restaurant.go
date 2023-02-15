package subscriber

import (
	"context"
	"go-food-delivery/common"
	component "go-food-delivery/component/appctx"
	restaurantstorage "go-food-delivery/module/restaurant/storage"
)

func DecreaseLikeCountAfterUserUnlikeRestaurant(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)

	store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}
