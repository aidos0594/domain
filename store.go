package domain

import (
	"time"
)

type StockEvent struct {
	StockId   string
	Tuid      string
	EventType string `json:"event_type, omitempty"`
	Body      Stock
}

type StockContact struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Email       string `json:"email,omitempty"`
}

type UserInfo struct {
	UserId    string       `json:"userId,omitempty"`
	FirstName string       `json:"firstname,omitempty"`
	LastName  string       `json:"lastname,omitempty"`
	JobTitle  string       `json:"jobTitle,omitempty"`
	Contacts  StockContact `json:"contacts,omitempty"`
}

type Stock struct {
	ID          string         `json:"id,omitempty" cql:"id"`
	MerchantId  *string        `json:"merchantId,omitempty" cql:"merchantId"`
	Type        WarehouseType  `json:"type,omitempty" cql:"type"`
	Title       *string        `json:"title,omitempty" cql:"title"`
	Owner       *string        `json:"owner,omitempty"`
	MOL         *MOLInfo       `json:"stock_contacts,omitempty"`
	Scope       *string        `json:"scope,omitempty"`
	Name        *string        `json:"name,omitempty"`
	Location    *Location      `json:"location,omitempty" cql:"location"`
	ZipCode     *string        `json:"zip_code,omitempty"`
	CreatedOn   time.Time      `json:"created_on,omitempty"`
	UpdatedOn   time.Time      `json:"updated_on,omitempty"`
	UpdatedTuid string         `json:"updated_tuid,omitempty"`
	Works       *[]WorkingTime `json:"works,omitempty"`
	Users       *[]UserInfo    `json:"users,omitempty"`
	IsVisible   *bool          `json:"isVisible,omitempty"`
}

type MOLInfo struct {
	FIO      string       `json:"fio,omitempty"`
	Contacts StockContact `json:"contacts,omitempty"`
}

type Location struct {
	Country  string    `json:"country,omitempty" cql:"country"`
	City     string    `json:"city,omitempty" cql:"city"`
	Address  string    `json:"address,omitempty" cql:"address"`
	LngLat   []float64 `json:"lnglat,omitempty" cql:"lnglat"`
	GeoHash  string    `json:"geohash,omitempty" cql:"geohash"`
	RegionId string    `json:"regionId,omitempty" cql:"regionId"`
}

type WorkingTime struct {
	Day    time.Weekday `json:"day,omitempty"`
	Open   int64        `json:"open,omitempty"`
	Close  int64        `json:"close,omitempty"`
	IsOpen bool         `json:"isOpen,omitempty"`
}

type HoldEvent struct {
	Sku        string    `json:"sku, omitempty"`
	StockId    string    `json:"stockId, omitempty"`
	TimeToHold time.Time `json:"timeToHold, omitempty"`
	Amount     float64   `json:"amount, omitempty"`
	FromId     string    `json:"fromId, omitempty"`
	HoldType   string    `json:"holdType, omitempty"`
	HoldTuid   string    `json:"holdTuid, omitempty"`
	WaybillId  string    `json:"waybillId, omitempty"`
}

func ConstructUniqueID(stockid, skuid string) string {
	return stockid + "$" + skuid
}

type StockTransfers struct {
	Sku                   string    `json:"sku,omitempty"`
	StockId               string    `json:"stockId,omitempty"`
	StockName             string    `json:"stockName,omitempty"`
	Price                 float64  `json:"price"`
	To                    string    `json:"to,omitempty"`
	Document              string    `json:"document,omitempty"`
	DocumentId            string    `json:"documentId,omitempty"`
	Comment               string    `json:"comment,omitempty"`
	Barcode               string    `json:"barcode,omitempty"`
	CreatedOn             time.Time `json:"created_on,omitempty"`
	Amount                float64  `json:"amount"`
	PartNumber            string    `json:"partNumber,omitempty"`
	StockPosition         string    `json:"stockPosition,omitempty"`
	UpdatedTuid           string    `json:"updatedTuid,omitempty"` //don't need for simple transaction
	ProductId             string    `json:"productId,omitempty"`
	SkuName               string    `json:"skuName,omitempty"`
	Reason                string    `json:"reason,omitempty"`
	InventoryId           string    `json:"inventory_id,omitempty"`
	Version               int       `json:"version,omitempty"`
	Actual                string    `json:"actual,omitempty"`
	Residue               string    `json:"residue,omitempty"`
	CalculatedTransaction bool      `json:"calculated_transaction"`
	LastAmount            float64   `json:"last_amount"`
	LastPrice             float64   `json:"last_price"`
}
