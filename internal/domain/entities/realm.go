package entities

import (
	"github.com/rendau/dop/dopTypes"
)

type RealmSt struct {
	Id   string      `json:"id" db:"id"`
	Data RealmDataSt `json:"data" db:"data"`
}

type RealmListParsSt struct {
	dopTypes.ListParams
}

type RealmCUSt struct {
	Data *RealmDataSt `json:"data" db:"data"`
}

// data

type RealmDataSt struct {
	Name              string                `json:"name"`
	PublicBaseUrl     string                `json:"public_base_url"`
	Timeout           string                `json:"timeout"`
	ReadHeaderTimeout string                `json:"read_header_timeout"`
	ReadTimeout       string                `json:"read_timeout"`
	DeployConf        RealmDataDeployConfSt `json:"deploy_conf"`
	CorsConf          RealmDataCorsConfSt   `json:"cors_conf"`
	JwtConf           RealmDataJwtConfSt    `json:"jwt_conf"`
}

type RealmDataDeployConfSt struct {
	ConfFile string `json:"conf_file"`
	Url      string `json:"url"`
	Method   string `json:"method"`
}

type RealmDataCorsConfSt struct {
	Enabled          bool     `json:"enabled"`
	AllowCredentials bool     `json:"allow_credentials"`
	MaxAge           string   `json:"max_age"`
	AllowOrigins     []string `json:"allow_origins,omitempty"`
	AllowMethods     []string `json:"allow_methods,omitempty"`
	AllowHeaders     []string `json:"allow_headers,omitempty"`
}

type RealmDataJwtConfSt struct {
	Alg                string `json:"alg"`
	JwkUrl             string `json:"jwk_url"`
	DisableJwkSecurity bool   `json:"disable_jwk_security"`
	Cache              bool   `json:"cache"`
	CacheDuration      int64  `json:"cache_duration"`
	RolesKey           string `json:"roles_key"`
	RolesKeyIsNested   bool   `json:"roles_key_is_nested"`
}
