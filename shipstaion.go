package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/mrmiguu/un"
)

// weight of whole order/item.
type weight struct {
	WeightUnits int
	Value       float64
	Units       string
}

// dimensions demensions of package form shipstaion.
type dimensions struct {
	Units  string
	Length float64
	Width  float64
	Height float64
}

// shipFrom data from where we ship.
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

// shipTo address info of customer.
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
	OrderID             int
	UserID              string
	CustomerEmail       string
	OrderNumber         string
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
	ShipTo              shipTo
	Weight              weight
	Dimensions          dimensions
	InsuranceOptions    interface{}
	AdvancedOptions     interface{}
	ShipmentItems       interface{}
	LabelData           string
	FormData            interface{}
}

// shipLable struct POST to shipstaions api for print.
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
	AdvancedOptions      advancedOptions
	TestLabel            bool
}

// Payload is the first level of a ShipStation HTTP response body.
type payload struct {
	Orders []*order
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
	WarehouseID       interface{}
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

// shipControl controller data for shipstaion calls.
type shipControl struct {
	ShipURL  string
	Username string
	Password string
	Client   http.Client
}

// limits takes requests left and seconds to wait form api header.
type limits struct {
	Requests int
	Seconds  int
}

var (
	lim   = make(chan bool, 1)
	idMap map[string]*order
)

// initLimits initiates a new limit for post call to shipstaion.
func initLimits() {
	lim <- true
}

// newShipStation initiates new shipstaion controler.
func newShipStation() *shipControl {
	client := &http.Client{}
	return &shipControl{
		ShipURL:  `https://ssapi.shipstation.com/`,
		Username: os.Getenv("SHIP_API_KEY"),
		Password: os.Getenv("SHIP_API_SECRET"),
		Client:   *client,
	}
}

func (s shipControl) getOrders(ordIDs []string) ([]*order, error) {
	orders := []*order{}
	pay, err := s.ssOrders()
	if err != nil {
		return nil, err
	}
	if idMap == nil {
		println("ok")
		idMap = map[string]*order{}
		for _, ord := range pay.Orders {
			idMap[ord.OrderNumber] = ord
		}
	}
	for _, id := range ordIDs {
		orders = append(orders, idMap[id])
	}
	return orders, nil
}

// getOrders gets awaiting_shipment.
func (s shipControl) ssOrders() (*payload, error) {
	pg := 1
	pay := &payload{}
	head, err := getPage(pg, pay, s)
	if err != nil {
		return pay, err
	}
	err = setLimits(head)
	if err != nil {
		return pay, err
	}
	for pay.Page < pay.Pages {
		pg++
		ords := pay.Orders
		pay = &payload{}
		println("limit before")
		<-lim
		println("limit after")
		head, err = getPage(pg, pay, s)
		if err != nil {
			return pay, err
		}
		err := setLimits(head)
		if err != nil {
			return pay, err
		}
		pay.Orders = append(ords, pay.Orders...)

	}
	return pay, nil
}

// getPage gets the extrea pages of getOrders.
func getPage(page int, pay *payload, s shipControl) (http.Header, error) {

	// make query.
	query := url.Values(map[string][]string{})
	query.Set(`page`, strconv.Itoa(page))
	query.Set(`orderStatus`, `awaiting_shipment`)
	query.Set(`pageSize`, `500`)

	// make request.
	req, err := http.NewRequest("GET", s.ShipURL+`orders?`+query.Encode(), nil)
	if err != nil {
		return nil, err
	}

	// set headers and authentication.
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(s.Username, s.Password)

	// calling shipstaion.
	resp, err := s.Client.Do(req)
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

	if reqs < 1 {
		println("waiting")
		time.Sleep(time.Duration(secs) * time.Second)
	}
	lim <- true
	return nil

}

func (s shipControl) printLabels(ords []*order) ([][]byte, error) {
	// initiate limit.
	initLimits()
	pd := make([][]byte, len(ords))

	// range over orders for printing.
	for i, ord := range ords {
		lr := &labelRes{}
		<-lim

		head, err := s.printLabel(ord, lr)
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

func (s shipControl) printLabel(ord *order, lr *labelRes) (http.Header, error) {

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
	req.SetBasicAuth(s.Username, s.Password)

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	un.Wrap(json.NewDecoder(resp.Body).Decode(&lr))

	return resp.Header, nil
}

func (s shipControl) postCustomFields(ords []*order, sID string) error {
	for _, ord := range ords {
		// adding session ID to order data
		ord.AdvancedOptions.CustomField3 = sID
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
	req.SetBasicAuth(s.Username, s.Password)

	resp, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
