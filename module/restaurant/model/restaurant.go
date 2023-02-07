package restaurantmodel

import (
	"errors"
	"go-food-delivery/common"
	"strings"
)

const EntityName = "Restaurant"

var ErrNameIsEmpty = errors.New("Nam can not be empty.")

type Restaurant struct {
	common.SQLModel
	Name    string             `json:"name" gorm:"column:name;"`
	Address string             `json:"address" gorm:"column:address;"`
	Status  int                `json:"status" gorm:"column:status;"`
	Logo    *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images     `json:"cover" gorm:"column:cover;"`
	User    *common.SimpleUser `json:"user" gorm:"preload:false;" `
	UserId  int                `json:"-" gorm:"column:user_id"`
}

type RestaurantCreate struct {
	common.SQLModel
	Name    string         `json:"name" gorm:"column:name;"`
	Address string         `json:"address" gorm:"column:address;"`
	UserId  int            `json:"-" gorm:"column:user_id"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

type RestaurantUpdate struct {
	Name    *string        `json:"name" gorm:"column:name;"`
	Address *string        `json:"address" gorm:"column:address;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}
