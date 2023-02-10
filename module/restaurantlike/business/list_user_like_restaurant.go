package restaurantlikebusiness

import (
	"context"
	"go-food-delivery/common"
	restaurantlikemodel "go-food-delivery/module/restaurantlike/model"
)

type ListUserLikeRestaurantStore interface {
	GetUserLikeRestaurant(
		ctx context.Context,
		condition map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBusiness struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBusiness(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBusiness {
	return &listUserLikeRestaurantBusiness{store: store}
}

func (business *listUserLikeRestaurantBusiness) ListUsersLikeRestaurant(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	users, err := business.store.GetUserLikeRestaurant(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCanNotListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil

}
