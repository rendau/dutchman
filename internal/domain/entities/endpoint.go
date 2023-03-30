package entities

import (
	"encoding/json"
)

type EndpointSt struct {
	Id     string          `json:"id" db:"id"`
	AppId  string          `json:"app_id" db:"app_id"`
	Active bool            `json:"active" db:"active"`
	Data   json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}

type EndpointCUSt struct {
	AppId  *string          `json:"app_id" db:"app_id"`
	Active *bool            `json:"active" db:"active"`
	Data   *json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}
