package usecases

import (
	"encoding/json"
	"fmt"
	"log"
	BuyerModel "pruebatecnica/pruebatecnicabackend/features/buyer/domain/models"
	ProductModel "pruebatecnica/pruebatecnicabackend/features/product/domain/models"
	ProductUsecase "pruebatecnica/pruebatecnicabackend/features/product/domain/usecases"
)

type Response struct {
	SameIpBuyers    []BuyerModel.Buyer     `json:"sameIpBuyers"`
	PurchaseHistory []ProductModel.Product `json:"purchaseHistory"`
	Recomendations  []ProductModel.Product `json:"recomendations"`
}

func GetBuyerInfo(id string) []byte {

	sameIpBuyers := GetBuyerWhitSameIP(id)
	purchaseHistory := ProductUsecase.GetProductsHistory(id)
	var recomendations []ProductModel.Product

	if len(sameIpBuyers) > 0 {
		id := sameIpBuyers[0].Id
		fmt.Println(id)
		recomendations = ProductUsecase.GetProductsHistory(sameIpBuyers[0].Id)
	}

	buyerInfo := Response{
		SameIpBuyers:    sameIpBuyers,
		PurchaseHistory: purchaseHistory,
		Recomendations:  recomendations,
	}

	pb, err := json.Marshal(buyerInfo)
	if err != nil {
		log.Fatal(err)
	}

	return pb
}
