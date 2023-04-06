package entities

import (
	"github.com/rendau/dop/dopTypes"
)

type RoleSt struct {
	Id        string  `json:"id" db:"id"`
	AppId     *string `json:"app_id" db:"app_id"`
	IsFetched bool    `json:"is_fetched" db:"is_fetched"`
	Code      string  `json:"code" db:"code"`
	Dsc       string  `json:"dsc" db:"dsc"`
}

type RoleListParsSt struct {
	dopTypes.ListParams

	AppId       *string `json:"app_id" form:"app_id"`
	AppIdOrNull *string `json:"app_id_or_null" form:"app_id_or_null"`
	IsFetched   *bool   `json:"is_fetched" form:"is_fetched"`
}

type RoleCUSt struct {
	AppId     *string  `json:"app_id" db:"-"`
	DbAppId   **string `json:"-" db:"app_id"`
	IsFetched *bool    `json:"is_fetched" db:"is_fetched"`
	Code      *string  `json:"code" db:"code"`
	Dsc       *string  `json:"dsc" db:"dsc"`
}

// remote

type RoleRemoteRepItemSt struct {
	Code string `json:"code"`
	Dsc  string `json:"dsc"`
}
