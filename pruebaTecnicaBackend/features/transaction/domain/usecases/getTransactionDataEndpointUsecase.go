package usecases

import (
	"pruebatecnica/pruebatecnicabackend/core/services"
	"pruebatecnica/pruebatecnicabackend/features/transaction/domain/models"
	"strings"
)

func GetTransactionDataEndpoint(url string) []models.Transaction {
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
