package main

import (
	"pruebaTecnica/pruebaTecnicaBackend/core/access/controllers"
)

func main() {
	/*
		p := models.NewBuyer("1245", "Gerardo", "22")
		pb, err := json.Marshal(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(p)
		fmt.Println(pb)
	*/
	controllers.EndpointInit()
	//usecases.GetPurchaseByDate("")
	//usecases.GetBuyerInfo("187d6d69")
	//fmt.Println(dgrahpdata.BdQuery(q))

}
