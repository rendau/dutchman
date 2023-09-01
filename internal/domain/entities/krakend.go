package entities

type KrakendSt struct {
	Schema            string                `json:"$schema"`
	Version           int                   `json:"version"`
	Timeout           string                `json:"timeout"`
	ReadHeaderTimeout string                `json:"read_header_timeout,omitempty"`
	ReadTimeout       string                `json:"read_timeout,omitempty"`
	Endpoints         []*KrakendEndpointSt  `json:"endpoints"`
	ExtraConfig       *KrakendExtraConfigSt `json:"extra_config,omitempty"`
}

type KrakendEndpointSt struct {
	Endpoint          string                        `json:"endpoint"`
	Method            string                        `json:"method"`
	OutputEncoding    string                        `json:"output_encoding"`
	InputQueryStrings []string                      `json:"input_query_strings"`
	InputHeaders      []string                      `json:"input_headers"`
	Backend           []KrakendEndpointBackendSt    `json:"backend"`
	ExtraConfig       *KrakendEndpointExtraConfigSt `json:"extra_config,omitempty"`
}

type KrakendEndpointBackendSt struct {
	Method     string   `json:"method,omitempty"`
	UrlPattern string   `json:"url_pattern"`
	Encoding   string   `json:"encoding"`
	Host       []string `json:"host"`
}

type KrakendEndpointExtraConfigSt struct {
	AuthValidator *KrakendEndpointExtraConfigAuthValidatorSt  `json:"auth/validator,omitempty"`
	ValidationCel []KrakendEndpointExtraConfigValidationCelSt `json:"validation/cel,omitempty"`
}

type KrakendEndpointExtraConfigAuthValidatorSt struct {
	Alg                string   `json:"alg"`
	JwkUrl             string   `json:"jwk_url"`
	DisableJwkSecurity bool     `json:"disable_jwk_security"`
	Cache              bool     `json:"cache"`
	CacheDuration      int64    `json:"cache_duration"`
	Roles              []string `json:"roles,omitempty"`
	RolesKey           string   `json:"roles_key,omitempty"`
	RolesKeyIsNested   bool     `json:"roles_key_is_nested"`
}

type KrakendEndpointExtraConfigValidationCelSt struct {
	CheckExpr string `json:"check_expr"`
}

type KrakendExtraConfigSt struct {
	SecurityCors *KrakendExtraConfigSecurityCorsSt `json:"security/cors,omitempty"`
}

type KrakendExtraConfigSecurityCorsSt struct {
	ExposeHeaders    []string `json:"expose_headers"`
	AllowCredentials bool     `json:"allow_credentials"`
	MaxAge           string   `json:"max_age,omitempty"`
	AllowOrigins     []string `json:"allow_origins"`
	AllowMethods     []string `json:"allow_methods"`
	AllowHeaders     []string `json:"allow_headers"`
}
