package query

const (
	BUYERSQUERY = `{
		people(func: has(age)){
			name
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
			ResponseString : name
		}
	}`
)
