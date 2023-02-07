package userbusiness

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
	usermodel "go-food-delivery/module/user/model"
)

type RegisterStore interface {
	FindUser(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
	CreateUser(
		context context.Context,
		data *usermodel.UserCreate,
	) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStore RegisterStore
	hasher        Hasher
}

func NewRegisterBusiness(registerStore RegisterStore, hasher Hasher) *registerBusiness {
	return &registerBusiness{registerStore: registerStore, hasher: hasher}
}

func (business *registerBusiness) Register(
	context context.Context,
	data *usermodel.UserCreate,
) error {
	user, _ := business.registerStore.FindUser(context, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExitsed
	}

	salt := common.GetSalt(50)

	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err := business.registerStore.CreateUser(context, data); err != nil {
		return common.ErrCanNotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
