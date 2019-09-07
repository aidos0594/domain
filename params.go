package domain

type Params struct {
	Uid             string         `json:"uid"`
	CategoryUid     string         `json:"category_uid"`
	Title           string         `json:"title,omitempty"`
	Code            string         `json:"code,omitempty"`
	ParamsV         []ParamsValues `json:"params_values,omitempty"`
	Medias          []Media        `json:"media,omitempty"`
	Ptype           string         `json:"ptype,omitempty"`
	Ftype           string         `json:"ftype,omitempty"`
	Unit            string         `json:"unit,omitempty"`
	IsRequired      bool           `json:"is_required"`
	PositionOnFront string         `json:"position_on_front,omitempty"`
	PositionSort    int            `json:"position_sort,omitempty"`
	MigrationState  string         `json:"migration_state,omitempty"`
}

type ParamsValues struct {
	Uid        string  `json:"uid_values,omitempty"`
	CodeValue  string  `json:"code_value,omitempty"`
	Unit       string  `json:"unit,omitempty"`
	Values     string  `json:"values,omitempty"`
	SortValues int     `json:"sort_values,omitempty"`
	Medias     []Media `json:"media,omitempty"`
}

type Media struct {
	MediaType string `json:"media_type,omitempty"`
	MediaUrl  string `json:"media_url,omitempty"`
	AltText   string `json:"alt_text,omitempty"`
}
