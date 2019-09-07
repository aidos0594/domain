package domain

type CityItem struct {
	Code       string  `json:"code,omitempty"`
	Name       string  `json:"name,omitempty"`
	Slug       string  `json:"slug,omitempty"`
	KatoId     string  `json:"katoId,omitempty"`
	KatoName   string  `json:"katoName,omitempty"`
	externalId int64   `json:"externalId,omitempty"`
	Lat        float64 `json:"lat,omitempty"`
	Lon        float64 `json:"lon,omitempty"`
	Radius     float64 `json:"radius,omitempty"`
	Priority   int     `json:"priority,omitempty"`
}

type CountryItem struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type City struct {
	Country CountryItem `json:"country"`
	Region  []CityItem  `json:"region"`
	City    CityItem    `json:"city"`
}
