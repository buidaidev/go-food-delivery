package restaurantlikebusiness

import (
	"context"
	"go-food-delivery/component/asyncjob"
	restaurantlikemodel "go-food-delivery/module/restaurantlike/model"
	"log"
)

type UserDislikeRestaurantStore interface {
	Delete(context context.Context, userId, restaurantId int) error
}

type DecreaseLikeCountRestaurantStore interface {
	DecreaseLikeCount(context context.Context, id int) error
}

type userDislikeRestaurantBusiness struct {
	store         UserDislikeRestaurantStore
	decreaseStore DecreaseLikeCountRestaurantStore
}

func NewUserDislikeRestaurantBusiness(store UserDislikeRestaurantStore, decreaseStore DecreaseLikeCountRestaurantStore) *userDislikeRestaurantBusiness {
	return &userDislikeRestaurantBusiness{store: store, decreaseStore: decreaseStore}
}

func (business *userDislikeRestaurantBusiness) DislikeRestaurant(
	ctx context.Context,
	userId, restaurantId int,
) error {
	err := business.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCanNotDislikeRestaurant(err)
	}

	j := asyncjob.NewJob(func(ctx context.Context) error {
		return business.decreaseStore.DecreaseLikeCount(ctx, restaurantId)
	})

	if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
		log.Println(err)
	}

	return nil
}
