package core

import (
	"context"

	"github.com/rendau/dutchman/internal/domain/entities"
)

type Dic struct {
	r *St
}

func NewDic(r *St) *Dic {
	return &Dic{r: r}
}

func (c *Dic) Get(ctx context.Context) (*entities.DicSt, error) {
	// var err error

	data := &entities.DicSt{}

	return data, nil
}
