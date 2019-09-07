package domain

import "time"

type Supplier struct {
	Id        string    `json:"id"`
	IIN       string    `json:"iin"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	Phone     []string  `json:"phone"`
	Email     string    `json:"email"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`

	//how to do merchant comment
}
