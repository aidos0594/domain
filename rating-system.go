package domain

type ProductCard struct {
	Uid            string  `json:"uid"`
	Name           string  `json:"name"`
	Media          []Media `json:"media"`
	AssessmentRate int     `json:"assessment_rate"`
	ProductComment string  `json:"product_comment"`
}

type ServiceCard struct {
	SpeedOfOrderProcessing           int `json:"speed_of_order_processing"`
	CallCenterConsultationPoliteness int `json:"call_center"`
	PostService                      int `json:"post_service"`
	DarbazarOverAll                  int `json:"darbazar_over_all"`
}

type DeliveryCard struct {
	SpeedOfDelivery      int `json:"speed_of_delivery"`
	CourierAppearance    int `json:"courier_appearance"`
	CourierCommunication int `json:"courier_communication"`
	ParcelAppearance     int `json:"parcel_appearance"`
	ChangesCOD           int `json:"changes_cod"`
}

type OrderAssessmentCard struct {
	JsonClass string        `json:"jsonClass,omitempty"`
	Delivery  DeliveryCard  `json:"delivery"`
	Products  []ProductCard `json:"products"`
	Service   ServiceCard   `json:"service"`
	OwnType   string        `json:"own_type"`
	Owner     string        `json:"owner"`
	Comment   string        `json:"comment"`
}
