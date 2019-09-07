package domain

import "time"

type Bundle struct {
	Uid               string    `json:"uid,omitempty"`
	Scope             []string    `json:"scope,omitempty"`
	MerchantID        string    `json:"merchant_id,omitempty"`
	Title             string    `json:"title,omitempty"`
	Label             string    `json:"label,omitempty"` //подарок или комбо
	Items             []Product `json:"items,omitempty"`
	CreatedOn         time.Time `json:"created_on,omitempty"`
	UpdatedOn         time.Time `json:"updated_on,omitempty"`
	CommonPrice       float64   `json:"common_price"`
	CommonRetailPrice float64   `json:"common_retail_price"`
	CommonCount       int       `json:"common_count"`
}
