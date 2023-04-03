package entities

import (
	"encoding/json"

	"github.com/rendau/dop/dopTypes"
)

type RoleSt struct {
	Id        string          `json:"id" db:"id"`
	AppId     string          `json:"app_id" db:"app_id"`
	IsFetched bool            `json:"is_fetched" db:"is_fetched"`
	Data      json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}

type RoleListParsSt struct {
	dopTypes.ListParams

	AppId *string `json:"app_id" form:"app_id"`
}

type RoleCUSt struct {
	AppId     *string          `json:"app_id" db:"app_id"`
	IsFetched *bool            `json:"is_fetched" db:"is_fetched"`
	Data      *json.RawMessage `json:"data" db:"data" swaggertype:"string"`
}

// remote

type RoleRemoteRepItemSt struct {
	Code string `json:"code"`
	Dsc  string `json:"dsc"`
}
