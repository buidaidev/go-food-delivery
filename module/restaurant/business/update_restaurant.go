package restaurantbusiness

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

type UpdateRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	Update(
		context context.Context,
		data *restaurantmodel.RestaurantUpdate,
		id int,
	) error
}

type updateRestaurantBusiness struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBusiness(store UpdateRestaurantStore) *updateRestaurantBusiness {
	return &updateRestaurantBusiness{store: store}
}

func (business *updateRestaurantBusiness) UpdateRestaurant(context context.Context, data *restaurantmodel.RestaurantUpdate, id int) error {
	oldData, err := business.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCanNotGetEntity(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, err)
	}

	if err := business.store.Update(context, data, id); err != nil {
		return common.ErrCanNotUpdateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
