package restaurantbusiness

import (
	"context"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(
		context context.Context,
		data *restaurantmodel.RestaurantCreate,
	) error
}

type createRestaurantBusiness struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBusiness(store CreateRestaurantStore) *createRestaurantBusiness {
	return &createRestaurantBusiness{store: store}
}

func (business *createRestaurantBusiness) CreateRestaurant(
	context context.Context,
	data *restaurantmodel.RestaurantCreate,
) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := business.store.Create(context, data); err != nil {
		return err
	}

	return nil
}
