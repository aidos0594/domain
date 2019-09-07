package domain


type LogexResponse struct {
	Status   string     `json:"status,omitempty"`
	Message  string     `json:"message,omitempty"`
	Data     DataItems  `json:"data,omitempty"`
}

type ItemsParameters struct {
	Cost               int         `json:"cost,omitempty"`
	Address            string      `json:"address"`
	Price              float64     `json:"price"`
	Count              float64         `json:"count"`
	City               string      `json:"city"`
	Weight             float64     `json:"weight"`
	MerchantName       string      `json:"name,omitempty"`
	MobilePhone        string      `json:"phone,omitempty"`
	ProductName        string      `json:"product,omitempty"`
	MerchantEmail      string      `json:"email,omitempty"`
	MerchantID         string      `json:"merchant_id,omitempty"`
	ProductDimensions  Dimensions  `json:"dimensions,omitempty"`
}

type DataItems struct {
	TotalCost    float64          `json:"total_cost,omitempty"`
	TotalWeight  int          `json:"total_weight,omitempty"`
	PaymentType  string       `json:"payment_type,omitempty"`
	PaymentSum   int          `json:"payment_sum,omitempty"`
	From         []ItemsParameters  `json:"from,omitempty"`
	To           ItemsParameters      `json:"to,omitempty"`
	DType 		 string       `json:"type,omitempty"`
	APIKey       string             `json:"api_key,omitempty"`
	SendDate     string             `json:"send_date,omitempty"`
	Description  string             `json:"description,omitempty"`
	OrderID      string                `json:"id,omitempty"`
	WaybillID    int      `json:"waybill_id,omitempty"`
}
