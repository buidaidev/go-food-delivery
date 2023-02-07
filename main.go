package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	component "go-food-delivery/component/appctx"
	"go-food-delivery/component/uploadprovider"
	"go-food-delivery/middleware"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")
	secretKey := os.Getenv("SECRET_KEY")
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3BucketName")
	s3APIKey := os.Getenv("S3BucketName")
	s3SecretKey := os.Getenv("S3BucketName")
	s3Domain := os.Getenv("S3BucketName")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()
	appContext := component.NewAppContext(db, s3Provider, secretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))
	r.Static("/static", "./static")

	v1 := r.Group("/v1")
	setupRoutes(appContext, v1)
	setupAdminRoutes(appContext, v1)

	r.Run()

}
