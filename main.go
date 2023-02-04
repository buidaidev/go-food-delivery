package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-food-delivery/component/appctx"
	ginrestaurant "go-food-delivery/module/restaurant/transport/gin"
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()
	appContext := appctx.NewAppContext(db)

	r := gin.Default()
	v1 := r.Group("/v1")
	restaurants := v1.Group("restaurants")
	{
		restaurants.POST("/", ginrestaurant.CreateRestaurant(appContext))

		restaurants.GET("/", ginrestaurant.ListRestaurant(appContext))

		restaurants.GET("/:id", ginrestaurant.FindRestaurant(appContext))

		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appContext))

		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
	}

	r.Run()

}
