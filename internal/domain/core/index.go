package core

import (
	"sync"

	"github.com/rendau/dop/adapters/cache"
	"github.com/rendau/dop/adapters/db"
	"github.com/rendau/dop/adapters/logger"

	"github.com/rendau/dutchman/internal/adapters/repo"
)

type St struct {
	lg                  logger.Lite
	cache               cache.Cache
	db                  db.RDBContextTransaction
	repo                repo.Repo
	authPassword        string
	sessionRefreshToken string
	testing             bool

	wg sync.WaitGroup

	Config   *Config
	Session  *Session
	Profile  *Profile
	Realm    *Realm
	App      *App
	Endpoint *Endpoint
	Role     *Role
}

func New(
	lg logger.Lite,
	cache cache.Cache,
	db db.RDBContextTransaction,
	repo repo.Repo,
	authPassword string,
	sessionRefreshToken string,
	k8sRestartResourceType string,
	testing bool,
) *St {
	c := &St{
		lg:                  lg,
		cache:               cache,
		db:                  db,
		repo:                repo,
		authPassword:        authPassword,
		sessionRefreshToken: sessionRefreshToken,
		testing:             testing,
	}

	c.Config = NewConfig(c)
	c.Session = NewSession(c)
	c.Profile = NewProfile(c)
	c.Realm = NewRealm(c, k8sRestartResourceType)
	c.App = NewApp(c)
	c.Endpoint = NewEndpoint(c)
	c.Role = NewRole(c)

	return c
}

func (c *St) WaitJobs() {
	c.wg.Wait()
}
