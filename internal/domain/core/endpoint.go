package core

import (
	"context"

	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
)

type Endpoint struct {
	r *St
}

func NewEndpoint(r *St) *Endpoint {
	return &Endpoint{r: r}
}

func (c *Endpoint) ValidateCU(ctx context.Context, obj *entities.EndpointCUSt, id string) error {
	// forCreate := id == ""

	return nil
}

func (c *Endpoint) List(ctx context.Context, pars *entities.EndpointListParsSt) ([]*entities.EndpointSt, int64, error) {
	items, tCount, err := c.r.repo.EndpointList(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	return items, tCount, nil
}

func (c *Endpoint) Get(ctx context.Context, id string, pars *entities.EndpointGetParsSt, errNE bool) (*entities.EndpointSt, error) {
	result, err := c.r.repo.EndpointGet(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		if errNE {
			return nil, dopErrs.ObjectNotFound
		}
		return nil, nil
	}

	if pars.WithApp {
		result.App, err = c.r.App.Get(ctx, result.AppId, true)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (c *Endpoint) IdExists(ctx context.Context, id string) (bool, error) {
	return c.r.repo.EndpointIdExists(ctx, id)
}

func (c *Endpoint) Create(ctx context.Context, obj *entities.EndpointCUSt) (string, error) {
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	// create
	result, err := c.r.repo.EndpointCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Endpoint) Update(ctx context.Context, id string, obj *entities.EndpointCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	err = c.r.repo.EndpointUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Endpoint) Delete(ctx context.Context, id string) error {
	return c.r.repo.EndpointDelete(ctx, id)
}
