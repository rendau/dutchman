package core

import (
	"context"
	"encoding/json"

	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
)

type Role struct {
	r *St
}

func NewRole(r *St) *Role {
	return &Role{r: r}
}

func (c *Role) ValidateCU(ctx context.Context, obj *entities.RoleCUSt, id string) error {
	// forCreate := id == ""

	return nil
}

func (c *Role) List(ctx context.Context, pars *entities.RoleListParsSt) ([]*entities.RoleSt, int64, error) {
	items, tCount, err := c.r.repo.RoleList(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	return items, tCount, nil
}

func (c *Role) Get(ctx context.Context, id string, errNE bool) (*entities.RoleSt, error) {
	result, err := c.r.repo.RoleGet(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		if errNE {
			return nil, dopErrs.ObjectNotFound
		}
		return nil, nil
	}

	return result, nil
}

func (c *Role) IdExists(ctx context.Context, id string) (bool, error) {
	return c.r.repo.RoleIdExists(ctx, id)
}

func (c *Role) Create(ctx context.Context, obj *entities.RoleCUSt) (string, error) {
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	// create
	result, err := c.r.repo.RoleCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Role) Update(ctx context.Context, id string, obj *entities.RoleCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	err = c.r.repo.RoleUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Role) Delete(ctx context.Context, id string) error {
	return c.r.repo.RoleDelete(ctx, id)
}

func (c *Role) parseRemoteJson(src []byte, path []string) ([]*entities.RoleRemoteRepItemSt, error) {
	if len(src) == 0 {
		return []*entities.RoleRemoteRepItemSt{}, nil
	}

	if len(path) == 0 {
		result := make([]*entities.RoleRemoteRepItemSt, 0)

		err := json.Unmarshal(src, &result)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	obj := map[string]json.RawMessage{}

	err := json.Unmarshal(src, &obj)
	if err != nil {
		return nil, err
	}

	return c.parseRemoteJson(obj[path[0]], path[1:])
}
