package restaurantrepository

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
	usermodel "go-food-delivery/module/user/model"
	"log"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantRepository struct {
	store     ListRestaurantStore
	likeStore LikeRestaurantStore
}

func NewListRestaurantRepository(store ListRestaurantStore, likeStore LikeRestaurantStore) *listRestaurantRepository {
	return &listRestaurantRepository{store: store, likeStore: likeStore}
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

	ids := make([]int, len(result))

	for i := range ids {
		ids[i] = result[i].Id
	}

	likeMap, err := repository.likeStore.GetRestaurantLikes(context, ids)

	if err != nil {
		log.Println(err)
		return result, nil
	}

	for i, item := range result {
		result[i].LikeCount = likeMap[item.Id]
	}

	return result, nil
}
