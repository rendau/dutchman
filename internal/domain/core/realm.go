package core

import (
	"context"
	"github.com/rendau/dutchman/internal/cns"
	"path"
	"strings"

	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
)

type Realm struct {
	r *St
}

func NewRealm(r *St) *Realm {
	return &Realm{r: r}
}

func (c *Realm) ValidateCU(ctx context.Context, obj *entities.RealmCUSt, id string) error {
	// forCreate := id == ""

	return nil
}

func (c *Realm) List(ctx context.Context, pars *entities.RealmListParsSt) ([]*entities.RealmSt, int64, error) {
	items, tCount, err := c.r.repo.RealmList(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	return items, tCount, nil
}

func (c *Realm) Get(ctx context.Context, id string, errNE bool) (*entities.RealmSt, error) {
	result, err := c.r.repo.RealmGet(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		if errNE {
			return nil, dopErrs.ObjectNotFound
		}
		return nil, nil
	}

	return result, nil
}

func (c *Realm) IdExists(ctx context.Context, id string) (bool, error) {
	return c.r.repo.RealmIdExists(ctx, id)
}

func (c *Realm) Create(ctx context.Context, obj *entities.RealmCUSt) (string, error) {
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	// create
	result, err := c.r.repo.RealmCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Realm) Update(ctx context.Context, id string, obj *entities.RealmCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	err = c.r.repo.RealmUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Realm) Delete(ctx context.Context, id string) error {
	return c.r.repo.RealmDelete(ctx, id)
}

func (c *Realm) GenerateConf(ctx context.Context, id string) (*entities.KrakendSt, error) {
	realm, err := c.Get(ctx, id, true)
	if err != nil {
		return nil, err
	}

	result := &entities.KrakendSt{
		Schema:            "https://www.krakend.io/schema/v3.json",
		Version:           3,
		Timeout:           realm.Data.Timeout,
		ReadHeaderTimeout: realm.Data.ReadHeaderTimeout,
		ReadTimeout:       realm.Data.ReadTimeout,
		Endpoints:         make([]entities.KrakendEndpointSt, 0),
	}

	if realm.Data.CorsConf.Enabled {
		if result.ExtraConfig == nil {
			result.ExtraConfig = &entities.KrakendExtraConfigSt{}
		}

		result.ExtraConfig.SecurityCors = &entities.KrakendExtraConfigSecurityCorsSt{
			ExposeHeaders:    "*",
			AllowCredentials: realm.Data.CorsConf.AllowCredentials,
			MaxAge:           realm.Data.CorsConf.MaxAge,
		}
		if len(realm.Data.CorsConf.AllowOrigins) > 0 {
			result.ExtraConfig.SecurityCors.AllowOrigins = realm.Data.CorsConf.AllowOrigins
		}
		if len(realm.Data.CorsConf.AllowMethods) > 0 {
			result.ExtraConfig.SecurityCors.AllowMethods = realm.Data.CorsConf.AllowMethods
		}
		if len(realm.Data.CorsConf.AllowHeaders) > 0 {
			result.ExtraConfig.SecurityCors.AllowHeaders = realm.Data.CorsConf.AllowHeaders
		}
	}

	apps, _, err := c.r.App.List(ctx, &entities.AppListParsSt{
		RealmId: &realm.Id,
		Active:  &cns.True,
	})
	if err != nil {
		return nil, err
	}

	for _, app := range apps {
		endpoints, _, err := c.r.Endpoint.List(ctx, &entities.EndpointListParsSt{
			AppId:  &app.Id,
			Active: &cns.True,
		})
		if err != nil {
			return nil, err
		}

		for _, endpoint := range endpoints {
			//switch endpoint.Id {
			//case "8252c666-62d7-4e26-9216-a678f92f5ae9",
			//	"c010c570-6805-45df-a662-1d7b9d935ea7":
			//default:
			//	continue
			//}

			epPath := endpoint.Data.Path
			if endpoint.Data.Backend.CustomPath {
				epPath = endpoint.Data.Backend.Path
			}

			resEndpoint := entities.KrakendEndpointSt{
				Endpoint:          "/" + strings.TrimLeft(path.Join(app.Data.Path, endpoint.Data.Path), "/"),
				Method:            endpoint.Data.Method,
				OutputEncoding:    "no-op",
				InputQueryStrings: []string{"*"},
				InputHeaders:      []string{"*"},
				Backend: []entities.KrakendEndpointBackendSt{
					{
						Method:     endpoint.Data.Method,
						UrlPattern: "/" + strings.TrimLeft(path.Join(app.Data.BackendBase.Path, epPath), "/"),
						Encoding:   "no-op",
						Host:       []string{app.Data.BackendBase.Host},
					},
				},
			}

			if endpoint.Data.JwtValidation.Enabled && realm.Data.JwtConf.JwkUrl != "" {
				if resEndpoint.ExtraConfig == nil {
					resEndpoint.ExtraConfig = &entities.KrakendEndpointExtraConfigSt{}
				}

				resEndpoint.ExtraConfig.AuthValidator = &entities.KrakendEndpointExtraConfigAuthValidatorSt{
					Alg:                realm.Data.JwtConf.Alg,
					JwkUrl:             realm.Data.JwtConf.JwkUrl,
					DisableJwkSecurity: realm.Data.JwtConf.DisableJwkSecurity,
					Cache:              realm.Data.JwtConf.Cache,
					CacheDuration:      realm.Data.JwtConf.CacheDuration,
				}

				if len(endpoint.Data.JwtValidation.Roles) > 0 {
					resEndpoint.ExtraConfig.AuthValidator.Roles = endpoint.Data.JwtValidation.Roles
					resEndpoint.ExtraConfig.AuthValidator.RolesKey = realm.Data.JwtConf.RolesKey
					resEndpoint.ExtraConfig.AuthValidator.RolesKeyIsNested = realm.Data.JwtConf.RolesKeyIsNested
				}
			}

			if endpoint.Data.IpValidation.Enabled && len(endpoint.Data.IpValidation.AllowedIps) > 0 {
				if resEndpoint.ExtraConfig == nil {
					resEndpoint.ExtraConfig = &entities.KrakendEndpointExtraConfigSt{}
				}

				resEndpoint.ExtraConfig.ValidationCel = []entities.KrakendEndpointExtraConfigValidationCelSt{
					{
						CheckExpr: `req_headers['X-Real-Ip'][0] in ['` + strings.Join(endpoint.Data.IpValidation.AllowedIps, "','") + `']`,
					},
				}
			}

			result.Endpoints = append(result.Endpoints, resEndpoint)
		}
	}

	return result, nil
}
