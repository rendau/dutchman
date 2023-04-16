package usecases

import (
	"context"

	"github.com/rendau/dutchman/internal/domain/entities"
)

func (u *St) AppList(ctx context.Context,
	pars *entities.AppListParsSt) ([]*entities.AppSt, int64, error) {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return nil, 0, err
	}

	// if err = dopTools.RequirePageSize(pars.ListParams, cns.MaxPageSize); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.App.List(ctx, pars)
}

func (u *St) AppGet(ctx context.Context, id string) (*entities.AppSt, error) {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return nil, err
	}

	return u.cr.App.Get(ctx, id, true)
}

func (u *St) AppCreate(ctx context.Context,
	obj *entities.AppCUSt) (string, error) {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return "", err
	}

	var result string

	err = u.db.TransactionFn(ctx, func(ctx context.Context) error {
		result, err = u.cr.App.Create(ctx, obj)
		return err
	})

	return result, err
}

func (u *St) AppUpdate(ctx context.Context,
	id string, obj *entities.AppCUSt) error {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return err
	}

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.App.Update(ctx, id, obj)
	})
}

func (u *St) AppDelete(ctx context.Context,
	id string) error {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return err
	}

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.App.Delete(ctx, id)
	})
}

func (u *St) AppSyncRoles(ctx context.Context,
	id string) {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return
	}

	_ = u.db.TransactionFn(ctx, func(ctx context.Context) error {
		u.cr.App.SyncRoles(ctx, id)
		return nil
	})
}
