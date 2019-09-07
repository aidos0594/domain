package domain

import "time"

type Basket struct {
	Uid                   string       `json:"uid,omitempty"`
	CreatedOn             time.Time    `json:"created_on,omitempty"`
	Scope                 string       `json:"scope,omitempty"`
	UpdatedOn             time.Time    `json:"updated_on,omitempty"`
	OwnType               string       `json:"own_type,omitempty"` //dar account или
	Owner                 string       `json:"owner,omitempty"`
	Comment               string       `json:"comment,omitempty"` //коментарий корзине
	CommonPrice           float64      `json:"common_price,omitempty"`
	CommonOldPrice        float64      `json:"common_old_price,omitempty"`
	CommonQuantity        float64      `json:"common_quantity,omitempty"`
	Items                 []BasketItem `json:"items,omitempty"`
	DeliveryPrice         float64      `json:"delivery_price"`
	EstimatedDeliveryDate *time.Time   `json:"estimated_delivery_date,omitempty"`
	DeliveryCommission    float64      `json:"delivery_commission,omitempty"`
	RefundCommission      float64      `json:"refund_commission,omitempty"`

	RecieverContacts    *Contacts  `json:"reciever_contacts,omitempty"`
	PayTypes            string     `json:"pay_types,omitempty"`
	PayTitle            string     `json:"pay_title,omitempty"`
	DeliveryTypes       string     `json:"delivery_types,omitempty"`
	DeliveryTitle       string     `json:"delivery_title,omitempty"`
	CityId              string     `json:"city_id,omitempty"`
	PickupDate          *time.Time `json:"pickup_date,omitempty"`
	CalculateByDelivery bool       `json:"calculate_by_delivery"`
}

type BasketItem struct {
	Nomenclature
	MerchantInfo     Contacts      `json:"merchant_info,omitempty"`
	ProductVariation string        `json:"product_variation"`
	PickupPoint      string        `json:"pickup_point,omitempty"` //otkuda zaberet item
	PickupAddress    *PickupAddres `json:"pickup_address,omitempty"`
	DeliveryDate     string        `json:"delivery_date,omitempty"`
}

type PickupAddres struct {
	RegionId  string  `json:"regionId,omitempty"`
	Street    string  `json:"street,omitempty"`
	StreetNum string  `json:"streetNum,omitempty"`
	AptNum    string  `json:"aptNum,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	AdditionalInfo string `json:"additional_info,omitempty"`
}
