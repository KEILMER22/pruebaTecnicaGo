package query

const (
	BUYERSQUERY = `{
		Response(func: has(age)){
			id
			name
			age
		}
	}`
	GETIPSBYBUYERID = `query all($a: string) {
		Response(func: eq(buyer_id, $a)) {
			ResponseString : ip
		}
	}`
	GETBUYERSBYIP = `query all($a: string) {
		Response(func: eq(ip, $a)) {
			ResponseString : buyer_id
		}
	}`

	GETNAMEBYID = `query all($a: string) {
		Response(func: eq(id, $a)) {
			Name : name
			Age : age
			Id : id
		}
	}`

	GETPRODUCTBYID = `query all($a: string) {
		Response(func: has(price, $a)) {
			Id : id
			Name : name
			Price : price
		}
	}`
)
