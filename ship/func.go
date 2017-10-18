package ship

import (
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/WedgeNix/CubbyChaser-shared"
	"github.com/WedgeNix/util"
)

// New initiates new shipstaion controller.
func New() *Control {
	return &Control{
		shipURL:  `https://ssapi.shipstation.com/`,
		username: util.MustGetenv("SHIP_API_KEY"),
		password: util.MustGetenv("SHIP_API_SECRET"),
		client:   http.Client{},
	}
}

// GetOrders converts order numbers to a list of orders.
func (c Control) GetOrders(nums []string) ([]shared.Order, error) {
	orders := []shared.Order{}

	var err error
	idOnce.Do(func() {
		var pay *payload
		pay, err = c.ssOrders()
		if err != nil {
			return
		}
		for _, ord := range pay.Orders {
			idMap[ord.OrderNumber] = ord
		}
	})
	if err != nil {
		return nil, err
	}

	var atLeastOne bool
	for _, num := range nums {
		ord, exists := idMap[num]
		if !exists {
			continue
		}
		orders = append(orders, ord)
		atLeastOne = true
	}
	if !atLeastOne {
		return orders, errors.New("all orders shipped and/or canceled")
	}
	return orders, nil
}

// PrintLabels creates to-generate files for printing.
func (c Control) PrintLabels(ords []shared.Order) ([][]byte, error) {
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
