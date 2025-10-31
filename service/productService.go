package service

import (
	"fmt"
	"github.com/Adeshss20/ecommerce-platform-golang.git/models"
	"github.com/google/uuid"
)

var products []models.Product

type ProductService struct {

}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (service *ProductService) GetProduct(productId string) (*models.Product, error) {
	for i := range products {
		if productId == products[i].Id {
			return &products[i],nil
		}
	}
	return nil, fmt.Errorf("product not found with product id - %s", productId)
}

func (service *ProductService) AddProduct(product models.Product) (string, error) {
	product.Id = uuid.New().String()
	products = append(products, product)
	return product.Id, nil
}

func (service *ProductService) GetAllProducts() ([]models.Product, error) {
	return products, nil
}

func (service *ProductService) CheckProductAvailibility(productId string, quantity int32) bool {
	product, err := service.GetProduct(productId)

	if err!=nil {
		return false
	}

	return (product.Quantity-quantity)>=0
}

func (service *ProductService) ReduceQuantity(productId string, quantity int32) {
	product, _ := service.GetProduct(productId)
	(*product).Quantity = (*product).Quantity-quantity
}