package restaurantlikebusiness

import (
	"context"
	"go-food-delivery/component/asyncjob"
	restaurantlikemodel "go-food-delivery/module/restaurantlike/model"
	"log"
)

type UserLikeRestaurantStore interface {
	Create(context context.Context, data *restaurantlikemodel.Like) error
}

type IncreaseLikeCountRestaurantStore interface {
	IncreaseLikeCount(context context.Context, id int) error
}

type userLikeRestaurantBusiness struct {
	store         UserLikeRestaurantStore
	increaseStore IncreaseLikeCountRestaurantStore
}

func NewUserLikeRestaurantBusiness(store UserLikeRestaurantStore, increaseStore IncreaseLikeCountRestaurantStore) *userLikeRestaurantBusiness {
	return &userLikeRestaurantBusiness{store: store, increaseStore: increaseStore}
}

func (business *userLikeRestaurantBusiness) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := business.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCanNotLikeRestaurant(err)
	}

	j := asyncjob.NewJob(func(ctx context.Context) error {
		return business.increaseStore.IncreaseLikeCount(ctx, data.RestaurantId)
	})

	if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
		log.Println(err)
	}

	return nil
}
