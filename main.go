package main

import (
	"encoding/json"
	"fmt"

	"github.com/joho/godotenv"
)

type test struct {
	T []*order
}

func main() {
	godotenv.Load()
	ss := newShipStation()
	oIDs := []string{"111-8045170-2502645", "16467144"}
	pld, err := ss.getOrders(oIDs)
	if err != nil {
		panic(err)
	}
	// test := []order{}
	t := test{T: pld}
	b, _ := json.Marshal(t)

	fmt.Println(string(b))

}
