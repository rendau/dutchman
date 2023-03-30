package pg

import (
	"context"
	"errors"

	"github.com/rendau/dop/adapters/db"
	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
)

func (d *St) RealmGet(ctx context.Context, id string) (*entities.RealmSt, error) {
	result := &entities.RealmSt{}

	err := d.HfGet(ctx, db.RDBGetOptions{
		Dst:    result,
		Tables: []string{"realm"},
		Conds:  []string{"id = ${id}"},
		Args:   map[string]any{"id": id},
	})
	if errors.Is(err, dopErrs.NoRows) {
		result = nil
		err = nil
	}

	return result, err
}

func (d *St) RealmList(ctx context.Context) ([]*entities.RealmSt, int64, error) {
	conds := make([]string, 0)
	args := map[string]any{}

	result := make([]*entities.RealmSt, 0, 100)

	tCount, err := d.HfList(ctx, db.RDBListOptions{
		Dst:    &result,
		Tables: []string{`realm t`},

		Conds: conds,
		Args:  args,
		AllowedSorts: map[string]string{
			"default": "t.id",
		},
	})

	return result, tCount, err
}

func (d *St) RealmIdExists(ctx context.Context, id string) (bool, error) {
	var cnt int

	err := d.DbQueryRow(ctx, `
        select count(*)
        from realm
        where id = $1
    `, id).Scan(&cnt)

	return cnt > 0, err
}

func (d *St) RealmCreate(ctx context.Context, obj *entities.RealmCUSt) (string, error) {
	var result string

	err := d.HfCreate(ctx, db.RDBCreateOptions{
		Table:  "realm",
		Obj:    obj,
		RetCol: "id",
		RetV:   &result,
	})

	return result, err
}

func (d *St) RealmUpdate(ctx context.Context, id string, obj *entities.RealmCUSt) error {
	return d.HfUpdate(ctx, db.RDBUpdateOptions{
		Table: "realm",
		Obj:   obj,
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}

func (d *St) RealmDelete(ctx context.Context, id string) error {
	return d.HfDelete(ctx, db.RDBDeleteOptions{
		Table: "realm",
		Conds: []string{"id = ${cond_id}"},
		Args:  map[string]any{"cond_id": id},
	})
}
