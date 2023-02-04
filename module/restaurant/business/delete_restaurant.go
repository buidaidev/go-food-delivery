package restaurantbusiness

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error
}

type deleteRestaurantBusiness struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBusiness(store DeleteRestaurantStore) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{store: store}
}

func (business *deleteRestaurantBusiness) DeleteRestaurant(context context.Context, id int) error {
	oldData, err := business.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, err)
	}

	if err := business.store.Delete(context, id); err != nil {
		return common.ErrCanNotDeleteEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
