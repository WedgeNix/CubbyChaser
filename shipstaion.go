package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type weight struct {
	WeightUnits int
	Value       float64
	Units       string
}
type dimensions struct {
	Units  string
	Length float64
	Width  float64
	Height float64
}
type shipFrom struct {
	Name        string
	Company     string
	Street1     string
	Street2     string
	Street3     interface{}
	City        string
	State       string
	PostalCode  string
	Country     string
	Phone       interface{}
	Residential bool
}
type shipTo struct {
	Name        string
	Company     string
	Street1     string
	Street2     string
	Street3     interface{}
	City        string
	State       string
	PostalCode  string
	Country     string
	Phone       interface{}
	Residential bool
}

type labelRes struct {
	ShipmentID          int
	OrderID             interface{}
	UserID              interface{}
	CustomerEmail       interface{}
	OrderNumber         interface{}
	CreateDate          string
	ShipDate            string
	ShipmentCost        float64
	InsuranceCost       float64
	TrackingNumber      string
	IsReturnLabel       bool
	BatchNumber         interface{}
	CarrierCode         string
	ServiceCode         string
	PackageCode         string
	Confirmation        string
	WarehouseID         interface{}
	Voided              bool
	VoidDate            interface{}
	MarketplaceNotified bool
	NotifyErrorMessage  interface{}
	ShipTo              interface{}
	Weight              weight
	Dimensions          dimensions
	InsuranceOptions    interface{}
	AdvancedOptions     interface{}
	ShipmentItems       interface{}
	LabelData           string
	FormData            interface{}
}

type shipLable struct {
	CarrierCode          string
	ServiceCode          string
	PackageCode          string
	Confirmation         string
	ShipDate             string
	Weight               weight
	Dimensions           dimensions
	ShipFrom             shipFrom
	ShipTo               shipTo
	InsuranceOptions     interface{}
	InternationalOptions interface{}
	AdvancedOptions      interface{}
	TestLabel            bool
}

// Payload is the first level of a ShipStation HTTP response body.
type payload struct {
	Orders []order
	Total  int
	Page   int
	Pages  int
}

// Item is the third level of a ShipStation HTTP response body.
type item struct {
	OrderItemID       int
	LineItemKey       string
	SKU               string
	Name              string
	ImageURL          string
	Weight            weight
	Quantity          int
	UnitPrice         float32
	TaxAmount         float32
	ShippingAmount    float32
	WarehouseLocation string
	Options           interface{}
	ProductID         int
	FulfillmentSKU    string
	Adjustment        bool
	UPC               string
	CreateDate        string
	ModifyDate        string
}

// Order is the second level of a ShipStation HTTP response body.
type order struct {
	OrderID                  int
	OrderNumber              string
	OrderKey                 string
	OrderDate                string
	CreateDate               string
	ModifyDate               string
	PaymentDate              string
	ShipByDate               string
	OrderStatus              string
	CustomerID               int
	CustomerUsername         string
	CustomerEmail            string
	BillTo                   interface{}
	ShipTo                   shipTo
	Items                    []item
	OrderTotal               float32
	AmountPaid               float32
	TaxAmount                float32
	ShippingAmount           float32
	CustomerNotes            string
	InternalNotes            string
	Gift                     bool
	GiftMessage              string
	PaymentMethod            string
	RequestedShippingService string
	CarrierCode              string
	ServiceCode              string
	PackageCode              string
	Confirmation             string
	ShipDate                 string
	HoldUntilDate            string
	Weight                   weight
	Dimensions               dimensions
	InsuranceOptions         interface{}
	InternationalOptions     interface{}
	AdvancedOptions          advancedOptions
	TagIDs                   []int
	UserID                   string
	ExternallyFulfilled      bool
	ExternallyFulfilledBy    string
}

// AdvancedOptions holds the "needed" custom fields for post-email tagging.
type advancedOptions struct {
	WarehouseID       int
	NonMachinable     bool
	SaturdayDelivery  bool
	ContainsAlcohol   bool
	MergedOrSplit     bool
	MergedIDs         interface{}
	ParentID          interface{}
	StoreID           int
	CustomField1      string
	CustomField2      string
	CustomField3      string
	Source            string
	BillToParty       interface{}
	BillToAccount     interface{}
	BillToPostalCode  interface{}
	BillToCountryCode interface{}
}

type shipControl struct {
	ShipURL  string
	Username string
	Password string
}

type limits struct {
	Requests int
	Seconds  int
}

var (
	lim = make(chan limits)
)

func initLimits() {
	lim <- limits{
		Requests: 1,
		Seconds:  0,
	}
}

func newShipStation() *shipControl {
	return &shipControl{
		ShipURL:  `https://ssapi.shipstation.com/`,
		Username: os.Getenv("SHIP_API_KEY"),
		Password: os.Getenv("SHIP_API_SECRET"),
	}
}

func (s shipControl) getOrders() (*payload, error) {
	pg := 1
	pay := &payload{}
	head, err := getPage(pg, pay, s)
	if err != nil {
		return pay, err
	}
	limit, err := setLimits(head)
	if err != nil {
		return pay, err
	}
	lim <- limit
	for pay.Page < pay.Pages {
		pg++
		ords := pay.Orders
		pay = &payload{}
		limit := <-lim
		if limit.Requests < 1 {
			time.Sleep(time.Duration(limit.Seconds) * time.Second)
		}
		println("sending pg:", pg)
		head, err = getPage(pg, pay, s)
		if err != nil {
			return pay, err
		}
		println("Done with pg:", pg)
		limit, err := setLimits(head)
		if err != nil {
			return pay, err
		}
		lim <- limit
		pay.Orders = append(ords, pay.Orders...)

	}
	return pay, nil
}

func getPage(page int, pay *payload, s shipControl) (http.Header, error) {
	client := http.Client{}

	// make query
	query := url.Values(map[string][]string{})
	query.Set(`page`, strconv.Itoa(page))
	query.Set(`orderStatus`, `awaiting_shipment`)
	query.Set(`pageSize`, `500`)

	// make request
	req, err := http.NewRequest("GET", s.ShipURL+`orders?`+query.Encode(), nil)
	if err != nil {
		return nil, err
	}

	// set headers and authentication
	req.Header.Add("Content-Type", "application/json")
	println(s.Username, s.Password)
	req.SetBasicAuth(s.Username, s.Password)

	// calling shipstaion
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	println(string(b))
	panic("panicking")
	// read response
	err = json.NewDecoder(resp.Body).Decode(pay)
	if err != nil {
		return nil, err
	}

	return resp.Header, nil
}

func setLimits(head http.Header) (lim limits, err error) {
	// throttle next call by header
	remaining := head.Get("X-Rate-Limit-Remaining")

	var reqs, secs int
	reqs, err = strconv.Atoi(remaining)
	if err != nil {
		return
	}
	reset := head.Get("X-Rate-Limit-Reset")
	secs, err = strconv.Atoi(reset)
	if err != nil {
		return
	}
	lim.Requests = reqs
	lim.Seconds = secs
	return

}
