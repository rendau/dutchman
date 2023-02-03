package usecases

import (
	"context"

	"github.com/rendau/dutchman/internal/domain/entities"
)

func (u *St) ProfileAuth(ctx context.Context,
	obj *entities.ProfileAuthReqSt) (*entities.ProfileAuthRepSt, error) {
	return u.cr.Profile.Auth(ctx, obj)
}

func (u *St) ProfileAuthByRefreshToken(ctx context.Context,
	obj *entities.ProfileAuthByRefreshTokenReqSt) (*entities.ProfileAuthByRefreshTokenRepSt, error) {
	return u.cr.Profile.AuthByRefreshToken(ctx, obj)
}

func (u *St) ProfileGet(ctx context.Context) (*entities.ProfileSt, error) {
	var err error

	ses := u.SessionGetFromContext(ctx)

	if err = u.SessionRequireAuth(ses); err != nil {
		return nil, err
	}

	return u.cr.Profile.Get(ctx)
}
