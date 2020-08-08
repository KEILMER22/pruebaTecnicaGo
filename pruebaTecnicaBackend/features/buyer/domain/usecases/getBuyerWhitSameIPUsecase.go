package usecases

import (
	"encoding/json"
	"log"
	dgrahpdata "pruebatecnica/pruebatecnicabackend/core/dgraphdata"
	"pruebatecnica/pruebatecnicabackend/features/buyer/query"
)

//GetBuyerInfo return
func GetBuyerWhitSameIP(id string) []string {
	var sameIp dgrahpdata.AuxResponseSingle
	var buyerIp dgrahpdata.AuxResponseSingle
	var buyerNames dgrahpdata.AuxResponseSingle
	var buyers []string
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
			error := json.Unmarshal(response, &buyerNames)
			if error != nil {
				log.Fatal(error)
			}
			name := buyerNames.Response[0].ResponseString
			buyers = append(buyers, name)
		}

	}
	buyers = dgrahpdata.RemoveDuplicatesUnordered(buyers)
	return buyers
}
