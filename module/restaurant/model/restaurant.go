package restaurantmodel

import (
	"errors"
	"go-food-delivery/common"
	"strings"
)

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:address;"`
	Status  int    `json:"status" gorm:"column:status;"`
}

type RestaurantCreate struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:address;"`
}

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:address;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

var ErrNameIsEmpty = errors.New("Nam can not be empty.")
