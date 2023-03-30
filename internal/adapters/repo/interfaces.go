package repo

import (
	"context"

	"github.com/rendau/dutchman/internal/domain/entities"
)

type Repo interface {
	// config
	ConfigGet(ctx context.Context) (*entities.ConfigSt, error)
	ConfigSet(ctx context.Context, config *entities.ConfigSt) error

	// realm
	RealmGet(ctx context.Context, id string) (*entities.RealmSt, error)
	RealmList(ctx context.Context) ([]*entities.RealmSt, int64, error)
	RealmIdExists(ctx context.Context, id string) (bool, error)
	RealmCreate(ctx context.Context, obj *entities.RealmCUSt) (string, error)
	RealmUpdate(ctx context.Context, id string, obj *entities.RealmCUSt) error
	RealmDelete(ctx context.Context, id string) error
}
