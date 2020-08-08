package main

import (
	"pruebaTecnica/pruebatecnicabackend/features/purchase/domain/usecases"
	_ "pruebaTecnica/pruebatecnicabackend/features/purchase/domain/usecases"
	_ "pruebaTecnica/pruebatecnicabackend/features/purchase/repository"
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
	//controllers.EndpointInit()
	//usecases.GetPurchaseByDate("")
	usecases.GetBuyerInfo("187d6d69")
	//fmt.Println(dgrahpdata.BdQuery(q))

}
