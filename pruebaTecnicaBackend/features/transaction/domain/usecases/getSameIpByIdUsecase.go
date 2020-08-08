package usecases

import (
	"encoding/json"
	"fmt"
	"log"
	dgrahpdata "pruebatecnica/pruebatecnicabackend/core/dgraphdata"
	"pruebatecnica/pruebatecnicabackend/features/transaction/query"
)

func IpBuyers(id string) []string {
	var sameIp dgrahpdata.AuxResponseSingle
	var buyerIp dgrahpdata.AuxResponseSingle
	var buyerNames dgrahpdata.AuxResponseSingle
	var sameIpNames []string
	response := dgrahpdata.BdQueryWithVars(query.GETIPBYBUYERID, id).Json
	error := json.Unmarshal(response, &buyerIp)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(buyerIp)
	for j := 0; j < len(buyerIp.Response); j++ {
		ip := buyerIp.Response[j].ResponseString
		response = dgrahpdata.BdQueryWithVars(query.GETUSERSBYIP, ip).Json
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
			sameIpNames = append(sameIpNames, name)
		}

	}
	sameIpNames = dgrahpdata.RemoveDuplicatesUnordered(sameIpNames)
	return sameIpNames
}
