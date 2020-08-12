package usecases

import (
	"encoding/json"
	"log"
	ProductUsecase "pruebatecnica/pruebatecnicabackend/features/product/domain/usecases"
)

type Response struct {
	SameIpBuyers    []Item `json:"sameIpBuyers"`
	PurchaseHistory []Item `json:"purchaseHistory"`
}
type Item struct {
	Item string `json:"item"`
}

func GetBuyerInfo(id string) []byte {

	sameIpBuyers := GetBuyerWhitSameIP(id)
	purchaseHistory := ProductUsecase.GetProductsHistory(id)
	var itemIp []Item
	var itemPurchase []Item

	for i := 0; i < len(sameIpBuyers); i++ {
		itemAux := Item{
			Item: sameIpBuyers[i],
		}
		itemIp = append(itemIp, itemAux)
	}

	for i := 0; i < len(purchaseHistory); i++ {
		itemAux := Item{
			Item: purchaseHistory[i],
		}
		itemPurchase = append(itemPurchase, itemAux)
	}

	buyerInfo := Response{
		SameIpBuyers:    itemIp,
		PurchaseHistory: itemPurchase,
	}

	pb, err := json.Marshal(buyerInfo)
	if err != nil {
		log.Fatal(err)
	}

	return pb
}
