package main

import (
	"fmt"

	"github.com/WedgeNix/CubbyChaser-shared"
	load "github.com/mrmiguu/Loading"
	"github.com/mrmiguu/rest"
)

func init() {
	rest.Connect("/")
}

func main() {
	done := load.New("opening session queue channel")
	_, queue := rest.New(shared.SessionQueueH).Bytes()
	done <- true

	done = load.New("reading session queue")
	q := shared.Btoq(queue())
	done <- true

	fmt.Println(q)
}
