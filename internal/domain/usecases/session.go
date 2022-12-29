package usecases

import (
	"context"

	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
)

func (u *St) SessionGetFromToken(token string) *entities.Session {
	return u.cr.Session.GetFromToken(token)
}

func (u *St) SessionRequireAuth(ses *entities.Session) error {
	if ses.Id == 0 {
		return dopErrs.NotAuthorized
	}

	return nil
}

func (u *St) SessionSetToContext(ctx context.Context, ses *entities.Session) context.Context {
	return u.cr.Session.SetToContext(ctx, ses)
}

func (u *St) SessionSetToContextByToken(ctx context.Context, token string) context.Context {
	return u.cr.Session.SetToContext(ctx, u.SessionGetFromToken(token))
}

func (u *St) SessionGetFromContext(ctx context.Context) *entities.Session {
	return u.cr.Session.GetFromContext(ctx)
}
