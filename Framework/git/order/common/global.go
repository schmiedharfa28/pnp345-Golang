package common

//Struct API
// Order struct (Model) ...

type Message struct {
	Code    int     `json:"code"`
	Remark  string  `json:"remark"`
	OrderID string  `json:"orderID"`
	Orders  *Orders `json:"orders,omitempty"`
	Result  *Result `json:"result,omitempty"`
}

type Orders struct {
	OrderID    string         `json:"orderID"`
	CustomerID string         `json:"customerID"`
	EmployeeID string         `json:"employeeID"`
	OrderDate  string         `json:"orderDate"`
	OrdersDet  []OrdersDetail `json:"ordersDetail"`
}

type Products struct {
	ProductID       int     `json:"ProductID"`
	ProductName     string  `json:"ProductName"`
	SupplierID      int     `json:"SupplierID"`
	CategoryID      int     `json:"CategoryID"`
	QuantityPerUnit string  `json:"QuantityPerUnit"`
	UnitPrice       float64 `json:"UnitPrice"`
	UnitsInStock    int     `json:"UnitsInStock"`
	UnitsOnOrder    int     `json:"UnitsOnOrder"`
	ReorderLevel    int     `json:"ReorderLevel"`
	Discontinued    int     `json:"Discontinued"`
	Description     string  `json:"Description"`
}

type OrdersDetail struct {
	OrderID     string  `json:"orderID"`
	ProductID   string  `json:"ProductID"`
	ProductName string  `json:"ProductName"`
	UnitPrice   float64 `json:"UnitPrice"`
	Quantity    int     `json:"Quantity"`
}

type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}

type Customers struct {
	CustomerID   string `json:"CustomerID"`
	CompanyName  string `json:"CompanyName"`
	ContactName  string `json:"ContactName"`
	ContactTitle string `json:"ContactTitle"`
	Address      string `json:"Address"`
	City         string `json:"City"`
	Country      string `json:"Country"`
	Phone        string `json:"Phone"`
	PostalCode   string `json:"PostalCode"`
}

//End Struct API

type FastPayRequest struct {
	Merchant   string `json:"merchant"`
	MerchantID string `json:"merchant_id"`
	Request    string `json:"request"`
	Signature  string `json:"signature"`
}

type FastPayResponse struct {
	Response       string           `json:"response"`
	Merchant       string           `json:"merchant"`
	MerchantID     string           `json:"merchant_id"`
	PaymentChannel []PaymentChannel `json:"payment_channel"`
	ResponseCode   string           `json:"response_code"`
	ResponseDesc   string           `json:"response_desc"`
}

type PaymentChannel struct {
	PgCode string `json:"pg_code"`
	PgName string `json:"pg_name"`
}

type InquiryPayment struct {
	Response          string `json:"response"`
	BillNo            string `json:"bill_no"`
	Merchant          string `json:"merchant"`
	MerchantID        string `json:"merchant_id"`
	PaymentDate       string `json:"payment_date"`
	PaymentReff       string `json:"payment_reff"`
	PaymentStatusCode string `json:"payment_status_code"`
	PaymentStatusDesc string `json:"payment_status_desc"`
	ResponseCode      string `json:"response_code"`
	ResponseDesc      string `json:"response_desc"`
	TrxID             string `json:"trx_id"`
}

type InquiryRequest struct {
	BillNo     string `json:"bill_no"`
	MerchantID string `json:"merchant_id"`
	Request    string `json:"request"`
	Signature  string `json:"signature"`
	TrxID      string `json:"trx_id"`
}
