package restaurantlikebusiness

import (
	"context"
	"go-food-delivery/common"
	restaurantlikemodel "go-food-delivery/module/restaurantlike/model"
	"go-food-delivery/pubsub"
	"log"
)

type UserDislikeRestaurantStore interface {
	Delete(context context.Context, userId, restaurantId int) error
}

type userDislikeRestaurantBusiness struct {
	store UserDislikeRestaurantStore
	ps    pubsub.Pubsub
}

func NewUserDislikeRestaurantBusiness(store UserDislikeRestaurantStore, ps pubsub.Pubsub) *userDislikeRestaurantBusiness {
	return &userDislikeRestaurantBusiness{store: store, ps: ps}
}

func (business *userDislikeRestaurantBusiness) DislikeRestaurant(
	ctx context.Context,
	userId, restaurantId int,
) error {
	err := business.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCanNotDislikeRestaurant(err)
	}

	if err := business.ps.
		Publish(ctx, common.TopicUserDisLikeRestaurant, pubsub.NewMessage(&restaurantlikemodel.Like{RestaurantId: restaurantId})); err != nil {
		log.Println(err)
	}
	return nil
}
