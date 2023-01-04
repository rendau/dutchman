package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rendau/dop/adapters/logger"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dutchman/internal/domain/usecases"
	swagFiles "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"
)

type St struct {
	lg  logger.Lite
	ucs *usecases.St
}

func GetHandler(
	lg logger.Lite,
	ucs *usecases.St,
	frontDir string,
	frontConfig string,
	withCors bool,
) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// middlewares
	setMiddlewares(r, lg, withCors)

	// doc
	addDocRoutes(r.Group("/doc"))

	// api
	addApiRoutes(r, &St{lg: lg, ucs: ucs})

	// healthcheck
	addHealthcheckRoute(r)

	return r
}

func setMiddlewares(r *gin.Engine, lg logger.WarnAndError, withCors bool) {
	r.Use(dopHttps.MwRecovery(lg, nil))
	if withCors {
		r.Use(dopHttps.MwCors())
	}
}

func addDocRoutes(r *gin.RouterGroup) {
	r.GET("/*any", ginSwag.WrapHandler(swagFiles.Handler, func(c *ginSwag.Config) {
		c.DefaultModelsExpandDepth = 0
		c.DocExpansion = "none"
	}))
}

func addApiRoutes(r *gin.Engine, s *St) {
	// dic
	r.GET("/dic", s.hDicGet)

	// config
	r.GET("/config", s.hConfigGet)
	r.PUT("/config", s.hConfigUpdate)
}

func addHealthcheckRoute(r *gin.Engine) {
	r.GET("/healthcheck", func(c *gin.Context) { c.Status(http.StatusOK) })
}

func (o *St) getRequestContext(c *gin.Context) context.Context {
	return o.ucs.SessionSetToContextByToken(nil, dopHttps.GetAuthToken(c))
}
