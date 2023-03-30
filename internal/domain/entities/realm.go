package entities

import (
	"encoding/json"
)

type RealmSt struct {
	Id   string          `json:"id" db:"id"`
	Data json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}

type RealmCUSt struct {
	Name *string          `json:"name" db:"name"`
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
