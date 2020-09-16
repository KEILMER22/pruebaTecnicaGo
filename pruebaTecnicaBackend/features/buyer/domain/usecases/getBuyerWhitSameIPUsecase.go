package usecases

import (
	"encoding/json"
	"log"
	dgrahpdata "pruebatecnica/pruebatecnicabackend/core/dgraphdata"
	"pruebatecnica/pruebatecnicabackend/features/buyer/domain/models"
	"pruebatecnica/pruebatecnicabackend/features/buyer/query"
	"strconv"
)

type ResponseBuyer struct {
	Response []BuyerA `json:"Response"`
}
type BuyerA struct {
	Name string `json:"Name,omitempty"`
	Age  string `json:"Age,omitempty"`
	Id   string `json:"Id,omitempty"`
}

func GetBuyerWhitSameIP(id string) []models.Buyer {
	var buyers []models.Buyer
	var buyer ResponseBuyer
	var sameIp dgrahpdata.AuxResponseSingle
	var buyerIp dgrahpdata.AuxResponseSingle
	response := dgrahpdata.BdQueryWithVars(query.GETIPSBYBUYERID, id).Json
	error := json.Unmarshal(response, &buyerIp)
	if error != nil {
		log.Fatal(error)
	}
	for j := 0; j < len(buyerIp.Response); j++ {
		ip := buyerIp.Response[j].ResponseString
		response = dgrahpdata.BdQueryWithVars(query.GETBUYERSBYIP, ip).Json
		error := json.Unmarshal(response, &sameIp)
		if error != nil {
			log.Fatal(error)
		}

		for i := 0; i < len(sameIp.Response); i++ {
			id := sameIp.Response[i].ResponseString
			response = dgrahpdata.BdQueryWithVars(query.GETNAMEBYID, id).Json
			error := json.Unmarshal(response, &buyer)
			if error != nil {
				log.Fatal(error)
			}
			for k := 0; k < len(buyer.Response); k++ {
				age, err := strconv.Atoi(buyer.Response[k].Age)
				if err != nil {
					log.Fatal(error)
				}
				temp := models.NewBuyer(buyer.Response[k].Id, buyer.Response[k].Name, age)
				buyers = append(buyers, temp)
			}
		}
	}

	return buyers
}
