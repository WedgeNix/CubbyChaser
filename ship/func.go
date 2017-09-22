package ship

import (
	"encoding/base64"
	"net/http"
	"os"
)

// New initiates new shipstaion controller.
func New() *Control {
	client := &http.Client{}
	return &Control{
		ShipURL:  `https://ssapi.shipstation.com/`,
		Username: os.Getenv("SHIP_API_KEY"),
		Password: os.Getenv("SHIP_API_SECRET"),
		Client:   *client,
	}
}

// GetOrders converts order IDs to a list of orders.
func (c Control) GetOrders(ids []string) ([]Order, error) {
	orders := []Order{}
	pay, err := c.ssOrders()
	if err != nil {
		return nil, err
	}
	if idMap == nil {
		println("ok")
		idMap = map[string]Order{}
		for _, ord := range pay.Orders {
			idMap[ord.OrderNumber] = ord
		}
	}
	for _, id := range ids {
		orders = append(orders, idMap[id])
	}
	return orders, nil
}

// PrintLabels creates to-generate files for printing.
func (c Control) PrintLabels(ords []Order) ([][]byte, error) {
	// initiate limit.
	initLimits()
	pd := make([][]byte, len(ords))

	// range over orders for printing.
	for i, ord := range ords {
		<-lim
		var lr labelRes
		head, err := c.printLabel(ord, &lr)
		if err != nil {
			return nil, err
		}
		err = setLimits(head)
		if err != nil {
			return nil, err
		}
		dec, err := base64.StdEncoding.DecodeString(lr.LabelData)
		if err != nil {
			return nil, err
		}
		pd[i] = dec
	}
	return pd, nil
}
