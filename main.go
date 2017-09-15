package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	ss := newShipStation()
	pld, err := ss.getOrders()
	if err != nil {
		panic(err)
	}
	fmt.Println(pld)

}
