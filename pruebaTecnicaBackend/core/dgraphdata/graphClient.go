package dgrahpdata

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

/*
type Purchases struct {
	Uid          string        `json:"uid,omitempty"`
	Buyers       []Buyer       `json:"Buyers,omitempty"`
	Products     []Product     `json:"Products,omitempty"`
	Transactions []Transaction `json:"Transactions,omitempty"`
}


*/
const (
	RUTA_BASE_DATOS = "localhost:9080"
	SCHEMA          = `
						id: string@index(term) .
						name: string .
						age: string .
						price: string .
						ip: string@index(term) .
						buyer_id: string@index(term) .
						device: string .
						productIds: string .
						cod: string@index(term).
						Buyer: [uid] .
						Product: [uid] .
						Transaction: [uid] .

						type Buyer{  
							id:   string 
							name: string 
							age:  int    
						}
						type Product{
							
							id:    string 
							name:  string 
							price: string 
						}
						type Transaction  {
							id:     string
							buyer_id:    string   
							ip:         string   
							device:     string   
							productIds: [string] 
						}
						type Purchases  {
							cod: string
							Buyer:       [Buyer]       
							Product:     [Product]   
							Transaction:  [Transaction]
						}
						`
	BUYERSQUERY = `{
						people(func: has(name)){
							name
						}
					}`

	GETPRODUCTSID = `query all($a: string) {
						Response(func: eq(buyer_id, $a)) {
							ResponseString : product_ids
						}
					  }`
	GETPRODUCT = `query all($a: string) {
						Response(func: eq(id, $a)) {
							ResponseString : name
						}
					}`

	GETIP = `query all($a: string) {
						Response(func: eq(buyer_id, $a)) {
							ResponseString : ip
						}
					}`
	GETUSERSBYIP = `query all($a: string) {
						Response(func: eq(ip, $a)) {
							ResponseString : buyer_id
						}
					}`

	GETNAMEBYID = `query all($a: string) {
						Response(func: eq(id, $a)) {
							ResponseString : name
						}
					}`
)

// Querys
func BuyersQuery() *api.Response {
	return BdQuery(BUYERSQUERY)

}

type Response struct {
	Response []ResponseString `json:"Response,omitempty"`
}
type ResponseString struct {
	ResponseString []string `json:"ResponseString,omitempty"`
}
type ResponseB struct {
	Response []ResponseStringB `json:"Response,omitempty"`
}
type ResponseStringB struct {
	ResponseString string `json:"ResponseString,omitempty"`
}

func BuyHistory(id string) []string {
	var data Response
	var product ResponseB
	var StringProduct []string
	response := BdQueryWithVars(GETPRODUCTSID, id).Json
	error := json.Unmarshal(response, &data)
	if error != nil {
		log.Fatal(error)
	}
	if len(data.Response) != 0 {
		for j := 0; j < len(data.Response); j++ {
			productIds := data.Response[j].ResponseString
			for i := 0; i < len(productIds); i++ {
				response = BdQueryWithVars(GETPRODUCT, productIds[i]).Json
				error := json.Unmarshal(response, &product)
				if error != nil {
					log.Fatal(error)
				}
				if len(product.Response) != 0 {
					StringProduct = append(StringProduct, product.Response[0].ResponseString)
				}
			}
		}
	}
	StringProduct = removeDuplicatesUnordered(StringProduct)
	return StringProduct
}

func IpBuyers(id string) []string {
	var sameIp ResponseB
	var buyerIp ResponseB
	var buyerNames ResponseB
	var buyers []string
	response := BdQueryWithVars(GETIP, id).Json
	error := json.Unmarshal(response, &buyerIp)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(buyerIp)
	for j := 0; j < len(buyerIp.Response); j++ {
		ip := buyerIp.Response[j].ResponseString
		response = BdQueryWithVars(GETUSERSBYIP, ip).Json
		error := json.Unmarshal(response, &sameIp)
		if error != nil {
			log.Fatal(error)
		}

		for i := 0; i < len(sameIp.Response); i++ {
			id := sameIp.Response[i].ResponseString
			response = BdQueryWithVars(GETNAMEBYID, id).Json
			error := json.Unmarshal(response, &buyerNames)
			if error != nil {
				log.Fatal(error)
			}
			name := buyerNames.Response[0].ResponseString
			buyers = append(buyers, name)
		}

	}
	buyers = removeDuplicatesUnordered(buyers)
	return buyers
}

func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

//-- End Querys
func GetClient() (*dgo.Dgraph, *grpc.ClientConn) {
	conn, err := grpc.Dial(RUTA_BASE_DATOS, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	return dgraphClient, conn
}

func BdQueryWithVars(query string, vars string) *api.Response {
	dgraphClient, conn := GetClient()
	defer conn.Close()
	ctx := context.Background()
	txn := dgraphClient.NewTxn()
	res, err := txn.QueryWithVars(ctx, query, map[string]string{"$a": vars})
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func BdQuery(query string) *api.Response {
	dgraphClient, conn := GetClient()
	defer conn.Close()
	ctx := context.Background()
	txn := dgraphClient.NewTxn()

	res, err := txn.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
func dropData() {
	dgraphClient, conn := GetClient()
	defer conn.Close()
	op := &api.Operation{}
	op.DropAll = true
	ctx := context.Background()
	err := dgraphClient.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}
}
func Schema(schema string) {
	dgraphClient, conn := GetClient()
	defer conn.Close()
	op := &api.Operation{}
	op.Schema = schema
	ctx := context.Background()
	err := dgraphClient.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}
}

func UploadData(pb []byte) {
	dgraphClient, conn := GetClient()
	defer conn.Close()
	dropData()
	Schema(SCHEMA)
	ctx := context.Background()
	mu := &api.Mutation{
		CommitNow: true,
	}
	mu.SetJson = pb
	_, err := dgraphClient.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}
}

func Temp() {

}
