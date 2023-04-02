package pg

import (
	"context"
	"errors"

	"github.com/rendau/dop/adapters/db"
	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
)

func (d *St) PermGet(ctx context.Context, id string) (*entities.PermSt, error) {
	result := &entities.PermSt{}

	err := d.HfGet(ctx, db.RDBGetOptions{
		Dst:    result,
		Tables: []string{"perm"},
		Conds:  []string{"id = ${id}"},
		Args:   map[string]any{"id": id},
	})
	if errors.Is(err, dopErrs.NoRows) {
		result = nil
		err = nil
	}

	return result, err
}

func (d *St) PermList(ctx context.Context, pars *entities.PermListParsSt) ([]*entities.PermSt, int64, error) {
	conds := make([]string, 0)
	args := map[string]any{}

	// filter
	if pars.AppId != nil {
		conds = append(conds, `t.app_id = ${app_id}`)
		args["app_id"] = *pars.AppId
	}

	result := make([]*entities.PermSt, 0, 100)

	tCount, err := d.HfList(ctx, db.RDBListOptions{
		Dst:    &result,
		Tables: []string{`perm t`},
		LPars:  pars.ListParams,
		Conds:  conds,
		Args:   args,
		AllowedSorts: map[string]string{
			"default": "t.id",
		},
	})

	return result, tCount, err
}

func (d *St) PermIdExists(ctx context.Context, id string) (bool, error) {
	var cnt int

	err := d.DbQueryRow(ctx, `
        select count(*)
        from perm
        where id = $1
    `, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) PermCreate(ctx context.Context, obj *entities.PermCUSt) (string, error) {
	var result string

	err := d.HfCreate(ctx, db.RDBCreateOptions{
		Table:  "perm",
		Obj:    obj,
		RetCol: "id",
		RetV:   &result,
	})

	return result, err
}

func (d *St) PermUpdate(ctx context.Context, id string, obj *entities.PermCUSt) error {
	return d.HfUpdate(ctx, db.RDBUpdateOptions{
		Table: "perm",
		Obj:   obj,
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}

func (d *St) PermDelete(ctx context.Context, id string) error {
	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: "perm",
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}
