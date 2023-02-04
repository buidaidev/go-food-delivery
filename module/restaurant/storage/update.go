package restaurantstorage

import (
	"context"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

func (s *sqlStore) Update(
	context context.Context,
	data *restaurantmodel.RestaurantUpdate,
	id int,
) error {
	if err := s.db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
