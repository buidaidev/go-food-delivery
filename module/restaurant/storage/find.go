package restaurantstorage

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant

	if err := s.db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where(condition).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(restaurantmodel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
