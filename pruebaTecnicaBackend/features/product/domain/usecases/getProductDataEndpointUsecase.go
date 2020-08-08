package usecases

import (
	"pruebatecnica/pruebatecnicabackend/core/services"
	"pruebatecnica/pruebatecnicabackend/features/product/domain/models"
	"strings"
)

func GetProductDataEndpoint(url string) []models.Product {
	var productsData []models.Product
	productsString := string(services.LoadData(url))
	products := strings.Split(productsString, "\n")
	for j := 0; j < len(products); j++ {
		//for j := 0; j < 5; j++ {
		product := products[j]
		sliceData := strings.Split(product, "'")
		if len(sliceData) == 3 {
			producto := models.NewProduct(sliceData[0], sliceData[1], sliceData[2])
			productsData = append(productsData, producto)
		}
	}
	return productsData
}
