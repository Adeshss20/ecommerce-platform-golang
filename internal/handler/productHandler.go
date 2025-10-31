package handler

import (
	"fmt"
	"net/http"
	"github.com/Adeshss20/ecommerce-platform-golang.git/models"
	"github.com/Adeshss20/ecommerce-platform-golang.git/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (handler *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("product-id")
	product, err := handler.service.GetProduct(id)

	if err!=nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message" : fmt.Sprintf("product not found for product id - %s", id)})
	}

	c.JSON(http.StatusOK, gin.H{ "product-details" : product})
}

func (handler *ProductHandler) AddProduct(c *gin.Context) {
	var req models.Product

	err := c.ShouldBindJSON(&req)
	if err!=nil {
		c.JSON(http.StatusBadRequest,
		gin.H{
			"error" : err.Error(), 
		})
	}

	id, _ := handler.service.AddProduct(req)
	c.JSON(http.StatusCreated,
		gin.H{"message" : fmt.Sprintf("Product added with id - %s", id)})
}

func (handler *ProductHandler) GetAllProducts(c *gin.Context) {
	products, _ := handler.service.GetAllProducts()
	c.JSON(http.StatusOK,
		gin.H{
			"products" : products,
		})
}