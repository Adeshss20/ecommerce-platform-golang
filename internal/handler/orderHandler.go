package handler

import (
	"fmt"
	"net/http"
	"github.com/Adeshss20/ecommerce-platform-golang.git/models"
	"github.com/Adeshss20/ecommerce-platform-golang.git/service"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *service.OrderSerice
}

func NewOrderHandler(service *service.OrderSerice) *OrderHandler {
	return &OrderHandler{service: service}
}

func (handler *OrderHandler) GetOrder(c *gin.Context) {
	orderId := c.Param("order-id")
	order, err := handler.service.GetOrder(orderId)
	
	if err!=nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"message" : fmt.Sprintf("order not found for order id - %s", orderId),
			})
	}

	c.JSON(http.StatusOK, gin.H{"order-detail" : *order})
}

func (handler *OrderHandler) AddOrder(c *gin.Context) {
	var req models.Order

	err := c.ShouldBindJSON(&req)
	if err!=nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error" : err.Error(),
			})
		return
	}

	id, err := handler.service.AddOrder(req)
	if err!=nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error" : err.Error(),
			})
		return
	}
	c.JSON(http.StatusCreated,
		gin.H{"message" : fmt.Sprintf("Order created with id - %s", id)})
}

func (handler *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, _ := handler.service.GetAllOrders()
	c.JSON(http.StatusOK,
		gin.H{
			"orders" : orders,
		})
}