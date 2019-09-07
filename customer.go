package domain

import (
	"time"
)


type Customer struct {
	Scope [] string `json:"scope,omitempty"`
	Uid string `json:"uid,omitempty"`
	CustomerCategoryId []string `json:"customer_category_id,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Address string `json:"address,omitempty"`
	Name string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	Email string `json:"email,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty"`
	Gender string `json:"gender,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty"`
	UpdatedOn time.Time `json:"updated_on,omitempty"`
	IIN string `json:"iin,omitempty"`
	IdCardNumber string `json:"id_card_number,omitempty"`
}

type ListCustomers struct {
	Uid string `json:"uid,omitempty"`
	CustomerCategoryId []string `json:"customer_category_id,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Birthday time.Time `json:"birthday,omitempty"`
	LastBought time.Time `json:"last_bought,omitempty"`
	ToWhomId string `json:"to_whom_id,omitempty"`
	ToWhomFullname string `json:"to_whom_fullname,omitempty"`
	OrdersCountDone int64 `json:"orders_count_done"`
	StoreId string `json:"store_id,omitempty"` 
	OrdersCount int64 `json:"orders_count"`
	LastNotification time.Time `json:"last_notification"`
	Discount float64 `json:"discount"`
	Gender string `json:"gender"`
}


type MerchantsCustomers struct {
	MerchantId string `json:"merchant_id,omitempty"`
	Status string `json:"status,omitempty"`
}

type CustomerCategory struct {
	Scope [] string `json:"scope,omitempty"`
	Uid string `json:"uid,omitempty"`
	Title string `json:"title,omitempty"`
	MerchantID string `json:"merchant_id,omitempty"`
	Sort       int64  `json:"sort"`
}


