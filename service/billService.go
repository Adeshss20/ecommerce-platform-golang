package service

import (
	"fmt"
	"github.com/Adeshss20/ecommerce-platform-golang.git/models"
	"github.com/google/uuid"
)

var bills []models.Bill

type BillService struct {
	
}

func NewBillService() *BillService {
	return &BillService{}
}

func (service *BillService) GetBillForOrder(orderId string) (*models.Bill, error) {
	for _, bill := range bills {
		if(bill.OrderId==orderId) {
			return &bill, nil;
		}
	}
	return nil, fmt.Errorf("bill not found for order id - %s", orderId)
}

func (service *BillService) AddBill(bill models.Bill) (string, error) {
	bill.Id = uuid.New().String()
	bills = append(bills, bill)
	return bill.Id, nil
}

func (service *BillService) GetAllBills() ([]models.Bill, error) {
	return bills, nil
}

func (service *BillService) GenerateBill(orderId string, totalCost int64) *models.Bill {
	id := uuid.New().String()
	return &models.Bill{Id: id, OrderId: orderId, TotalCost: totalCost}
}