package main

import (
	"github.com/WedgeNix/CubbyChaser-shared"
	"github.com/mrmiguu/Loading"
	"github.com/mrmiguu/rest"
)

func main() {
	connection, _ := rest.Bool()
	done := load.Ing(`connection()`)
	connection(true)
	done <- true

	_, session := rest.Bytes()

	println(shared.Btos(session()).String())

	// for i := range sess.Cubbies {
	// 	i := i
	// 	_, rCub := rest.Bytes()
	// 	go func() {

	// 		cub, _ := shared.Btoc(rCub())
	// 		sess.Cubbies[i] = cub
	// 	}()
	// 	// go func() {

	// 	// }()
	// }
}
