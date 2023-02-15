package component

import (
	"go-food-delivery/component/uploadprovider"
	"go-food-delivery/pubsub"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, uploadProvider uploadprovider.UploadProvider, secretKey string, ps pubsub.Pubsub) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider, secretKey: secretKey, ps: ps}
}

func (ctx *appCtx) GetMaiDBConnection() *gorm.DB                  { return ctx.db }
func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider { return ctx.uploadProvider }
func (ctx *appCtx) SecretKey() string                             { return ctx.secretKey }
func (ctx *appCtx) GetPubSub() pubsub.Pubsub                      { return ctx.ps }
