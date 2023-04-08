package entities

import (
	"github.com/rendau/dop/dopTypes"
)

type EndpointSt struct {
	Id     string         `json:"id" db:"id"`
	AppId  string         `json:"app_id" db:"app_id"`
	Active bool           `json:"active" db:"active"`
	Data   EndpointDataSt `json:"data" db:"data"`
}

type EndpointGetParsSt struct {
}

type EndpointListParsSt struct {
	dopTypes.ListParams

	AppId  *string `json:"app_id" form:"app_id"`
	Active *bool   `json:"active" form:"active"`
}

type EndpointCUSt struct {
	AppId  *string         `json:"app_id" db:"app_id"`
	Active *bool           `json:"active" db:"active"`
	Data   *EndpointDataSt `json:"data" db:"data"`
}

// data

type EndpointDataSt struct {
	Method        string                      `json:"method"`
	Path          string                      `json:"path"`
	Backend       EndpointDataBackendSt       `json:"backend"`
	JwtValidation EndpointDataJwtValidationSt `json:"jwt_validation"`
	IpValidation  EndpointDataIpValidationSt  `json:"ip_validation"`
}

type EndpointDataBackendSt struct {
	CustomPath bool   `json:"custom_path"`
	Path       string `json:"path"`
}

type EndpointDataJwtValidationSt struct {
	Enabled bool     `json:"enabled"`
	Roles   []string `json:"roles"`
}

type EndpointDataIpValidationSt struct {
	Enabled    bool     `json:"enabled"`
	AllowedIps []string `json:"allowed_ips"`
}
