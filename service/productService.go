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
	for _, product := range products {
		if productId == product.Id {
			return &product,nil
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