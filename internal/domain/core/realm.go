package core

import (
	"context"

	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
)

type Realm struct {
	r *St
}

func NewRealm(r *St) *Realm {
	return &Realm{r: r}
}

func (c *Realm) ValidateCU(ctx context.Context, obj *entities.RealmCUSt, id string) error {
	// forCreate := id == ""

	return nil
}

func (c *Realm) List(ctx context.Context) ([]*entities.RealmSt, int64, error) {
	items, tCount, err := c.r.repo.RealmList(ctx)
	if err != nil {
		return nil, 0, err
	}

	return items, tCount, nil
}

func (c *Realm) Get(ctx context.Context, id string, errNE bool) (*entities.RealmSt, error) {
	result, err := c.r.repo.RealmGet(ctx, id)
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

func (c *Realm) IdExists(ctx context.Context, id string) (bool, error) {
	return c.r.repo.RealmIdExists(ctx, id)
}

func (c *Realm) Create(ctx context.Context, obj *entities.RealmCUSt) (string, error) {
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	// create
	result, err := c.r.repo.RealmCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Realm) Update(ctx context.Context, id string, obj *entities.RealmCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	err = c.r.repo.RealmUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Realm) Delete(ctx context.Context, id string) error {
	return c.r.repo.RealmDelete(ctx, id)
}

func (c *Realm) Deploy(ctx context.Context, id string, obj *entities.RealmDeployReqSt) error {
	// var err error
	//
	// realm, err := c.Get(ctx, id, true)
	// if err != nil {
	// 	return err
	// }
	//
	// err = os.WriteFile(filepath.Join(c.r.confDir, obj.ConfFile), obj.Config, os.ModePerm)
	// if err != nil {
	// 	return errs.FailToSaveFile
	// }
	//
	// if obj.Url != "" {
	// 	hClient := httpclient.New(c.r.lg, &httpc.OptionsSt{
	// 		Client: &http.Client{
	// 			Timeout:   15 * time.Second,
	// 			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	// 		},
	// 		Method:    obj.Method,
	// 		LogPrefix: "Deploy webhook:",
	// 	})
	//
	// 	_, err = hClient.Send(&httpc.OptionsSt{
	// 		Uri: obj.Url,
	// 	})
	// 	if err != nil {
	// 		return errs.FailToSendDeployWebhook
	// 	}
	// }

	return nil
}
