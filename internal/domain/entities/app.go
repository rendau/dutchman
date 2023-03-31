package entities

import (
	"encoding/json"

	"github.com/rendau/dop/dopTypes"
)

type AppSt struct {
	Id      string          `json:"id" db:"id"`
	RealmId string          `json:"realm_id" db:"realm_id"`
	Active  bool            `json:"active" db:"active"`
	Data    json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}

type AppListParsSt struct {
	dopTypes.ListParams

	RealmId *string `json:"realm_id" form:"realm_id"`
	Active  *bool   `json:"active" form:"active"`
}

type AppCUSt struct {
	RealmId *string          `json:"realm_id" db:"realm_id"`
	Active  *bool            `json:"active" db:"active"`
	Data    *json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}
