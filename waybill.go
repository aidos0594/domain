package domain

import (
	"time"
)

type ElasticWaybill struct {
	NumericId     int       `json:"numeric_id,omitempty"`
	WaybillId     string    `json:"waybill_id,omitempty"`
	StockId       string    `json:"stock_id,omitempty"`
	Scope         string    `json:"scope,omitempty"` // ?
	StockName     string    `json:"stock_name,omitempty"`
	Status        string    `json:"status,omitempty"`
	MerchantID    string    `json:"merchant_id,omitempty"`
	QuantityItems float64   `json:"quantity_items,omitempty"` // ?
	CommonSum     float64   `json:"common_sum,omitempty"`
	CreatedOn     time.Time `json:"created_on,omitempty"`
	UpdatedOn     time.Time `json:"updated_on,omitempty"`
	SkusNumber    []string  `json:"skus_number,omitempty"` // ?
	Comment       string    `json:"comment,omitempty"`
	Supplier      string    `json:"supplier,omitempty"`      // postavchik
	EmployeeId    string    `json:"employee_id,omitempty"`   // sotrudniki
	EmployeeName  string    `json:"employee_name,omitempty"` // sotrudniki
	AccountId     string    `json:"account_id,omitempty"`    // kassa1 or kassa2 or bankId
	WaybillType   string    `json:"waybill_type,omitempty"`  // konsignasia or oplachennaya
	DocumentType  string    `json:"document_type,omitempty"` // type of document
}

type Waybill_SkuId struct {
	DocumentQuantity float64 `json:"document_quantity,omitempty"` // ?
	ReceivedQuantity float64 `json:"quantity"`                    // kolichestvo priniaemogo tovara
	Price            float64 `json:"price,omitempty"`             // sena prodaji
	SkuId            string  `json:"sku_id,omitempty"`            // id tovara s nawei sisteme
	ProductName      string  `json:"product_name,omitempty"`      // imya produkta
	Barcode          string  `json:"barcode, omitempty"`          // shtihcode
	Residue          float64 `json:"residue,omitempty"`           // ostatok kolichestvo
	Category         string  `json:"category,omitempty"`          // kategoria
	Measure          string  `json:"measure,omitempty"`           // edinitsa izmerenie
	WaybillPrice     float64 `json:"waybill_price,omitempty"`     // sena po nakladnoi|priemka
	CostPrice        float64 `json:"cost_price,omitempty"`        // sebestoimost
	Markup           float64 `json:"markup,omitempty"`            // nasenka|priemka
	TotalMoney       float64 `json:"total_money,omitempty"`       // itogovaya summa
	Comment          string  `json:"comment,omitempty"`           // komment|spisanie
	WaybillSkuType   string  `json:"waybill_sku_type,omitempty"`  // type списание|spisanie
}

type Waybill struct {
	DocumentType string `json:"document_type"` // prihodnaya or raskladnaya
	//DocumentId   string          `json:"document_id, omitempty"` // ?
	Skus           []Waybill_SkuId `json:"skus,omitempty"`           // tovary
	NumericId      int             `json:"numeric_id,omitempty"`     // document id dlya klientov
	WaybillId      string          `json:"waybill_id,omitempty"`     // id priemki v nawei sisteme
	CreatedOn      time.Time       `json:"created_on,omitempty"`     // data sozdanie
	UpdatedOn      time.Time       `json:"updated_on,omitempty"`     // data obnovlenie
	StockName      string          `json:"stock_name,omitempty"`     // imia stoka
	StockId        string          `json:"stock_id,omitempty"`       // id stoka v nawei sisteme
	MerchantID     string          `json:"merchant_id,omitempty"`    // id merchanta
	Status         string          `json:"status,omitempty"`         // draft or priemka
	Scope          string          `json:"scope,omitempty"`          // scope
	RecieveMoney   float64         `json:"recieve_money,omitempty"`  // sena priemki
	SellMoney      float64         `json:"sell_money,omitempty"`     // sena prodaziy
	QuantityItems  int64           `json:"quantity_items,omitempty"` // kolichestvo tovarov
	Supplier       string          `json:"supplier,omitempty"`       // postavwik
	EmployeeId     string          `json:"employee_id,omitempty"`    // poluchatel id
	EmployeeName   string          `json:"employee_name,omitempty"`  // imia poluchatela
	AccountId      string          `json:"account_id,omitempty"`     // kassa1 or kassa2 or bankId
	Comment        string          `json:"comment,omitempty"`        // koment
	WaybillType    string          `json:"waybill_type,omitempty"`   // konsignasia or oplachennaya
	Consignment    bool            `json:"consignment"`              // pod konsignasiu
	CommonSum      float64         `json:"common_sum"`
	ReversedTime   time.Time       `json:"reversed_time"`
	DocumentNumber string          `json:"document_number"`
}
