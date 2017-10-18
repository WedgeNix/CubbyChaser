package session

import (
	"sync"

	"github.com/WedgeNix/CubbyChaser-shared"
	"github.com/WedgeNix/CubbyChaser/ship"
	"github.com/WedgeNix/warn"
	load "github.com/mrmiguu/Loading"
)

var (
	set sync.Once
	ss  *ship.Control
)

func New(id int, ordNums []string) (shared.Session, error) {
	set.Do(func() { ss = ship.New() })

	s := shared.Session{ID: id}

	done := load.New("getting orders from ShipStation")
	ords, err := ss.GetOrders(ordNums)
	if err != nil {
		return s, err
	}
	s.Cubbies = ords

	if testOrder {
		warn.Do("using test order")
		s.Cubbies = []shared.Order{{
			OrderNumber: "ABC123",
			Items: []shared.Item{{
				Quantity: 1,
				UPC:      "918273645647",
			}, {
				Quantity: 3,
				UPC:      "008273645600",
			}},
		}}
	}

	done <- true

	return s, nil
}
