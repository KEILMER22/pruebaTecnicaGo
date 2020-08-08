package models

type Transaction struct {
	Uid        string   `json:"uid,omitempty"`
	Id         string   `json:"id,omitempty"`
	BuyerID    string   `json:"buyer_id,omitempty"`
	Ip         string   `json:"ip,omitempty"`
	Device     string   `json:"device,omitempty"`
	ProductIds []string `json:"product_ids,omitempty"`
}

func NewTransaction(id string, buyerID string, ip string, device string, productsIds []string) Transaction {
	transaction := Transaction{
		Id:         id,
		BuyerID:    buyerID,
		Ip:         ip,
		Device:     device,
		ProductIds: productsIds,
	}
	return transaction
}
