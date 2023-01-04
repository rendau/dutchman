package entities

import (
	"encoding/json"
)

type DataSt struct {
	Id   string          `json:"id" db:"id"`
	Name string          `json:"name" db:"name"`
	Val  json.RawMessage `json:"val" db:"val"`
}

type DataListSt struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type DataCUSt struct {
	Name *string          `json:"name" db:"name"`
	Val  *json.RawMessage `json:"val" db:"val"`
}
