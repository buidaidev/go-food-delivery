package restaurantlikestorage

import (
	"context"
	"go-food-delivery/common"
	restaurantlikemodel "go-food-delivery/module/restaurantlike/model"
)

func (s *sqlStore) Create(
	context context.Context,
	data *restaurantlikemodel.Like,
) error {
	if err := s.db.
		Table(restaurantlikemodel.Like{}.TableName()).
		Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
