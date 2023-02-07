package userstorage

import (
	"context"
	"go-food-delivery/common"
	usermodel "go-food-delivery/module/user/model"
)

func (s *sqlStore) CreateUser(
	context context.Context,
	data *usermodel.UserCreate,
) error {
	db := s.db.Begin()

	if err := db.
		Table(data.TableName()).
		Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
