package usecases

import (
	"context"

	"github.com/rendau/dop/dopTools"
	"github.com/rendau/dutchman/internal/cns"
	"github.com/rendau/dutchman/internal/domain/entities"
)

func (u *St) RoleList(ctx context.Context,
	pars *entities.RoleListParsSt) ([]*entities.RoleSt, int64, error) {
	var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	if err = dopTools.RequirePageSize(pars.ListParams, cns.MaxPageSize); err != nil {
		return nil, 0, err
	}

	return u.cr.Role.List(ctx, pars)
}

func (u *St) RoleGet(ctx context.Context, id string) (*entities.RoleSt, error) {
	// var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Role.Get(ctx, id, true)
}

func (u *St) RoleCreate(ctx context.Context,
	obj *entities.RoleCUSt) (string, error) {
	var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return "", err
	// }

	var result string

	err = u.db.TransactionFn(ctx, func(ctx context.Context) error {
		result, err = u.cr.Role.Create(ctx, obj)
		return err
	})

	return result, err
}

func (u *St) RoleUpdate(ctx context.Context,
	id string, obj *entities.RoleCUSt) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Role.Update(ctx, id, obj)
	})
}

func (u *St) RoleDelete(ctx context.Context,
	id string) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Role.Delete(ctx, id)
	})
}
