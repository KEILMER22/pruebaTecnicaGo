package models

import (
	Modelbuyer "pruebatecnica/pruebatecnicabackend/features/buyer/domain/models"
	Modelproduct "pruebatecnica/pruebatecnicabackend/features/product/domain/models"
	Modeltransaction "pruebatecnica/pruebatecnicabackend/features/transaction/domain/models"
)

type Purchases struct {
	Buyers       []Modelbuyer.Buyer             `json:"Buyers,omitempty"`
	Products     []Modelproduct.Product         `json:"Products,omitempty"`
	Transactions []Modeltransaction.Transaction `json:"Transactions,omitempty"`
}

func NewPurchases(buyers []Modelbuyer.Buyer, products []Modelproduct.Product, transactions []Modeltransaction.Transaction) Purchases {
	purchases := Purchases{
		Buyers:       buyers,
		Products:     products,
		Transactions: transactions,
	}
	return purchases
}
