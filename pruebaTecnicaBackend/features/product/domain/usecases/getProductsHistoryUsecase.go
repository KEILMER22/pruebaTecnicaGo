package usecases

import (
	"encoding/json"
	"log"
	dgrahpdata "pruebatecnica/pruebatecnicabackend/core/dgraphdata"
	"pruebatecnica/pruebatecnicabackend/features/product/query"
)

func GetProductsHistory(id string) []string {
	var data dgrahpdata.AuxResponse
	var product dgrahpdata.AuxResponseSingle
	var StringProduct []string
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
					StringProduct = append(StringProduct, product.Response[0].ResponseString)
				}
			}
		}
	}
	StringProduct = dgrahpdata.RemoveDuplicatesUnordered(StringProduct)
	return StringProduct
}
