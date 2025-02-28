package usecases

import (
	"context"

	"github.com/rendau/dutchman/internal/domain/entities"
)

func (u *St) RealmList(ctx context.Context,
	pars *entities.RealmListParsSt) ([]*entities.RealmSt, int64, error) {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return nil, 0, err
	}

	// if err = dopTools.RequirePageSize(pars.ListParams, cns.MaxPageSize); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Realm.List(ctx, pars)
}

func (u *St) RealmGet(ctx context.Context, id string) (*entities.RealmSt, error) {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return nil, err
	}

	return u.cr.Realm.Get(ctx, id, true)
}

func (u *St) RealmCreate(ctx context.Context,
	obj *entities.RealmCUSt) (string, error) {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return "", err
	}

	var result string

	err = u.db.TransactionFn(ctx, func(ctx context.Context) error {
		result, err = u.cr.Realm.Create(ctx, obj)
		return err
	})

	return result, err
}

func (u *St) RealmUpdate(ctx context.Context,
	id string, obj *entities.RealmCUSt) error {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return err
	}

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Realm.Update(ctx, id, obj)
	})
}

func (u *St) RealmDelete(ctx context.Context,
	id string) error {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return err
	}

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Realm.Delete(ctx, id)
	})
}

func (u *St) RealmExportConf(ctx context.Context,
	id string) (*entities.KrakendSt, error) {
	return u.cr.Realm.GenerateConf(ctx, id)
}

func (u *St) RealmPreviewConf(ctx context.Context,
	id string) (*entities.KrakendSt, error) {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return nil, err
	}

	return u.cr.Realm.GenerateConf(ctx, id)
}

func (u *St) RealmDeploy(ctx context.Context,
	id string) error {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return err
	}

	return u.cr.Realm.Deploy(ctx, id)
}

func (u *St) RealmImportConf(ctx context.Context,
	id string, cfg *entities.KrakendSt) error {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return err
	}

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Realm.ImportConf(ctx, id, cfg)
	})
}
