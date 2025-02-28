package cmd

import (
	"os"
	"time"

	dopCache "github.com/rendau/dop/adapters/cache"
	dopCacheMem "github.com/rendau/dop/adapters/cache/mem"
	dopCacheRedis "github.com/rendau/dop/adapters/cache/redis"
	dopDbPg "github.com/rendau/dop/adapters/db/pg"
	dopLoggerZap "github.com/rendau/dop/adapters/logger/zap"
	dopServerHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTools"

	"github.com/rendau/dutchman/docs"
	"github.com/rendau/dutchman/internal/adapters/repo/pg"
	"github.com/rendau/dutchman/internal/adapters/server/rest"
	"github.com/rendau/dutchman/internal/domain/core"
	"github.com/rendau/dutchman/internal/domain/usecases"
)

func Execute() {
	var err error

	app := struct {
		lg         *dopLoggerZap.St
		cache      dopCache.Cache
		db         *dopDbPg.St
		repo       *pg.St
		core       *core.St
		ucs        *usecases.St
		restApi    *rest.St
		restApiSrv *dopServerHttps.St
	}{}

	confLoad()

	app.lg = dopLoggerZap.New(conf.LogLevel, conf.Debug)

	if conf.RedisUrl == "" {
		app.cache = dopCacheMem.New()
	} else {
		app.cache = dopCacheRedis.New(
			app.lg,
			conf.RedisUrl,
			conf.RedisPsw,
			conf.RedisDb,
			conf.RedisKeyPrefix,
		)
	}

	app.db, err = dopDbPg.New(conf.Debug, app.lg, dopDbPg.OptionsSt{
		Dsn: conf.PgDsn,
	})
	if err != nil {
		app.lg.Fatal(err)
	}

	app.repo = pg.New(app.db, app.lg)

	app.ucs = usecases.New(app.lg, app.db)

	app.core = core.New(
		app.lg,
		app.cache,
		app.db,
		app.repo,
		conf.AuthPassword,
		conf.SessionRefreshToken,
		false,
	)

	app.ucs.SetCore(app.core)

	docs.SwaggerInfo.Host = conf.SwagHost
	docs.SwaggerInfo.BasePath = conf.SwagBasePath
	docs.SwaggerInfo.Schemes = []string{conf.SwagSchema}
	docs.SwaggerInfo.Title = "Dutchman service"

	// START

	app.lg.Infow("Starting...")

	app.restApiSrv = dopServerHttps.Start(
		conf.HttpListen,
		rest.GetHandler(
			app.lg,
			app.ucs,
			conf.HttpCors,
		),
		app.lg,
	)

	var exitCode int

	select {
	case <-dopTools.StopSignal():
	case <-app.restApiSrv.Wait():
		exitCode = 1
	}

	// STOP

	app.lg.Infow("Shutting down...")

	if !app.restApiSrv.Shutdown(20 * time.Second) {
		exitCode = 1
	}

	app.lg.Infow("Wait routines...")

	app.core.WaitJobs()

	app.lg.Infow("Exit")

	os.Exit(exitCode)
}
