package service

import (
	"fmt"
	"github.com/Adeshss20/ecommerce-platform-golang.git/models"
	"github.com/google/uuid"
)

var orders []models.Order

type OrderSerice struct {

}

func NewOrderService() *OrderSerice {
	return &OrderSerice{}
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
	return order.Id, nil
}

func (service *OrderSerice) GetAllOrders() ([]models.Order, error) {
	return orders, nil
}