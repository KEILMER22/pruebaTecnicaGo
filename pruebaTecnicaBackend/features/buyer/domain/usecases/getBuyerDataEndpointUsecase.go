package usecases

import (
	"encoding/json"
	"log"
	"pruebatecnica/pruebatecnicabackend/core/services"
	"pruebatecnica/pruebatecnicabackend/features/buyer/domain/models"
)

func GetBuyerDataEndpoint(url string) []models.Buyer {
	var buyers []models.Buyer
	data := services.LoadData(url)
	error := json.Unmarshal(data, &buyers)
	if error != nil {
		log.Fatal(error)
	}
	return buyers
}
