package pg

import (
	"context"
	"errors"

	"github.com/rendau/dop/adapters/db"
	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
)

func (d *St) EndpointGet(ctx context.Context, id string) (*entities.EndpointSt, error) {
	result := &entities.EndpointSt{}

	err := d.HfGet(ctx, db.RDBGetOptions{
		Dst:    result,
		Tables: []string{"endpoint"},
		Conds:  []string{"id = ${id}"},
		Args:   map[string]any{"id": id},
	})
	if errors.Is(err, dopErrs.NoRows) {
		result = nil
		err = nil
	}

	return result, err
}

func (d *St) EndpointList(ctx context.Context, pars *entities.EndpointListParsSt) ([]*entities.EndpointSt, int64, error) {
	conds := make([]string, 0)
	args := map[string]any{}

	// filter
	if pars.AppId != nil {
		conds = append(conds, `t.app_id = ${app_id}`)
		args["app_id"] = *pars.AppId
	}
	if pars.Active != nil {
		conds = append(conds, `t.active = ${active}`)
		args["active"] = *pars.Active
	}

	result := make([]*entities.EndpointSt, 0, 100)

	tCount, err := d.HfList(ctx, db.RDBListOptions{
		Dst:    &result,
		Tables: []string{`endpoint t`},
		LPars:  pars.ListParams,
		Conds:  conds,
		Args:   args,
		AllowedSorts: map[string]string{
			"default": "t.id",
		},
	})

	return result, tCount, err
}

func (d *St) EndpointIdExists(ctx context.Context, id string) (bool, error) {
	var cnt int

	err := d.DbQueryRow(ctx, `
        select count(*)
        from endpoint
        where id = $1
    `, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) EndpointCreate(ctx context.Context, obj *entities.EndpointCUSt) (string, error) {
	var result string

	err := d.HfCreate(ctx, db.RDBCreateOptions{
		Table:  "endpoint",
		Obj:    obj,
		RetCol: "id",
		RetV:   &result,
	})

	return result, err
}

func (d *St) EndpointUpdate(ctx context.Context, id string, obj *entities.EndpointCUSt) error {
	return d.HfUpdate(ctx, db.RDBUpdateOptions{
		Table: "endpoint",
		Obj:   obj,
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}

func (d *St) EndpointDelete(ctx context.Context, id string) error {
	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: "endpoint",
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}
