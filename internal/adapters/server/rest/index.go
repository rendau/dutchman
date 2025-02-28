package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rendau/dop/adapters/logger"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	swagFiles "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"

	"github.com/rendau/dutchman/internal/domain/usecases"
)

type St struct {
	lg  logger.Lite
	ucs *usecases.St
}

func GetHandler(
	lg logger.Lite,
	ucs *usecases.St,
	withCors bool,
) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// middlewares
	if withCors {
		r.Use(dopHttps.MwCors())
	}
	r.Use(dopHttps.MwRecovery(lg, nil))

	// healthcheck
	r.GET("/healthcheck", func(c *gin.Context) { c.Status(http.StatusOK) })

	// doc
	r.GET("/doc/*any", ginSwag.WrapHandler(swagFiles.Handler, func(c *ginSwag.Config) {
		c.DefaultModelsExpandDepth = 0
		c.DocExpansion = "none"
	}))

	// api
	s := &St{lg: lg, ucs: ucs}

	// config
	r.GET("/config", s.hConfigGet)
	r.PUT("/config", s.hConfigUpdate)

	// profile
	r.GET("/profile", s.hProfileGet)
	r.POST("/profile/auth", s.hProfileAuth)
	r.POST("/profile/auth/token", s.hProfileAuthByRefreshToken)

	// app
	r.GET("/app", s.hAppList)
	r.POST("/app", s.hAppCreate)
	r.GET("/app/:id", s.hAppGet)
	r.PUT("/app/:id", s.hAppUpdate)
	r.DELETE("/app/:id", s.hAppDelete)
	r.POST("/app/:id/duplicate", s.hAppDuplicate)
	r.POST("/app/:id/sync_roles", s.hAppSyncRoles)

	// endpoint
	r.GET("/endpoint", s.hEndpointList)
	r.POST("/endpoint", s.hEndpointCreate)
	r.GET("/endpoint/:id", s.hEndpointGet)
	r.PUT("/endpoint/:id", s.hEndpointUpdate)
	r.DELETE("/endpoint/:id", s.hEndpointDelete)

	// realm
	r.GET("/realm", s.hRealmList)
	r.POST("/realm", s.hRealmCreate)
	r.GET("/realm/:id", s.hRealmGet)
	r.PUT("/realm/:id", s.hRealmUpdate)
	r.DELETE("/realm/:id", s.hRealmDelete)
	r.GET("/realm/:id/export", s.hRealmExportConf)
	r.GET("/realm/:id/preview_conf", s.hRealmPreviewConf)
	r.POST("/realm/:id/deploy", s.hRealmDeploy)
	r.POST("/realm/:id/import_conf", s.hRealmImportConf)

	// role
	r.GET("/role", s.hRoleList)
	r.POST("/role", s.hRoleCreate)
	r.GET("/role/:id", s.hRoleGet)
	r.PUT("/role/:id", s.hRoleUpdate)
	r.DELETE("/role/:id", s.hRoleDelete)
	r.POST("/role/fetch_remote_uri", s.hRoleFetchRemoteUri)

	return r
}

func (o *St) getRequestContext(c *gin.Context) context.Context {
	return o.ucs.SessionSetToContextByToken(nil, dopHttps.GetAuthToken(c))
}
