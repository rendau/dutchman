package repo

import (
	"context"

	"github.com/rendau/dutchman/internal/domain/entities"
)

type Repo interface {
	// config
	ConfigGet(ctx context.Context) (*entities.ConfigSt, error)
	ConfigSet(ctx context.Context, config *entities.ConfigSt) error

	// data
	DataGet(ctx context.Context, id string) (*entities.DataSt, error)
	DataList(ctx context.Context) ([]*entities.DataListSt, int64, error)
	DataIdExists(ctx context.Context, id string) (bool, error)
	DataCreate(ctx context.Context, obj *entities.DataCUSt) (string, error)
	DataUpdate(ctx context.Context, id string, obj *entities.DataCUSt) error
	DataDelete(ctx context.Context, id string) error
}
