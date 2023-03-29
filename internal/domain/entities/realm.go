package entities

import (
	"encoding/json"
)

type RealmSt struct {
	Id   string          `json:"id" db:"id"`
	Name string          `json:"name" db:"name"`
	Conf json.RawMessage `json:"conf" db:"conf" swaggertype:"string"`
}

type RealmListSt struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type RealmCUSt struct {
	Name *string          `json:"name" db:"name"`
	Conf *json.RawMessage `json:"conf" db:"conf" swaggertype:"string"`
}

// Deploy

type RealmDeployReqSt struct {
	Data json.RawMessage `json:"data" swaggertype:"string"`
}
