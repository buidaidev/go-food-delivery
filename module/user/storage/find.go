package userstorage

import (
	"context"
	"go-food-delivery/common"
	usermodel "go-food-delivery/module/user/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var user usermodel.User

	if err := s.db.
		Where(condition).
		First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
