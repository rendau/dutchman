package entities

import (
	"encoding/json"

	"github.com/rendau/dop/dopTypes"
)

type RealmSt struct {
	Id   string          `json:"id" db:"id"`
	Data json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}

type RealmListParsSt struct {
	dopTypes.ListParams
}

type RealmCUSt struct {
	Data *json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}

type RealmDeployConfSt struct {
	Method   string `json:"method"`
	Url      string `json:"url"`
	ConfFile string `json:"conf_file"`
}

// Deploy

type RealmDeployReqSt struct {
	Config json.RawMessage `json:"config" swaggertype:"string"`
}
