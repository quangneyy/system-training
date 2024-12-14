package controller

import (
	"bankingSystem/common"
	"bankingSystem/module/product/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (api APIController) CreateProductAPI(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check & parse data from body
		var productData domain.ProductCreationDTO

		if err := c.Bind(&productData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		productData.Id = common.GenUUID()

		if err := api.createUseCase.CreateProduct(c.Request.Context(), &productData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// response to client
		c.JSON(http.StatusCreated, gin.H{"data": productData.Id})
	}
}
