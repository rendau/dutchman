package entities

import (
	"encoding/json"

	"github.com/rendau/dop/dopTypes"
)

type PermSt struct {
	Id    string          `json:"id" db:"id"`
	AppId string          `json:"app_id" db:"app_id"`
	Data  json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}

type PermListParsSt struct {
	dopTypes.ListParams

	AppId *string `json:"app_id" form:"app_id"`
}

type PermCUSt struct {
	AppId *string          `json:"app_id" db:"app_id"`
	Data  *json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}

// remote

type PermRemoteRepSt struct {
	Perms []*PermRemoteRepItemSt `json:"perms"`
}

type PermRemoteRepItemSt struct {
	Code  string `json:"code"`
	IsAll bool   `json:"is_all"`
	Dsc   string `json:"dsc"`
}
