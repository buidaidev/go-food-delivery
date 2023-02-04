package uploadbusiness

import (
	"bytes"
	"context"
	"fmt"
	"go-food-delivery/common"
	"go-food-delivery/component/uploadprovider"
	uploadmodel "go-food-delivery/module/upload/model"
	"image"
	"io"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStore interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBusiness struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStore
}

func NewUploadBusiness(provider uploadprovider.UploadProvider, imgStore CreateImageStore) *uploadBusiness {
	return &uploadBusiness{provider: provider, imgStore: imgStore}
}

func (business *uploadBusiness) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := business.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCanNotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)

	if err != nil {
		return 0, 0, nil
	}

	return img.Width, img.Height, nil
}
