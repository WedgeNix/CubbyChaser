package session

import (
	"github.com/WedgeNix/CubbyChaser-shared"
	"github.com/mrmiguu/rest"
)

type Queue struct {
	Add    chan<- shared.Session
	Delete chan<- shared.Session

	s map[int]shared.Session
	h map[int]*rest.Handler
}
