package restaurantstorage

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

func (s *sqlStore) Create(
	context context.Context,
	data *restaurantmodel.RestaurantCreate,
) error {
	if err := s.db.
		Table(restaurantmodel.RestaurantCreate{}.TableName()).
		Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
