package query

const (
	GETIPBYBUYERID = `query all($a: string) {
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
