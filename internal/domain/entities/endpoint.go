package entities

import (
	"encoding/json"

	"github.com/rendau/dop/dopTypes"
)

type EndpointSt struct {
	Id     string          `json:"id" db:"id"`
	AppId  string          `json:"app_id" db:"app_id"`
	Active bool            `json:"active" db:"active"`
	Data   json.RawMessage `json:"data" db:"data" swaggertype:"string"`

	App *AppSt `json:"app" db:"-"`
}

type EndpointGetParsSt struct {
	WithApp bool `json:"with_app" form:"with_app"`
}

type EndpointListParsSt struct {
	dopTypes.ListParams

	AppId  *string `json:"app_id" form:"app_id"`
	Active *bool   `json:"active" form:"active"`
}

type EndpointCUSt struct {
	AppId  *string          `json:"app_id" db:"app_id"`
	Active *bool            `json:"active" db:"active"`
	Data   *json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}
