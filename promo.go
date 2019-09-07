package domain

import "time"


type Agent struct {
	Uid               string    `json:"agent_uid,omitempty"`
	Name              string    `json:"name,omitempty"`
	Surname           string    `json:"surname,omitempty"`
	Promo			  string    `json:"promo_code,omitempty"`
	Cellphone         string    `json:"cellphone,omitempty"`
	Email			  string    `json:"email,omitempty"`
	UseLimit		  int       `json:"use_limit,omitempty"`
	Discount          Discount  `json:"discount,omitempty"`
	CreatedOn         time.Time `json:"created_on,omitempty"`
	UpdatedOn         time.Time `json:"updated_on,omitempty"`
}

type Discount struct {
	Merchants 		 []string    `json:"merchants,omitempty"`
	Percent 		 float32	 `json:"percent,omitempty"`
	MinTreshold 	 float32	 `json:"min_treshold,omitempty"`
	MaxTreshold 	 float32	 `json:"max_treshold,omitempty"`
	TriggerPrice	 float32	 `json:"trigger_price,omitempty"`
}
