package ship

import "github.com/WedgeNix/CubbyChaser-shared"

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
	ShipTo              shared.ShipTo
	Weight              shared.Weight
	Dimensions          shared.Dimensions
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
	Weight               shared.Weight
	Dimensions           shared.Dimensions
	ShipFrom             shipFrom
	ShipTo               shared.ShipTo
	InsuranceOptions     interface{}
	InternationalOptions interface{}
	AdvancedOptions      shared.AdvancedOptions
	TestLabel            bool
}

// payload is the first level of a ShipStation HTTP response body.
type payload struct {
	Orders []shared.Order
	Total  int
	Page   int
	Pages  int
}

// limits takes requests left and seconds to wait form api header.
type limits struct {
	Requests int
	Seconds  int
}
