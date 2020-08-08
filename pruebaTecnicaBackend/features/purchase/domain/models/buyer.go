package models

// Buyer: Type for the list Buyer
type Buyer struct {
	Uid  string `json:"uid,omitempty"`
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func NewBuyer(id string, name string, age int) Buyer {
	buyer := Buyer{
		Id:   id,
		Name: name,
		Age:  age,
	}
	return buyer
}
