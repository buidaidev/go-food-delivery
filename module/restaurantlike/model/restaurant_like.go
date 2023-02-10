package restaurantlikemodel

import (
	"fmt"
	"go-food-delivery/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at	;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Like) TableName() string { return "restaurant_likes" }

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

func ErrCanNotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Can not like restaurant"),
		fmt.Sprintf("ErrorCanNotLikeRestaurant"),
	)
}

func ErrCanNotDislikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Can not unlike restaurant"),
		fmt.Sprintf("ErrorCanNotDislikeRestaurant"),
	)
}
