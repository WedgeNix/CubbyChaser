package ship

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/WedgeNix/CubbyChaser-shared"
	"github.com/mrmiguu/un"
)

// initLimits initiates a new limit for post call to shipstaion.
func initLimits() {
	lim <- true
}

// ssOrders gets awaiting_shipment.
func (c Control) ssOrders() (*payload, error) {
	// done := load.New("ssOrders")
	// defer func() { done <- true }()

	pg := 1
	pay := &payload{}
	head, err := c.getPage(pg, pay)
	if err != nil {
		return pay, err
	}
	// done <- false

	err = setLimits(head)
	if err != nil {
		return pay, err
	}
	// done <- false

	for pay.Page < pay.Pages {
		pg++
		ords := pay.Orders
		pay = &payload{}
		// println("limit before")
		<-lim
		// println("limit after")
		head, err = c.getPage(pg, pay)
		if err != nil {
			return pay, err
		}
		err := setLimits(head)
		if err != nil {
			return pay, err
		}
		pay.Orders = append(ords, pay.Orders...)

	}
	<-lim
	return pay, nil
}

// getPage gets the extrea pages of getOrders.
func (c Control) getPage(page int, pay *payload) (http.Header, error) {

	// make query.
	query := url.Values(map[string][]string{})
	query.Set(`page`, strconv.Itoa(page))
	if expiredSessions {
		// warn.Do("using expired sessions")
		query.Set(`createDateStart`, time.Now().Add(-672*time.Hour).Format(`2006-01-02`)+` 00:00:00`)
	} else {
		query.Set(`orderStatus`, `awaiting_shipment`)
	}
	query.Set(`pageSize`, `500`)

	// make request.
	req, err := http.NewRequest("GET", c.shipURL+`orders?`+query.Encode(), nil)
	if err != nil {
		return nil, err
	}

	// set headers and authentication.
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.username, c.password)

	// calling shipstaion.
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read response.
	err = json.NewDecoder(resp.Body).Decode(pay)
	if err != nil {
		return nil, err
	}

	return resp.Header, nil
}

// setLimits limts api calls based on http.Header.
func setLimits(head http.Header) error {
	// done := load.New("setLimits")
	// defer func() { done <- true }()

	// throttle next call by header.
	remaining := head.Get("X-Rate-Limit-Remaining")

	reqs, err := strconv.Atoi(remaining)
	if err != nil {
		return err
	}
	reset := head.Get("X-Rate-Limit-Reset")
	secs, err := strconv.Atoi(reset)
	if err != nil {
		return err
	}
	// done <- false

	if reqs < 1 {
		println("waiting")
		time.Sleep(time.Duration(secs) * time.Second)
	}
	// done <- false

	lim <- true
	// done <- false

	return nil

}

func (c Control) printLabel(ord shared.Order, lr *labelRes) (http.Header, error) {

	shipLab := &shipLable{
		CarrierCode:  ord.CarrierCode,
		ServiceCode:  ord.ServiceCode,
		PackageCode:  ord.PackageCode,
		Confirmation: ord.Confirmation,
		ShipDate:     time.Now().Format("2006-01-02"),
		Weight:       ord.Weight,
		Dimensions:   ord.Dimensions,
		ShipFrom: shipFrom{
			Name:       "Shipping Dept.",
			Company:    "WedgeNix",
			Street1:    "1538 Howard Access RD",
			Street2:    "Unit A",
			City:       "Upland",
			State:      "CA",
			PostalCode: "91786",
		},
		ShipTo:    ord.ShipTo,
		TestLabel: true,
	}
	body, err := json.Marshal(shipLab)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://ssapi.shipstation.com/shipments/createlabel", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.username, c.password)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	un.Wrap(json.NewDecoder(resp.Body).Decode(lr))

	return resp.Header, nil
}

func (c Control) postCustomFields(ords []shared.Order, sID string) error {
	for i := range ords {
		// adding session ID to Order data
		ords[i].AdvancedOptions.CustomField3 = sID
	}

	body, err := json.Marshal(ords)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://ssapi.shipstation.com/", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.username, c.password)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
