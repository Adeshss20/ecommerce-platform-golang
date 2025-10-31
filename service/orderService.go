package service

import (
	"fmt"
	"github.com/Adeshss20/ecommerce-platform-golang.git/models"
	"github.com/google/uuid"
)

var orders []models.Order

type OrderSerice struct {
	productService *ProductService
	billService *BillService
}

func NewOrderService(productService *ProductService, billService *BillService) *OrderSerice {
	return &OrderSerice{ productService: productService, billService: billService}
}

func (service *OrderSerice) GetOrder(orderId string) (*models.Order, error) {
	for _, order := range orders {
		if order.Id==orderId {
			return &order, nil
		}
	}
	return nil, fmt.Errorf("order not found for order id - %s", orderId)
}

func (service *OrderSerice) AddOrder(order models.Order) (string, error) {
	order.Id = uuid.New().String()
	orders = append(orders, order)
	totalCost, _ := service.GetTotalCost(order)
	service.billService.GenerateBill(order.Id, totalCost)
	return order.Id, nil
}

func (service *OrderSerice) GetAllOrders() ([]models.Order, error) {
	return orders, nil
}

func (service *OrderSerice) GetTotalCost(order models.Order) (int64, error) {
	productService := service.productService

	product, err := productService.GetProduct(order.ProductId)
	if err!=nil {
		return 0, fmt.Errorf("unable to fetch product for order - %s", order.ProductId)
	}

	totalCost := (order.Quantity * int32(product.Price))

	return int64(totalCost), nil
}