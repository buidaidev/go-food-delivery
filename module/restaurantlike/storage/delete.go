package restaurantlikestorage

import (
	"context"
	"go-food-delivery/common"
	restaurantlikemodel "go-food-delivery/module/restaurantlike/model"
)

func (s *sqlStore) Delete(
	context context.Context,
	userId, restaurantId int,
) error {
	if err := s.db.
		Table(restaurantlikemodel.Like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
