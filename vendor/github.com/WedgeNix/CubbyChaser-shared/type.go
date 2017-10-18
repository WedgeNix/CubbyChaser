package shared

type Order struct {
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
	ShipTo                   ShipTo
	Items                    []Item
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
	Weight                   Weight
	Dimensions               Dimensions
	InsuranceOptions         interface{}
	InternationalOptions     interface{}
	AdvancedOptions          AdvancedOptions
	TagIDs                   []int
	UserID                   string
	ExternallyFulfilled      bool
	ExternallyFulfilledBy    string
}

type Item struct {
	OrderItemID       int
	LineItemKey       string
	SKU               string
	Name              string
	ImageURL          string
	Weight            Weight
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

type Weight struct {
	WeightUnits int
	Value       float64
	Units       string
}

type Dimensions struct {
	Units  string
	Length float64
	Width  float64
	Height float64
}

type ShipTo struct {
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

type AdvancedOptions struct {
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

type Queue map[int]int

type Session struct {
	ID      int
	Cubbies []Order
}
