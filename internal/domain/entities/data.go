package entities

import (
	"encoding/json"
)

type DataSt struct {
	Id   string          `json:"id" db:"id"`
	Name string          `json:"name" db:"name"`
	Val  json.RawMessage `json:"val" db:"val" swaggertype:"string"`
}

type DataListSt struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type DataCUSt struct {
	Name *string          `json:"name" db:"name"`
	Val  *json.RawMessage `json:"val" db:"val" swaggertype:"string"`
}

// Deploy

type DataDeployReqSt struct {
	ConfFile string          `json:"conf_file"`
	Url      string          `json:"url"`
	Method   string          `json:"method"`
	Data     json.RawMessage `json:"data" swaggertype:"string"`
}
