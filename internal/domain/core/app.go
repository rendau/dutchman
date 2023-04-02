package core

import (
	"context"

	"github.com/rendau/dop/dopErrs"
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

	return nil
}

func (c *App) Delete(ctx context.Context, id string) error {
	return c.r.repo.AppDelete(ctx, id)
}

// func (c *App) FetchRoles(ctx context.Context, id string) []*entities.RoleRemoteRepItemSt {
// 	app, err := c.Get(ctx, id, true)
// 	if err != nil {
// 		return []*entities.RoleRemoteRepItemSt{}
// 	}
//
// 	return c.r.Role.FetchRemoteUri(app, app.RemotePath)
// }
