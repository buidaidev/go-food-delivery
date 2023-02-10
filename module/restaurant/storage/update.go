package restaurantstorage

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"

	"gorm.io/gorm"
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
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikeCount(context context.Context, id int) error {
	db := s.db

	if err := db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikeCount(context context.Context, id int) error {
	db := s.db

	if err := db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
