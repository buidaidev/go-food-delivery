package restaurantlikebusiness

import (
	"context"
	"go-food-delivery/common"
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
	context context.Context,
	userId, restaurantId int,
) error {
	err := business.store.Delete(context, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCanNotDislikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover()

		if err := business.decreaseStore.DecreaseLikeCount(context, restaurantId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
