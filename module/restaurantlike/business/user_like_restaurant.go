package restaurantlikebusiness

import (
	"context"
	"go-food-delivery/common"
	restaurantlikemodel "go-food-delivery/module/restaurantlike/model"
	"go-food-delivery/pubsub"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(context context.Context, data *restaurantlikemodel.Like) error
}

type userLikeRestaurantBusiness struct {
	store UserLikeRestaurantStore
	ps    pubsub.Pubsub
}

func NewUserLikeRestaurantBusiness(store UserLikeRestaurantStore, ps pubsub.Pubsub) *userLikeRestaurantBusiness {
	return &userLikeRestaurantBusiness{store: store, ps: ps}
}

func (business *userLikeRestaurantBusiness) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := business.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCanNotLikeRestaurant(err)
	}

	if err := business.ps.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	return nil
}
