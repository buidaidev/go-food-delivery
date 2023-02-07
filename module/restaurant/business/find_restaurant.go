package restaurantbusiness

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
	usermodel "go-food-delivery/module/user/model"
)

type FindRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type findRestaurantBusiness struct {
	requester common.Requester
	store     FindRestaurantStore
}

func NewFindRestaurantBusiness(store FindRestaurantStore, requester common.Requester) *findRestaurantBusiness {
	return &findRestaurantBusiness{store: store, requester: requester}
}

func (business *findRestaurantBusiness) FindRestaurant(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	data, err := business.store.FindDataWithCondition(context, condition, usermodel.EntityName)

	if err != nil {
		return nil, common.ErrCanNotGetEntity(restaurantmodel.EntityName, err)
	}

	if data.UserId != business.requester.GetUserId() {
		return nil, common.ErrNoPermission(nil)
	}

	return data, nil
}
