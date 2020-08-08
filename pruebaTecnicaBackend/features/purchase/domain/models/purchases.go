package models

type Purchases struct {
	Cod          string        `json:"cod,omitempty"`
	Buyers       []Buyer       `json:"Buyers,omitempty"`
	Products     []Product     `json:"Products,omitempty"`
	Transactions []Transaction `json:"Transactions,omitempty"`
}

func NewPurchases(uid string, buyers []Buyer, products []Product, transactions []Transaction) Purchases {

	purchases := Purchases{
		Cod:          uid,
		Buyers:       buyers,
		Products:     products,
		Transactions: transactions,
	}
	return purchases
}
