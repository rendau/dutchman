package repo

import (
	"context"

	"github.com/rendau/dutchman/internal/domain/entities"
)

type Repo interface {
	// config
	ConfigGet(ctx context.Context) (*entities.ConfigSt, error)
	ConfigSet(ctx context.Context, config *entities.ConfigSt) error

	// app
	AppGet(ctx context.Context, id string) (*entities.AppSt, error)
	AppList(ctx context.Context, pars *entities.AppListParsSt) ([]*entities.AppSt, int64, error)
	AppIdExists(ctx context.Context, id string) (bool, error)
	AppCreate(ctx context.Context, obj *entities.AppCUSt) (string, error)
	AppUpdate(ctx context.Context, id string, obj *entities.AppCUSt) error
	AppDelete(ctx context.Context, id string) error

	// endpoint
	EndpointGet(ctx context.Context, id string) (*entities.EndpointSt, error)
	EndpointList(ctx context.Context, pars *entities.EndpointListParsSt) ([]*entities.EndpointSt, int64, error)
	EndpointIdExists(ctx context.Context, id string) (bool, error)
	EndpointCreate(ctx context.Context, obj *entities.EndpointCUSt) (string, error)
	EndpointUpdate(ctx context.Context, id string, obj *entities.EndpointCUSt) error
	EndpointDelete(ctx context.Context, id string) error

	// realm
	RealmGet(ctx context.Context, id string) (*entities.RealmSt, error)
	RealmList(ctx context.Context, pars *entities.RealmListParsSt) ([]*entities.RealmSt, int64, error)
	RealmIdExists(ctx context.Context, id string) (bool, error)
	RealmCreate(ctx context.Context, obj *entities.RealmCUSt) (string, error)
	RealmUpdate(ctx context.Context, id string, obj *entities.RealmCUSt) error
	RealmDelete(ctx context.Context, id string) error

	// role
	RoleGet(ctx context.Context, id string) (*entities.RoleSt, error)
	RoleList(ctx context.Context, pars *entities.RoleListParsSt) ([]*entities.RoleSt, int64, error)
	RoleIdExists(ctx context.Context, id string) (bool, error)
	RoleCreate(ctx context.Context, obj *entities.RoleCUSt) (string, error)
	RoleUpdate(ctx context.Context, id string, obj *entities.RoleCUSt) error
	RoleDelete(ctx context.Context, id string) error
}
