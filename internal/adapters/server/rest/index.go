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
	withCors bool,
) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// middlewares
	r.Use(dopHttps.MwRecovery(lg, nil))
	if withCors {
		r.Use(dopHttps.MwCors())
	}

	// healthcheck
	r.GET("/healthcheck", func(c *gin.Context) { c.Status(http.StatusOK) })

	// doc
	r.GET("/doc/*any", ginSwag.WrapHandler(swagFiles.Handler, func(c *ginSwag.Config) {
		c.DefaultModelsExpandDepth = 0
		c.DocExpansion = "none"
	}))

	// api
	s := &St{lg: lg, ucs: ucs}

	// dic
	r.GET("/dic", s.hDicGet)

	// config
	r.GET("/config", s.hConfigGet)
	r.PUT("/config", s.hConfigUpdate)

	// profile
	r.GET("/profile", s.hProfileGet)
	r.POST("/profile/auth", s.hProfileAuth)
	r.POST("/profile/auth/token", s.hProfileAuthByRefreshToken)

	// data
	r.GET("/data", s.hDataList)
	r.POST("/data", s.hDataCreate)
	r.GET("/data/:id", s.hDataGet)
	r.PUT("/data/:id", s.hDataUpdate)
	r.DELETE("/data/:id", s.hDataDelete)
	r.POST("/data/deploy", s.hDataDeploy)

	return r
}

func (o *St) getRequestContext(c *gin.Context) context.Context {
	return o.ucs.SessionSetToContextByToken(nil, dopHttps.GetAuthToken(c))
}
