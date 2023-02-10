package restaurantrepository

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
	usermodel "go-food-delivery/module/user/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantRepository struct {
	store ListRestaurantStore
}

func NewListRestaurantRepository(store ListRestaurantStore) *listRestaurantRepository {
	return &listRestaurantRepository{store: store}
}

func (repository *listRestaurantRepository) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := repository.store.ListDataWithCondition(context, filter, paging, usermodel.EntityName)

	if err != nil {
		return nil, common.ErrCanNotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
