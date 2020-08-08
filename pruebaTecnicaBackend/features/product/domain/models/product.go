package models

type Product struct {
	Uid   string `json:"uid,omitempty"`
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price string `json:"price,omitempty"`
}

func NewProduct(id string, name string, price string) Product {
	product := Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
	return product
}
