package ship

import "github.com/WedgeNix/CubbyChaser-shared"

var (
	lim   = make(chan bool, 1)
	idMap map[string]shared.Order
)
