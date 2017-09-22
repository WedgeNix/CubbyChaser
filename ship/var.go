package ship

var (
	lim   = make(chan bool, 1)
	idMap map[string]Order
)
