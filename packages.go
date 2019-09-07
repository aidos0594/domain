package domain

import (
	"time"
)

type Package struct {
	Uid             string         `json:"uid,omitempty"` //номер заказа
	Scope           string         `json:"scope,omitempty"`
	Status          string         `json:"status,omitempty"`
	StatusTitle     string         `json:"status_title,omitempty"` //elastic
	PackageStatus   []Statuses     `json:"package_status,omitempty"`
	CourierCache    string         `json:"courier_cache,omitempty"` //
	ResiverContacts *Contacts      `json:"resiver_contacts,omitempty"`
	DeliveryPrice   float64        `json:"delivery_price,omitempty"`
	TrackNumber     string         `json:"track_number,omitempty"`
	CreatedOn       time.Time      `json:"created_on,omitempty"`
	UpdatedOn       time.Time      `json:"created_on,omitempty"`
	DeliveryType    string         `json:"delivery_type,omitempty"`
	DeliveryInfo    string         `json:"delivery_info,omitempty"`
	Products        []PackageItems `json:"products"`
	CommonPrice     float64        `json:"common_price,omitempty"`
	CommonOldPrice  float64        `json:"common_old_price,omitempty"`
	CommonQuantity  int            `json:"common_quantity,omitempty"`
}

type PackageItems struct {
	Product          Product  `json:"product,omitempty"` ///поиск по названию товара в заказе
	SenderContact    Contacts `json:"sender_contact,omitempty"`
	ProductVariation string   `json:"product_variation,omitempty"`
	Comment          string   `json:"comment,omitempty"`
	BundleId         string   `json:"bundle_id,omitempty"`
	BundleTitle      string   `json:"bundle_title,omitempty"`
	BundleQuantity   float64  `json:"bundle_quantity,omitempty"`
	Quantity         float64  `json:"quantity,omitempty"`
	//StockAddress
}

type Contacts struct {
	Uid          string `json:"uid,omitempty"`
	Name         string `json:"name,omitempty"`
	Surname      string `json:"surname,omitempty"`
	Title        string `json:"title,omitempty"` // поиск по названию
	City         string `json:"city,omitempty"`
	Address      string `json:"address,omitempty"`
	Email        string `json:"email,omitempty"`
	HouseNumber  string `json:"house_number,omitempty"`
	FlatNumber   string `json:"flat_number,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	WalletMobile string `json:"wallet_mobile,omitempty"`
	LogoToken    string `json:"logoToken"`
	Iin          string `json:"iin,omitempty"`
}

type DeliveryPriceItems struct {
	StartPrice    float64 `json:"startPrice,omitempty"`
	EndPrice      float64 `json:"endPrice,omitempty"`
	DeliveryPrice float64 `json:"deliveryPrice,omitempty"`
}

type MerchantsIfo struct {
	Contacts
	RefundCommission []float64             `json:"orderRefusalFee,omitempty"`
	Deliveries       *[]DeliveryPriceItems `json:"deliveries,omitempty"`
	MinPrice         float64               `json:"min_price,omitempty"`
}
