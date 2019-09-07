package domain

import (
	"encoding/json"
	"strings"
)

type WarehouseType int

const (
	WarehouseTypeUnknown WarehouseType = iota
	WarehouseTypeStorage
	WarehouseTypeCentral
)

var WarehouseTypes = []WarehouseType{
	WarehouseTypeUnknown,
	WarehouseTypeStorage,
	WarehouseTypeCentral,
}

func (a WarehouseType) String() string {
	switch a {
	default:
		return "unknown"
	case WarehouseTypeCentral:
		return "central"
	case WarehouseTypeStorage:
		return "storage"
	}
	return "unknown"
}

func (a *WarehouseType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = WarehouseTypeUnknown
	case "central":
		*a = WarehouseTypeCentral
	case "storage":
		*a = WarehouseTypeStorage
	}

	return nil
}

func (a WarehouseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}
