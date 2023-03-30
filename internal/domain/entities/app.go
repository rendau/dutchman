package entities

import (
	"encoding/json"
)

type AppSt struct {
	Id      string          `json:"id" db:"id"`
	RealmId string          `json:"realm_id" db:"realm_id"`
	Active  bool            `json:"active" db:"active"`
	Data    json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}

type AppCUSt struct {
	RealmId *string          `json:"realm_id" db:"realm_id"`
	Active  *bool            `json:"active" db:"active"`
	Data    *json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}
