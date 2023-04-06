package pg

import (
	"context"
	"errors"

	"github.com/rendau/dop/adapters/db"
	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
)

func (d *St) RoleGet(ctx context.Context, id string) (*entities.RoleSt, error) {
	result := &entities.RoleSt{}

	err := d.HfGet(ctx, db.RDBGetOptions{
		Dst:    result,
		Tables: []string{`"role"`},
		Conds:  []string{"id = ${id}"},
		Args:   map[string]any{"id": id},
	})
	if errors.Is(err, dopErrs.NoRows) {
		result = nil
		err = nil
	}

	return result, err
}

func (d *St) RoleList(ctx context.Context, pars *entities.RoleListParsSt) ([]*entities.RoleSt, int64, error) {
	conds := make([]string, 0)
	args := map[string]any{}

	// filter
	if pars.AppId != nil {
		if *pars.AppId == "-" {
			conds = append(conds, `t.app_id is null`)
		} else {
			conds = append(conds, `t.app_id = ${app_id}`)
			args["app_id"] = *pars.AppId
		}
	}
	if pars.AppIdOrNull != nil {
		conds = append(conds, `(t.app_id = ${app_id_or_null} or t.app_id is null)`)
		args["app_id_or_null"] = *pars.AppIdOrNull
	}
	if pars.IsFetched != nil {
		if *pars.IsFetched {
			conds = append(conds, `t.is_fetched`)
		} else {
			conds = append(conds, `not t.is_fetched`)
		}
	}

	result := make([]*entities.RoleSt, 0, 100)

	tCount, err := d.HfList(ctx, db.RDBListOptions{
		Dst:    &result,
		Tables: []string{`"role" t`},
		LPars:  pars.ListParams,
		Conds:  conds,
		Args:   args,
		AllowedSorts: map[string]string{
			"default": "t.app_id, t.code",
		},
	})

	return result, tCount, err
}

func (d *St) RoleIdExists(ctx context.Context, id string) (bool, error) {
	var cnt int

	err := d.DbQueryRow(ctx, `
        select count(*)
        from "role"
        where id = $1
    `, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) RoleCreate(ctx context.Context, obj *entities.RoleCUSt) (string, error) {
	var result string

	err := d.HfCreate(ctx, db.RDBCreateOptions{
		Table:  `"role"`,
		Obj:    obj,
		RetCol: "id",
		RetV:   &result,
	})

	return result, err
}

func (d *St) RoleUpdate(ctx context.Context, id string, obj *entities.RoleCUSt) error {
	return d.HfUpdate(ctx, db.RDBUpdateOptions{
		Table: `"role"`,
		Obj:   obj,
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}

func (d *St) RoleDelete(ctx context.Context, id string) error {
	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: `"role"`,
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}
