package usecases

import (
	"encoding/json"
	"log"
	dgrahpdata "pruebaTecnica/pruebatecnicabackend/core/dgraphdata"
	"pruebaTecnica/pruebatecnicabackend/core/services"
	"pruebaTecnica/pruebatecnicabackend/features/purchase/domain/models"
	"strings"

	"github.com/dgraph-io/dgo/protos/api"
)

const (
	layoutISO    = "2006-01-02"
	URL          = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/"
	BUYERS       = "buyers"
	PRODUCTS     = "products"
	TRANSACTIONS = "transactions"
)

func GetPurchaseByDate(date string) []byte {
	unixDate := services.DateToUnix(date)

	buyersData := getBuyerData(URL + BUYERS + unixDate)
	productsData := getProductData(URL + PRODUCTS + unixDate)
	transactionsData := getTransactionData(URL + TRANSACTIONS + unixDate)

	purchases := models.NewPurchases("123", buyersData, productsData, transactionsData)

	pb, error := json.Marshal(purchases)
	if error != nil {
		log.Fatal(error)
	}
	dgrahpdata.UploadData(pb)
	return pb
}

func GetBuyerInfo(id string) {
	dgrahpdata.IpBuyers(id)
}

func getBuyerData(url string) []models.Buyer {
	var buyers []models.Buyer
	data := services.LoadData(url)
	error := json.Unmarshal(data, &buyers)
	if error != nil {
		log.Fatal(error)
	}
	return buyers
}

func GetBuyers() *api.Response {
	return dgrahpdata.BuyersQuery()
}

func getProductData(url string) []models.Product {
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

func getTransactionData(url string) []models.Transaction {
	transactionsData := string(services.LoadData(url))
	transactionsDataFragments := strings.Split(string(transactionsData), "#")
	var transactions []models.Transaction

	for i := 0; i < len(transactionsDataFragments); i++ {
		//for i := 0; i < 5; i++ {
		transaction := string(transactionsDataFragments[i])
		partsBad := strings.Replace(transaction, "\x00", " ", -1)

		parts := strings.Split(partsBad, " ")
		if len(parts) > 1 {

			products := strings.Replace(parts[4], "(", "", 1)
			products = strings.Replace(products, ")", "", 1)
			producsIds := strings.Split(products, ",")

			transactionModel := models.NewTransaction(parts[0], parts[1], parts[2], parts[3], producsIds)
			transactions = append(transactions, transactionModel)

		}

	}
	return transactions
}
