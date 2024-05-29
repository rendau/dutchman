package entities

import (
	"github.com/rendau/dop/dopTypes"
)

type AppSt struct {
	Id      string    `json:"id" db:"id"`
	RealmId string    `json:"realm_id" db:"realm_id"`
	Active  bool      `json:"active" db:"active"`
	Data    AppDataSt `json:"data" db:"data"`
}

type AppListParsSt struct {
	dopTypes.ListParams

	RealmId *string `json:"realm_id" form:"realm_id"`
	Active  *bool   `json:"active" form:"active"`
}

type AppCUSt struct {
	RealmId *string    `json:"realm_id" db:"realm_id"`
	Active  *bool      `json:"active" db:"active"`
	Data    *AppDataSt `json:"data" db:"data"`
}

// data

type AppDataSt struct {
	Name        string               `json:"name"`
	Path        string               `json:"path"`
	BackendBase AppDataBackendBaseSt `json:"backend_base"`
	RemoteRoles AppDataRemoteRolesSt `json:"remote_roles"`
}

type AppDataBackendBaseSt struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

type AppDataRemoteRolesSt struct {
	Url      string `json:"url"`
	JsonPath string `json:"json_path"`
}

// duplicate

type AppDuplicateReq struct {
	NewRealmId *string `json:"new_realm_id"`
	NewName    *string `json:"new_name"`
}
