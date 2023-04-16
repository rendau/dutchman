package core

import (
	"context"

	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dop/dopTools"
	"github.com/rendau/dutchman/internal/cns"
	"github.com/rendau/dutchman/internal/domain/entities"
)

type App struct {
	r *St
}

func NewApp(r *St) *App {
	return &App{r: r}
}

func (c *App) ValidateCU(ctx context.Context, obj *entities.AppCUSt, id string) error {
	// forCreate := id == ""

	return nil
}

func (c *App) List(ctx context.Context, pars *entities.AppListParsSt) ([]*entities.AppSt, int64, error) {
	items, tCount, err := c.r.repo.AppList(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	return items, tCount, nil
}

func (c *App) Get(ctx context.Context, id string, errNE bool) (*entities.AppSt, error) {
	result, err := c.r.repo.AppGet(ctx, id)
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

func (c *App) IdExists(ctx context.Context, id string) (bool, error) {
	return c.r.repo.AppIdExists(ctx, id)
}

func (c *App) Create(ctx context.Context, obj *entities.AppCUSt) (string, error) {
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	// create
	result, err := c.r.repo.AppCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	c.SyncRoles(ctx, result)

	return result, nil
}

func (c *App) Update(ctx context.Context, id string, obj *entities.AppCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	err = c.r.repo.AppUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	c.SyncRoles(ctx, id)

	return nil
}

func (c *App) Delete(ctx context.Context, id string) error {
	return c.r.repo.AppDelete(ctx, id)
}

func (c *App) DeleteMany(ctx context.Context, pars *entities.AppListParsSt) error {
	return c.r.repo.AppDeleteMany(ctx, pars)
}

func (c *App) SyncRoles(ctx context.Context, id string) {
	app, err := c.Get(ctx, id, true)
	if err != nil {
		return
	}

	dbRoles, _, err := c.r.Role.List(ctx, &entities.RoleListParsSt{
		AppId:     &id,
		IsFetched: &cns.True,
	})
	if err != nil {
		return
	}

	if app.Data.RemoteRoles.Url == "" {
		if len(dbRoles) > 0 {
			// delete
			for _, dbRole := range dbRoles {
				err = c.r.Role.Delete(ctx, dbRole.Id)
				if err != nil {
					return
				}
			}
		}

		return
	}

	remoteItems := c.r.Role.FetchRemoteUri(app.Data.RemoteRoles.Url, app.Data.RemoteRoles.JsonPath)
	if err != nil {
		return
	}

	var found bool

	for _, dbRole := range dbRoles {
		found = false

		for _, freshRole := range remoteItems {
			if freshRole.Code == dbRole.Code {
				found = true
				break
			}
		}

		if !found {
			// delete
			err = c.r.Role.Delete(ctx, dbRole.Id)
			if err != nil {
				return
			}
		}
	}

	for _, freshRole := range remoteItems {
		if freshRole.Code == "" {
			continue
		}

		found = false

		for _, dbRole := range dbRoles {
			if freshRole.Code == dbRole.Code {
				found = true

				if freshRole.Dsc != dbRole.Dsc {
					// update
					err = c.r.Role.Update(ctx, dbRole.Id, &entities.RoleCUSt{
						Dsc:       &freshRole.Dsc,
						IsFetched: &cns.True,
					})
					if err != nil {
						return
					}
				}

				break
			}
		}

		if !found {
			// create
			_, err = c.r.Role.Create(ctx, &entities.RoleCUSt{
				RealmId:   &app.RealmId,
				DbAppId:   dopTools.NewPtr(&id),
				IsFetched: &cns.True,
				Code:      &freshRole.Code,
				Dsc:       &freshRole.Dsc,
			})
			if err != nil {
				return
			}
		}
	}
}

func (c *App) FetchRoles(ctx context.Context, id string) []*entities.RoleFetchRemoteRepItemSt {
	app, err := c.Get(ctx, id, true)
	if err != nil {
		return []*entities.RoleFetchRemoteRepItemSt{}
	}

	return c.r.Role.FetchRemoteUri(app.Data.RemoteRoles.Url, app.Data.RemoteRoles.JsonPath)
}
