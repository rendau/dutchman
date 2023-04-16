package entities

import (
	"github.com/rendau/dop/dopTypes"
)

type RoleSt struct {
	Id        string  `json:"id" db:"id"`
	RealmId   string  `json:"realm_id" db:"realm_id"`
	AppId     *string `json:"app_id" db:"app_id"`
	IsFetched bool    `json:"is_fetched" db:"is_fetched"`
	Code      string  `json:"code" db:"code"`
	Dsc       string  `json:"dsc" db:"dsc"`
}

type RoleListParsSt struct {
	dopTypes.ListParams

	RealmId     *string `json:"realm_id" form:"realm_id"`
	AppId       *string `json:"app_id" form:"app_id"`
	AppIdOrNull *string `json:"app_id_or_null" form:"app_id_or_null"`
	IsFetched   *bool   `json:"is_fetched" form:"is_fetched"`
}

type RoleCUSt struct {
	RealmId   *string  `json:"realm_id" db:"realm_id"`
	AppId     *string  `json:"app_id" db:"-"`
	DbAppId   **string `json:"-" db:"app_id"`
	IsFetched *bool    `json:"is_fetched" db:"is_fetched"`
	Code      *string  `json:"code" db:"code"`
	Dsc       *string  `json:"dsc" db:"dsc"`
}

// remote

type RoleFetchRemoteReqSt struct {
	Uri  string `json:"uri"`
	Path string `json:"path"`
}

type RoleFetchRemoteRepItemSt struct {
	Code string `json:"code"`
	Dsc  string `json:"dsc"`
}
