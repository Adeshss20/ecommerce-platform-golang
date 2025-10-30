package models

type Product struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Price int64 `json:"price"`
	Quantity int32 `json:"quantity"`
}

type Order struct {
	Id string `json:"id"`
	ProductId string `json:"product-id"`
	Quantity int32 `json:"quantity"`
	BillId string `json:"bill-id"`
}

type Bill struct {
	Id string `json:"id"`
	OrderId string `json:"order-id"`
	TotalCost int64 `json:"total-cost"`
}