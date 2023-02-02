package core

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/rendau/dop/adapters/client/httpc"
	"github.com/rendau/dop/adapters/client/httpc/httpclient"
	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
	"github.com/rendau/dutchman/internal/domain/errs"
)

type Data struct {
	r *St
}

func NewData(r *St) *Data {
	return &Data{r: r}
}

func (c *Data) ValidateCU(ctx context.Context, obj *entities.DataCUSt, id string) error {
	// forCreate := id == ""

	return nil
}

func (c *Data) List(ctx context.Context) ([]*entities.DataListSt, int64, error) {
	items, tCount, err := c.r.repo.DataList(ctx)
	if err != nil {
		return nil, 0, err
	}

	return items, tCount, nil
}

func (c *Data) Get(ctx context.Context, id string, errNE bool) (*entities.DataSt, error) {
	result, err := c.r.repo.DataGet(ctx, id)
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

func (c *Data) IdExists(ctx context.Context, id string) (bool, error) {
	return c.r.repo.DataIdExists(ctx, id)
}

func (c *Data) Create(ctx context.Context, obj *entities.DataCUSt) (string, error) {
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	// create
	result, err := c.r.repo.DataCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Data) Update(ctx context.Context, id string, obj *entities.DataCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	err = c.r.repo.DataUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Data) Delete(ctx context.Context, id string) error {
	return c.r.repo.DataDelete(ctx, id)
}

func (c *Data) Deploy(ctx context.Context, obj *entities.DataDeployReqSt) error {
	if obj.ConfFile == "" {
		return errs.ConfFileNameRequired
	}

	err := os.WriteFile(filepath.Join(c.r.confDir, obj.ConfFile), obj.Data, os.ModePerm)
	if err != nil {
		return errs.FailToSaveFile
	}

	if obj.Url != "" {
		hClient := httpclient.New(c.r.lg, &httpc.OptionsSt{
			Client: &http.Client{
				Timeout:   15 * time.Second,
				Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
			},
			Method:    obj.Method,
			LogPrefix: "Deploy webhook:",
		})

		_, err = hClient.Send(&httpc.OptionsSt{
			Uri: obj.Url,
		})
		if err != nil {
			return errs.FailToSendDeployWebhook
		}
	}

	return nil
}
