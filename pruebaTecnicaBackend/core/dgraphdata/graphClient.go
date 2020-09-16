package dgrahpdata

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

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
)

type AuxResponse struct {
	Response []AuxResponseString `json:"Response,omitempty"`
}
type AuxResponseString struct {
	ResponseString []string `json:"ResponseString,omitempty"`
}
type AuxResponseSingle struct {
	Response []AuxResponseSingleString `json:"Response,omitempty"`
}
type AuxResponseSingleString struct {
	ResponseString string `json:"ResponseString,omitempty"`
}

func RemoveDuplicatesUnordered(elements []string) []string {
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
