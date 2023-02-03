package core

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/dutchman/internal/domain/entities"
	"github.com/rendau/dutchman/internal/domain/errs"
)

const (
	accessTokenDuration = 10 * time.Minute
)

type Profile struct {
	r *St
}

func NewProfile(r *St) *Profile {
	return &Profile{r: r}
}

func (c *Profile) Auth(ctx context.Context, obj *entities.ProfileAuthReqSt) (*entities.ProfileAuthRepSt, error) {
	if obj.Password != c.r.authPassword {
		return nil, errs.WrongPassword
	}

	return &entities.ProfileAuthRepSt{
		RefreshToken: c.r.sessionRefreshToken,
		AccessToken:  c.generateAccessToken(),
	}, nil
}

func (c *Profile) AuthByRefreshToken(ctx context.Context, obj *entities.ProfileAuthByRefreshTokenReqSt) (*entities.ProfileAuthByRefreshTokenRepSt, error) {
	if obj.RefreshToken != c.r.sessionRefreshToken {
		return nil, dopErrs.NotAuthorized
	}

	return &entities.ProfileAuthByRefreshTokenRepSt{
		AccessToken: c.generateAccessToken(),
	}, nil
}

func (c *Profile) Get(ctx context.Context) (*entities.ProfileSt, error) {
	return &entities.ProfileSt{}, nil
}

func (c *Profile) generateAccessToken() string {
	token := fmt.Sprintf("%x", sha256.Sum256([]byte(time.Now().String())))

	c.r.cache.Set(token, []byte("1"), accessTokenDuration)

	return token
}
