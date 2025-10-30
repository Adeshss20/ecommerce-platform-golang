package handler

import (
	"fmt"
	"net/http"
	"github.com/Adeshss20/ecommerce-platform-golang.git/models"
	"github.com/Adeshss20/ecommerce-platform-golang.git/service"
	"github.com/gin-gonic/gin"
)

type BillHandler struct {
	service *service.BillService
}

func NewBillHandler(service *service.BillService) *BillHandler {
	return &BillHandler{service: service}
}

func (handler *BillHandler) GetBillForOrder(c *gin.Context) {
	orderId := c.Param("order-id")
	bill, err := handler.service.GetBillForOrder(orderId)

	if err!=nil {
		c.JSON(http.StatusOK,
			gin.H{
				"message" : fmt.Sprintf("Bill not found for order id - %s", orderId),
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bill-detail" : bill})
}

func (handler *BillHandler) AddBill(c *gin.Context) {
	var req models.Bill

	err := c.ShouldBindJSON(&req) 
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

	id, _ := handler.service.AddBill(req)

	c.JSON(http.StatusCreated,
		gin.H{
			"message" : fmt.Sprintf("Bill created with id - %s", id),
		})
}

func (handler *BillHandler) GetAllBills(c *gin.Context) {
	bills, _ := handler.service.GetAllBills()
	c.JSON(http.StatusAccepted,
	gin.H{
		"bills" : bills,
	})
}