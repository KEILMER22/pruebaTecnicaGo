package usecases

import (
	"encoding/json"
	"log"
	dgrahpdata "pruebatecnica/pruebatecnicabackend/core/dgraphdata"
	"pruebatecnica/pruebatecnicabackend/features/product/domain/models"
	"pruebatecnica/pruebatecnicabackend/features/product/query"
)

type Response struct {
	Response []models.Product
}

func GetProductsHistory(id string) []models.Product {
	var data dgrahpdata.AuxResponse
	var product Response
	var products []models.Product
	response := dgrahpdata.BdQueryWithVars(query.GETPRODUCTSIDBYBUYERID, id).Json
	error := json.Unmarshal(response, &data)
	if error != nil {
		log.Fatal(error)
	}
	if len(data.Response) != 0 {
		for j := 0; j < len(data.Response); j++ {
			productIds := data.Response[j].ResponseString
			for i := 0; i < len(productIds); i++ {
				response = dgrahpdata.BdQueryWithVars(query.GETPRODUCT, productIds[i]).Json
				error := json.Unmarshal(response, &product)
				if error != nil {
					log.Fatal(error)
				}
				if len(product.Response) != 0 {
					products = append(products, product.Response[0])
				}
			}
		}
	}
	return products
}
