package restaurantbusiness

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

type ListRestaurantRepository interface {
	ListRestaurant(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBusiness struct {
	repository ListRestaurantRepository
}

func NewListRestaurantBusiness(repository ListRestaurantRepository) *listRestaurantBusiness {
	return &listRestaurantBusiness{repository: repository}
}

func (business *listRestaurantBusiness) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := business.repository.ListRestaurant(context, filter, paging)

	if err != nil {
		return nil, common.ErrCanNotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
