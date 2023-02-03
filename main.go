package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:address;"`
}

type RestaurantCreate struct {
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:address;"`
}

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:address;"`
}

func (Restaurant) TableName() string { return "restaurants" }

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

	r := gin.Default()
	v1 := r.Group("/v1")

	restaurants := v1.Group("restaurants")
	{
		restaurants.POST("/", func(ctx *gin.Context) {
			var data RestaurantCreate

			if err := ctx.ShouldBind(&data); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			db.Create(&data)

			ctx.JSON(http.StatusOK, gin.H{
				"data": data,
			})
		})

		restaurants.GET("/", func(ctx *gin.Context) {
			var data []Restaurant

			type Paging struct {
				Page  int `json:"page" form:"page"`
				Limit int `json:"limit" form:"limit"`
			}

			var pagingData Paging

			if err := ctx.ShouldBind(&pagingData); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			if pagingData.Page <= 0 {
				pagingData.Page = 1
			}

			if pagingData.Limit <= 0 {
				pagingData.Limit = 5
			}

			db.
				Offset((pagingData.Page - 1) * pagingData.Limit).
				Order("id desc").
				Limit(pagingData.Limit).
				Find(&data)

			ctx.JSON(http.StatusOK, gin.H{
				"data": data,
			})
		})

		restaurants.GET("/:id", func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Param("id"))

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			var data Restaurant

			db.
				Where("id = ?", id).
				First(&data)

			ctx.JSON(http.StatusOK, gin.H{
				"data": data,
			})
		})

		restaurants.PATCH("/:id", func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Param("id"))

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			var data RestaurantUpdate

			if err := ctx.ShouldBind(&data); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			db.
				Where("id = ?", id).
				Updates(&data)

			ctx.JSON(http.StatusOK, gin.H{
				"data": data,
			})
		})

		restaurants.DELETE("/:id", func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Param("id"))

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			db.
				Table(Restaurant{}.TableName()).
				Where("id = ?", id).
				Delete(nil)

			ctx.JSON(http.StatusOK, gin.H{
				"data": 1,
			})
		})
	}

	r.Run()

}
