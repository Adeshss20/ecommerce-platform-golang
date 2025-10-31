package router

import (
	"github.com/Adeshss20/ecommerce-platform-golang.git/internal/handler"
	"github.com/Adeshss20/ecommerce-platform-golang.git/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	productService := service.NewProductService()
	productHandler := handler.NewProductHandler(productService)

	productApi := r.Group("/product") 
	{
		productApi.GET("/", productHandler.GetAllProducts)
		productApi.GET("/:product-id", productHandler.GetProduct)
		productApi.POST("/", productHandler.AddProduct)
	}

	billService := service.NewBillService()
	billHandler := handler.NewBillHandler(billService)

	billApi := r.Group("/bill") 
	{
		billApi.GET("/", billHandler.GetAllBills)
		billApi.POST("/", billHandler.AddBill)
		billApi.GET("/:order-id", billHandler.GetBillForOrder)
	}

	orderService := service.NewOrderService(productService, billService)
	OrderHandler := handler.NewOrderHandler(orderService)

	orderApi := r.Group("/order") 
	{
		orderApi.GET("/", OrderHandler.GetAllOrders)
		orderApi.GET("/:order-id", OrderHandler.GetOrder)
		orderApi.POST("/", OrderHandler.AddOrder)
	}

	return r
}