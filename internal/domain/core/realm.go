package core

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/rendau/dop/adapters/client/httpc"
	"github.com/rendau/dop/adapters/client/httpc/httpclient"
	"github.com/rendau/dop/dopErrs"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/rendau/dutchman/internal/cns"
	"github.com/rendau/dutchman/internal/domain/entities"
	"github.com/rendau/dutchman/internal/domain/errs"
	"github.com/rendau/dutchman/internal/domain/util"
)

type Realm struct {
	r                      *St
	k8sRestartResourceType string
}

func NewRealm(r *St, k8sRestartResourceType string) *Realm {
	return &Realm{
		r:                      r,
		k8sRestartResourceType: k8sRestartResourceType,
	}
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
		Endpoints:         make([]*entities.KrakendEndpointSt, 0),
	}

	if realm.Data.CorsConf.Enabled {
		if result.ExtraConfig == nil {
			result.ExtraConfig = &entities.KrakendExtraConfigSt{}
		}

		result.ExtraConfig.SecurityCors = &entities.KrakendExtraConfigSecurityCorsSt{
			ExposeHeaders:    []string{"*"},
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
			// switch endpoint.Id {
			// case "8252c666-62d7-4e26-9216-a678f92f5ae9",
			//	"c010c570-6805-45df-a662-1d7b9d935ea7":
			// default:
			//	continue
			// }

			epPath := endpoint.Data.Path
			if endpoint.Data.Backend.CustomPath {
				epPath = endpoint.Data.Backend.Path
			}

			resEndpoint := &entities.KrakendEndpointSt{
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
					OperationDebug:     true,
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

func (c *Realm) ImportConf(ctx context.Context, id string, cfg *entities.KrakendSt) error {
	// KrakendExtraConfigSecurityCorsSt.extra_config.security/cors.expose_headers

	ipRegexp := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)

	realm, err := c.Get(ctx, id, true)
	if err != nil {
		return err
	}

	// delete all apps
	err = c.r.App.DeleteMany(ctx, &entities.AppListParsSt{
		RealmId: &realm.Id,
	})
	if err != nil {
		return err
	}

	type endpointSt struct {
		commonSuffix string
		ep           *entities.KrakendEndpointSt
	}

	type appSt struct {
		beHost       string
		bePathPrefix string
		epPathPrefix string
		endpoints    []*endpointSt
	}

	appMap := map[string]*appSt{}

	for _, endpoint := range cfg.Endpoints {
		if len(endpoint.Backend) == 0 {
			continue
		}

		be := endpoint.Backend[0]

		if strings.ToLower(be.Method) != strings.ToLower(endpoint.Method) {
			continue
		}

		if len(be.Host) == 0 {
			continue
		}

		beHost := be.Host[0]

		epPathPrefix, bePathPrefix, commonSuffix := util.StrTrimCommonSuffix(endpoint.Endpoint, be.UrlPattern)

		epPathPrefix = strings.TrimLeft(epPathPrefix, "/")
		bePathPrefix = strings.TrimLeft(bePathPrefix, "/")
		commonSuffix = strings.TrimLeft(commonSuffix, "/")

		appKey := fmt.Sprintf("%s__**%s__**%s", beHost, bePathPrefix, epPathPrefix)

		app, ok := appMap[appKey]
		if ok {
			app.endpoints = append(app.endpoints, &endpointSt{
				commonSuffix: commonSuffix,
				ep:           endpoint,
			})
		} else {
			appMap[appKey] = &appSt{
				beHost:       beHost,
				bePathPrefix: bePathPrefix,
				epPathPrefix: epPathPrefix,
				endpoints: []*endpointSt{
					{
						commonSuffix: commonSuffix,
						ep:           endpoint,
					},
				},
			}
		}
	}

	for _, app := range appMap {
		name := app.epPathPrefix
		if name == "" {
			name = "---"
		}

		appId, err := c.r.App.Create(ctx, &entities.AppCUSt{
			RealmId: &realm.Id,
			Active:  &cns.True,
			Data: &entities.AppDataSt{
				Name: name,
				Path: app.epPathPrefix,
				BackendBase: entities.AppDataBackendBaseSt{
					Host: app.beHost,
					Path: app.bePathPrefix,
				},
			},
		})
		if err != nil {
			return err
		}

		for _, endpoint := range app.endpoints {
			epData := &entities.EndpointDataSt{
				Method: strings.ToUpper(endpoint.ep.Method),
				Path:   endpoint.commonSuffix,
				Backend: entities.EndpointDataBackendSt{
					CustomPath: false,
					Path:       "",
				},
				JwtValidation: entities.EndpointDataJwtValidationSt{Roles: []string{}},
				IpValidation:  entities.EndpointDataIpValidationSt{AllowedIps: []string{}},
			}

			if endpoint.ep.ExtraConfig != nil {
				if endpoint.ep.ExtraConfig.AuthValidator != nil {
					if endpoint.ep.ExtraConfig.AuthValidator.Alg != "" && realm.Data.JwtConf.Alg != endpoint.ep.ExtraConfig.AuthValidator.Alg {
						realm.Data.JwtConf.Alg = endpoint.ep.ExtraConfig.AuthValidator.Alg
					}
					if realm.Data.JwtConf.JwkUrl == "" && endpoint.ep.ExtraConfig.AuthValidator.JwkUrl != "" {
						realm.Data.JwtConf.JwkUrl = endpoint.ep.ExtraConfig.AuthValidator.JwkUrl
					}
					if realm.Data.JwtConf.DisableJwkSecurity != endpoint.ep.ExtraConfig.AuthValidator.DisableJwkSecurity {
						realm.Data.JwtConf.DisableJwkSecurity = endpoint.ep.ExtraConfig.AuthValidator.DisableJwkSecurity
					}
					if realm.Data.JwtConf.RolesKey == "" && endpoint.ep.ExtraConfig.AuthValidator.RolesKey != "" {
						realm.Data.JwtConf.RolesKey = endpoint.ep.ExtraConfig.AuthValidator.RolesKey
					}
					if realm.Data.JwtConf.RolesKeyIsNested != endpoint.ep.ExtraConfig.AuthValidator.RolesKeyIsNested {
						realm.Data.JwtConf.RolesKeyIsNested = endpoint.ep.ExtraConfig.AuthValidator.RolesKeyIsNested
					}

					epData.JwtValidation.Enabled = true
					if endpoint.ep.ExtraConfig.AuthValidator.Roles != nil {
						epData.JwtValidation.Roles = endpoint.ep.ExtraConfig.AuthValidator.Roles
					}
				}

				if len(endpoint.ep.ExtraConfig.ValidationCel) == 1 {
					vCel := endpoint.ep.ExtraConfig.ValidationCel[0]
					if strings.HasPrefix(vCel.CheckExpr, "req_headers['X-Real-Ip'][0] in ") {
						ips := ipRegexp.FindAllString(vCel.CheckExpr[31:], -1)
						if len(ips) > 0 {
							epData.IpValidation.Enabled = true
							epData.IpValidation.AllowedIps = ips
						}
					}
				}
			}

			_, err := c.r.Endpoint.Create(ctx, &entities.EndpointCUSt{
				AppId:  &appId,
				Active: &cns.True,
				Data:   epData,
			})
			if err != nil {
				return err
			}
		}
	}

	// update realm

	realm.Data.Timeout = cfg.Timeout
	realm.Data.ReadHeaderTimeout = cfg.ReadHeaderTimeout
	realm.Data.ReadTimeout = cfg.ReadTimeout

	if cfg.ExtraConfig != nil {
		if cfg.ExtraConfig.SecurityCors != nil {
			realm.Data.CorsConf.Enabled = true
			realm.Data.CorsConf.AllowCredentials = cfg.ExtraConfig.SecurityCors.AllowCredentials
			realm.Data.CorsConf.MaxAge = cfg.ExtraConfig.SecurityCors.MaxAge
			if cfg.ExtraConfig.SecurityCors.AllowOrigins == nil {
				cfg.ExtraConfig.SecurityCors.AllowOrigins = []string{}
			}
			realm.Data.CorsConf.AllowOrigins = cfg.ExtraConfig.SecurityCors.AllowOrigins
			if cfg.ExtraConfig.SecurityCors.AllowMethods == nil {
				cfg.ExtraConfig.SecurityCors.AllowMethods = []string{}
			}
			realm.Data.CorsConf.AllowMethods = cfg.ExtraConfig.SecurityCors.AllowMethods
			if cfg.ExtraConfig.SecurityCors.AllowHeaders == nil {
				cfg.ExtraConfig.SecurityCors.AllowHeaders = []string{}
			}
			realm.Data.CorsConf.AllowHeaders = cfg.ExtraConfig.SecurityCors.AllowHeaders
		}
	}
	err = c.Update(ctx, realm.Id, &entities.RealmCUSt{Data: &realm.Data})
	if err != nil {
		return err
	}

	return nil
}

func (c *Realm) Deploy(ctx context.Context, id string) error {
	realm, err := c.Get(ctx, id, true)
	if err != nil {
		return err
	}

	if realm.Data.DeployConf.Url != "" {
		if strings.HasPrefix(realm.Data.DeployConf.Url, "http") { // webhook
			method := realm.Data.DeployConf.Method
			if method == "" {
				method = "GET"
			}

			hClient := httpclient.New(c.r.lg, &httpc.OptionsSt{
				Client: &http.Client{
					Timeout:   15 * time.Second,
					Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
				},
				Method:    method,
				LogPrefix: "Deploy webhook:",
			})

			_, err = hClient.Send(&httpc.OptionsSt{
				Uri: realm.Data.DeployConf.Url,
			})
			if err != nil {
				return errs.FailToSendDeployWebhook
			}
		} else { // try to restart deployment in k8s, (url == deployment-name)
			bgCtx := context.Background()

			config, err := rest.InClusterConfig() // Использовать внутри кластера
			if err != nil {
				return fmt.Errorf("failed to load Kubernetes config: %v", err)
			}

			clientSet, err := kubernetes.NewForConfig(config) // Инициализация клиента
			if err != nil {
				return fmt.Errorf("failed to create Kubernetes client: %v", err)
			}

			namespace, resourceName := parseK8sResourceName(realm.Data.DeployConf.Url)
			if namespace == "" {
				namespace = "default"
			}

			if c.k8sRestartResourceType == "daemonset" {
				err = c.restartDaemonSet(bgCtx, clientSet, namespace, resourceName)
				if err != nil {
					return fmt.Errorf("failed to restart daemonset: %w", err)
				}
			} else if c.k8sRestartResourceType == "deployment" {
				err = c.restartDeployment(bgCtx, clientSet, namespace, resourceName)
				if err != nil {
					return fmt.Errorf("failed to restart deployment: %w", err)
				}
			} else {
				return fmt.Errorf("unknown k8s-resource type: %s", c.k8sRestartResourceType)
			}
		}
	}

	return nil
}

func (c *Realm) restartDaemonSet(ctx context.Context, clientSet *kubernetes.Clientset, namespace, name string) error {
	c.r.lg.Infow("Start restart DaemonSet", "namespace", namespace, "name", name)

	resClient := clientSet.AppsV1().DaemonSets(namespace)
	res, err := resClient.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to find DaemonSets: %v", err)
	}

	if res.Spec.Template.ObjectMeta.Annotations == nil {
		res.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
	}
	res.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)

	_, err = resClient.Update(ctx, res, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update deployment: %v", err)
	}

	return nil
}

func (c *Realm) restartDeployment(ctx context.Context, clientSet *kubernetes.Clientset, namespace, name string) error {
	c.r.lg.Infow("Start restart Deployment", "namespace", namespace, "name", name)

	resClient := clientSet.AppsV1().Deployments(namespace)
	res, err := resClient.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to find Deployment: %v", err)
	}

	if res.Spec.Template.ObjectMeta.Annotations == nil {
		res.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
	}
	res.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)

	_, err = resClient.Update(ctx, res, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update deployment: %v", err)
	}

	return nil
}

// parseK8sResourceName splits a Kubernetes resource name into its namespace and name components based on the ':' separator.
// Returns an empty string and the input if no separator is found.
func parseK8sResourceName(v string) (string, string) {
	parts := strings.Split(v, ":")
	if len(parts) == 1 {
		return "", v
	}

	return parts[0], parts[1]
}
