package query

const (
	GETPRODUCTSIDBYBUYERID = `query all($a: string) {
		Response(func: eq(buyer_id, $a)) {
			ResponseString : product_ids
		}
	  }`
	GETPRODUCT = `query all($a: string) {
		Response(func: eq(id, $a)) {
			ResponseString : name
		}
	}`
)
