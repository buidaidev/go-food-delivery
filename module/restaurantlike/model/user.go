package restaurantlikemodel

import (
	"go-food-delivery/common"
	"time"
)

type User struct {
	common.SimpleUser
	LikeAt *time.Time `json:"created_at,omitempty" gorm:"created_at"`	
}
