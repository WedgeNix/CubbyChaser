package ship

import (
	"sync"

	"github.com/WedgeNix/CubbyChaser-shared"
)

var (
	lim    = make(chan bool, 1)
	idOnce sync.Once
	idMap  = map[string]shared.Order{}
)
