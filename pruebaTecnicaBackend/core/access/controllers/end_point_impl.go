package controllers

import (
	"fmt"
	"net/http"
	PurchaseUsecase "pruebaTecnica/pruebaTecnicaBackend/features/purchase/domain/usecases"
	BuyerUsecase "pruebatecnica/pruebatecnicabackend/features/buyer/domain/usecases"

	"github.com/go-chi/chi"
)

//PORT -> serve and listen port
var PORT string = ":8090"

func EndpointInit() {

	router := chi.NewRouter()
	//post
	router.Route("/load_data", func(r chi.Router) {
		r.Get("/{date}", loadData)
		r.Get("/", loadData)
	})
	router.Get("/list_buyers", func(w http.ResponseWriter, r *http.Request) {
		listBuyers(w, r)
	})
	router.Route("/buyer_info", func(r chi.Router) {
		r.Get("/{id}", buyerInfo)
	})
	err := http.ListenAndServe(PORT, router)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}

}

func loadData(w http.ResponseWriter, r *http.Request) {
	date := chi.URLParam(r, "date")
	if date != "" {
		w.Write(PurchaseUsecase.GetPurchaseByDate(date))
	} else {
		w.Write(PurchaseUsecase.GetPurchaseByDate(""))
	}
}

func listBuyers(w http.ResponseWriter, r *http.Request) {

	w.Write(BuyerUsecase.GetAllBuyers().Json)

}

func buyerInfo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	buyerInfo := BuyerUsecase.GetBuyerInfo(id)
	w.Write(buyerInfo)
}
