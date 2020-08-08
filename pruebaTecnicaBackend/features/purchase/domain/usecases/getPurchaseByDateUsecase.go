package usecases

import (
	"encoding/json"
	"log"
	dgrahpdata "pruebatecnica/pruebatecnicabackend/core/dgraphdata"
	"pruebatecnica/pruebatecnicabackend/core/services"
	"pruebatecnica/pruebatecnicabackend/features/purchase/constants"
	"pruebatecnica/pruebatecnicabackend/features/purchase/domain/models"

	UsecaseBuyer "pruebatecnica/pruebatecnicabackend/features/buyer/domain/usecases"
	UsecaseProduct "pruebatecnica/pruebatecnicabackend/features/product/domain/usecases"
	UsecaseTransaction "pruebatecnica/pruebatecnicabackend/features/transaction/domain/usecases"
)

func GetPurchaseByDate(date string) []byte {
	unixDate := services.DateToUnix(date)
	buyersData := UsecaseBuyer.GetBuyerDataEndpoint(constants.URL + constants.BUYERS + unixDate)
	productsData := UsecaseProduct.GetProductDataEndpoint(constants.URL + constants.PRODUCTS + unixDate)
	transactionsData := UsecaseTransaction.GetTransactionDataEndpoint(constants.URL + constants.TRANSACTIONS + unixDate)

	purchases := models.NewPurchases(buyersData, productsData, transactionsData)

	pb, error := json.Marshal(purchases)
	if error != nil {
		log.Fatal(error)
	}
	dgrahpdata.UploadData(pb)
	return pb
}
