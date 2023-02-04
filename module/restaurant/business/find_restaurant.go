package restaurantbusiness

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

type FindRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type findRestaurantBusiness struct {
	store FindRestaurantStore
}

func NewFindRestaurantBusiness(store FindRestaurantStore) *findRestaurantBusiness {
	return &findRestaurantBusiness{store: store}
}

func (business *findRestaurantBusiness) FindRestaurant(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	data, err := business.store.FindDataWithCondition(context, condition)

	if err != nil {
		return nil, common.ErrCanNotGetEntity(restaurantmodel.EntityName, err)
	}

	return data, nil
}
