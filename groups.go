package domain

type Groups struct {
	Uid   string `json:"uid,omitempty"`
	Scope string `json:"scope,omitempty"`
	//TypeOfGroup string `json:"type_of_group"`
	//GroupID
	MerchantID string  `json:"merchant_id,omitempty"`
	Title      string  `json:"title,omitempty"`
	Sort       int64   `json:"sort"`
	Medias     []Media `json:"media,omitempty"`
}
