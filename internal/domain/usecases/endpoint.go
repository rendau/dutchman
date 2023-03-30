package usecases

import (
	"context"

	"github.com/rendau/dutchman/internal/domain/entities"
)

func (u *St) EndpointList(ctx context.Context) ([]*entities.EndpointSt, int64, error) {
	// var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Endpoint.List(ctx)
}

func (u *St) EndpointGet(ctx context.Context, id string) (*entities.EndpointSt, error) {
	// var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Endpoint.Get(ctx, id, true)
}

func (u *St) EndpointCreate(ctx context.Context,
	obj *entities.EndpointCUSt) (string, error) {
	var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return "", err
	// }

	var result string

	err = u.db.TransactionFn(ctx, func(ctx context.Context) error {
		result, err = u.cr.Endpoint.Create(ctx, obj)
		return err
	})

	return result, err
}

func (u *St) EndpointUpdate(ctx context.Context,
	id string, obj *entities.EndpointCUSt) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Endpoint.Update(ctx, id, obj)
	})
}

func (u *St) EndpointDelete(ctx context.Context,
	id string) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Endpoint.Delete(ctx, id)
	})
}
