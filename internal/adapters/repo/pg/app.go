package pg

import (
	"context"
	"errors"

	"github.com/rendau/dop/adapters/db"
	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
)

func (d *St) AppGet(ctx context.Context, id string) (*entities.AppSt, error) {
	result := &entities.AppSt{}

	err := d.HfGet(ctx, db.RDBGetOptions{
		Dst:    result,
		Tables: []string{"app"},
		Conds:  []string{"id = ${id}"},
		Args:   map[string]any{"id": id},
	})
	if errors.Is(err, dopErrs.NoRows) {
		result = nil
		err = nil
	}

	return result, err
}

func (d *St) AppList(ctx context.Context, pars *entities.AppListParsSt) ([]*entities.AppSt, int64, error) {
	conds, args := d.appGetConds(ctx, pars)

	result := make([]*entities.AppSt, 0, 100)

	tCount, err := d.HfList(ctx, db.RDBListOptions{
		Dst:    &result,
		Tables: []string{`app`},
		LPars:  pars.ListParams,
		Conds:  conds,
		Args:   args,
		AllowedSorts: map[string]string{
			"default": "data ->> 'name', id",
		},
	})

	return result, tCount, err
}

func (d *St) appGetConds(ctx context.Context, pars *entities.AppListParsSt) ([]string, map[string]any) {
	conds := make([]string, 0)
	args := map[string]any{}

	// filter
	if pars.RealmId != nil {
		conds = append(conds, `realm_id = ${realm_id}`)
		args["realm_id"] = *pars.RealmId
	}
	if pars.Active != nil {
		conds = append(conds, `active = ${active}`)
		args["active"] = *pars.Active
	}

	return conds, args
}

func (d *St) AppIdExists(ctx context.Context, id string) (bool, error) {
	var cnt int

	err := d.DbQueryRow(ctx, `
        select count(*)
        from app
        where id = $1
    `, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) AppCreate(ctx context.Context, obj *entities.AppCUSt) (string, error) {
	var result string

	err := d.HfCreate(ctx, db.RDBCreateOptions{
		Table:  "app",
		Obj:    obj,
		RetCol: "id",
		RetV:   &result,
	})

	return result, err
}

func (d *St) AppUpdate(ctx context.Context, id string, obj *entities.AppCUSt) error {
	return d.HfUpdate(ctx, db.RDBUpdateOptions{
		Table: "app",
		Obj:   obj,
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}

func (d *St) AppDelete(ctx context.Context, id string) error {
	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: "app",
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}

func (d *St) AppDeleteMany(ctx context.Context, pars *entities.AppListParsSt) error {
	conds, args := d.appGetConds(ctx, pars)

	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: "app",
		Conds: conds,
		Args:  args,
	})
}
