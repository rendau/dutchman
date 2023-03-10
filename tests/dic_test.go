package tests

import (
	"context"
	"testing"

	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
	"github.com/stretchr/testify/require"
)

func TestCfg(t *testing.T) {
	prepareDbForNewTest()

	bgCtx := context.Background()
	admCtx := app.ucs.SessionSetToContext(context.Background(), &entities.Session{
		Id: 1,
	})

	cfg, err := app.ucs.ConfigGet(bgCtx)
	require.Nil(t, err)
	require.NotNil(t, cfg)
	require.Empty(t, cfg.Contacts.Phone)
	require.Empty(t, cfg.Contacts.Email)

	err = app.ucs.ConfigSet(bgCtx, &entities.ConfigSt{})
	require.Equal(t, dopErrs.NotAuthorized, err)

	err = app.ucs.ConfigSet(admCtx, &entities.ConfigSt{
		Contacts: entities.ConfigContactsSt{
			Phone: "71230000001",
			Email: "qwe@asd.com",
		},
	})
	require.Nil(t, err)

	cfg, err = app.ucs.ConfigGet(bgCtx)
	require.Nil(t, err)
	require.NotNil(t, cfg)
	require.Equal(t, "71230000001", cfg.Contacts.Phone)
	require.Equal(t, "qwe@asd.com", cfg.Contacts.Email)
}
