package uploadmodel

import (
	"go-food-delivery/common"
)

type Upload struct {
	common.SQLModel
	common.Image
}

func (Upload) TableName() string {
	return "uploads"
}

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(err, "file is not image", "ErrorFileIsNotImage")
}

func ErrCanNotSaveFile(err error) *common.AppError {
	return common.NewCustomError(err, "cannot save uploaded file", "ErrorCanNotSaveFile")
}

func ErrFileTooLarge(err error) *common.AppError {
	return common.NewCustomError(err, "file too large", "ErrFileTooLarge")
}
