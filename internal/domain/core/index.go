package core

import (
	"sync"

	"github.com/rendau/dop/adapters/cache"
	"github.com/rendau/dop/adapters/db"
	"github.com/rendau/dop/adapters/logger"
	"github.com/rendau/dutchman/internal/adapters/repo"
)

type St struct {
	lg      logger.Lite
	cache   cache.Cache
	db      db.RDBContextTransaction
	repo    repo.Repo
	testing bool

	wg sync.WaitGroup

	Config  *Config
	Dic     *Dic
	Session *Session
}

func New(
	lg logger.Lite,
	cache cache.Cache,
	db db.RDBContextTransaction,
	repo repo.Repo,
	testing bool,
) *St {
	c := &St{
		lg:      lg,
		cache:   cache,
		db:      db,
		repo:    repo,
		testing: testing,
	}

	c.Config = NewConfig(c)
	c.Dic = NewDic(c)
	c.Session = NewSession(c)

	return c
}

func (c *St) WaitJobs() {
	c.wg.Wait()
}
