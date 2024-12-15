package main

import (
	"bankingSystem/common"
	"bankingSystem/module/product/controller"
	productusecase "bankingSystem/module/product/domain/usecase"
	productmysql "bankingSystem/module/product/repository/mysql"
	"bankingSystem/module/user/infras/httpservice"
	"bankingSystem/module/user/infras/repository"
	"bankingSystem/module/user/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Setup dependencies
	repo := productmysql.NewMysqlRepository(db)
	useCase := productusecase.NewCreateProductUseCase(repo)
	api := controller.NewAPIController(useCase)

	v1 := r.Group("/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", api.CreateProductAPI(db))
		}
	}

	userUC := usecase.NewUseCase(repository.NewUserRepo(db), &common.Hasher{})
	httpservice.NewUserService(userUC).Routes(v1)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
