package domain

import "time"

type Characteristic struct {
	Uid               string     `json:"uid"`
	NameRu            string     `json:"nameRu"`
	ValueTitleRu      string     `json:"valueTitleRu,omitempty"`
	NameEn            string     `json:"nameEn"`
	ValueTitleEn      string     `json:"valueTitleEn,omitempty"`
	NameKk            string     `json:"nameKk"`
	ValueTitleKk      string     `json:"valueTitleKk,omitempty"`
	ValueId           string     `json:"valueId,omitempty"`
	Mask              string     `json:"mask,omitempty"`
	SortPriority      float64    `json:"sortPriority,omitempty"`
	ValueSortPriority float64    `json:"valueSortPriority,omitempty"`
	CreatedOn         *time.Time `json:"created_on,omitempty"`
	UpdatedOn         *time.Time `json:"updated_on,omitempty"`
}

type ValueCharacteristic struct { // unused
	Uid               string      `json:"uid,omitempty"`
	ValueId           string      `json:"value_id,omitempty"`
	ValueSortPriority float64     `json:"value_sort_priority,omitempty"`
	ValueNames        []Languages `json:"value_names,omitempty"`
	CreatedOn         *time.Time  `json:"created_on,omitempty"`
	UpdatedOn         *time.Time  `json:"updated_on,omitempty"`
}

type CharacteristicSet struct {
	Uid                  string                `json:"uid"`
	Names                []Languages           `json:"names,omitempty"`
	CreatedOn            *time.Time            `json:"created_on,omitempty"`
	UpdatedOn            *time.Time            `json:"updated_on,omitempty"`
	CharacteristicGroups []CharacteristicGroup `json:"characteristic_groups"`
}

type Languages struct {
	Title      string `json:"title"`
	LangPrefix string `json:"lang_prefix"`
}

type CharacteristicGroup struct {
	GroupNames      []Languages `json:"group_names"`
	Characteristics []string    `json:"characteristics"`
}

type OldCharToNewCharsMapping struct {
	NewCharNameId   string   `json:"newCharNameId"`
	NewCharValueId  string   `json:"newCharValueId"`
	OldCharValueIds []string `json:"oldCharValueIds"`
}

type CharacNames struct {
	NameRu             string     `json:"nameRu,omitempty"`
	NameEn             string     `json:"nameEn,omitempty"`
	NameKk             string     `json:"nameKk,omitempty"`
	Uid                string     `json:"uid,omitempty"`
	SortPriority       float64    `json:"sortPriority,omitempty"`
	CreatedOn          *time.Time `json:"created_on,omitempty"`
	UpdatedOn          *time.Time `json:"updated_on,omitempty"`
	IsNecessarily      bool       `json:"isNecessarily"`
	UseOnFilter        bool       `json:"useOnFilter"`
	OnShowcase         bool       `json:"onShowcase"`
	UpdatedByProfileId string     `json:"updated_by_profile_id,omitempty"`
	UpdatedByFio       string     `json:"updated_by_fio,omitempty"`
}