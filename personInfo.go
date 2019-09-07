package domain

import "time"

type PersonInfo struct {
	FinancingInfo    FinancingInfo   `json:"financing_info"`
	ProductInfo      ProductInfo     `json:"product_info"`
	Surname          string          `json:"surname" validate:"required"`
	Name             string          `json:"name" validate:"required"`
	Patronymic       string          `json:"patronymic,omitempty"`
	BirthDate        time.Time       `json:"birthdate" validate:"required"`
	Sex              int             `json:"sex" validate:"required,numeric,min=1,max=2"`
	IIN              string          `json:"iin" validate:"required,numeric,len=12"`
	MaritalStatus    string          `json:"marital_status" validate:"required,numeric,len=9"`
	DocumentType     string          `json:"document_type,omitempty"`
	DocumentId       string          `json:"document_id" validate:"required,numeric"`
	IssuedBy         string          `json:"issued_by" validate:"required,numeric,len=9"`
	IssuedDate       time.Time       `json:"issued_date" validate:"required"`
	ValidityDate     time.Time       `json:"validity_date" validate:"required"`
	PlaceOfBirthId   string          `json:"place_of_birth_id" validate:"required"`
	PlaceOfBirthName string          `json:"place_of_birth_name" validate:"required"`
	MobilePhone      string          `json:"mobile_phone" validate:"required,numeric"`
	RegPostalIndex   string          `json:"reg_postal_index"`
	RegLocality      string          `json:"reg_locality" validate:"required"`
	RegLocalityName  string          `json:"reg_locality_name,omitempty"`
	RegMicrodistrict string          `json:"reg_microdistrict" validate:"required"`
	RegStreet        string          `json:"reg_street" validate:"required"`
	RegHouseNum      string          `json:"reg_house_number" validate:"required"`
	RegApartmentNum  string          `json:"reg_apart_number,omitempty"`
	RegTerm          int             `json:"reg_term,omitempty"`
	RegLiveWith      string          `json:"reg_live_with" validate:"required,numeric"`
	RegMatchRes      bool            `json:"reg_match_res"`
	ResPostalIndex   string          `json:"res_postal_index,omitempty"`
	ResLocality      string          `json:"res_locality,omitempty"`
	ResLocalityName  string          `json:"res_locality_name,omitempty"`
	ResMicrodistrict string          `json:"res_microdistrict,omitempty"`
	ResStreet        string          `json:"res_street,omitempty"`
	ResHouseNum      string          `json:"res_house_number,omitempty"`
	ResApartmentNum  string          `json:"res_apart_number,omitempty"`
	ResLiveWith      string          `json:"res_live_with,omitempty"`
	Wage             float64         `json:"wage" validate:"required,numeric"`
	ContactPersons   []ContactPerson `json:"contact_persons" validate:"required"`
	ClientHostInfo   string          `json:"client_host_info,omitempty"`
	Timestamp        time.Time       `json:"timestamp,omitempty"`
}

type ContactPerson struct {
	Surname     string `json:"surname" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Patronymic  string `json:"patronymic,omitempty"`
	MobilePhone string `json:"mobile_phone" validate:"required,numeric"`
	Relation    string `json:"relation" validate:"required"`
}

type FinancingInfo struct {
	Product          string  `json:"product,omitempty"`
	ProductSubCode   string  `json:"product_sub_code,omitempty"`
	Sum              float64 `json:"sum,omitempty"`
	Period           int     `json:"period" validate:"required,numeric,min=4,max=24"`
	FinancingPurpose string  `json:"fin_purpose,omitempty"`
}

type ProductInfo struct {
	Price    float64 `json:"price,omitempty"`
	Model    string  `json:"description,omitempty"`
	Category string  `json:"category,omitempty"`
}
