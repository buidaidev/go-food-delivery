package restaurantlikebusiness

import (
	"context"
	"go-food-delivery/common"
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
	context context.Context,
	data *restaurantlikemodel.Like,
) error {
	err := business.store.Create(context, data)

	if err != nil {
		return restaurantlikemodel.ErrCanNotLikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()

		if err := business.increaseStore.IncreaseLikeCount(context, data.RestaurantId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
