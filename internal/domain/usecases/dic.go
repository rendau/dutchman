package usecases

import (
	"context"

	"github.com/rendau/dutchman/internal/domain/entities"
)

func (u *St) DicGet(ctx context.Context) (*entities.DicSt, error) {
	return u.cr.Dic.Get(ctx)
}
