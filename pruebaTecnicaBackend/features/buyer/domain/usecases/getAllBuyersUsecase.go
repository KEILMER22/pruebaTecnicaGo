package usecases

import (
	dgrahpdata "pruebatecnica/pruebatecnicabackend/core/dgraphdata"
	"pruebatecnica/pruebatecnicabackend/features/buyer/query"

	"github.com/dgraph-io/dgo/protos/api"
)

// GetAllBuyers returns all the data base buyers name in a string array
func GetAllBuyers() *api.Response {
	return dgrahpdata.BdQuery(query.BUYERSQUERY)
}
