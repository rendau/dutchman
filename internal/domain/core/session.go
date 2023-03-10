package core

import (
	"context"

	"github.com/rendau/dutchman/internal/domain/entities"
)

const sessionContextKey = "user_session"

type Session struct {
	r *St
}

func NewSession(r *St) *Session {
	return &Session{r: r}
}

func (c *Session) Get(ctx context.Context, token string) *entities.Session {
	result := &entities.Session{}

	if token == "" {
		return result
	}

	// check in cache
	_, ok, err := c.r.cache.Get(token)
	if err == nil && ok {
		result.Authed = true
	}

	return result
}

func (c *Session) SetToContext(ctx context.Context, ses *entities.Session) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithValue(ctx, sessionContextKey, ses)
}

func (c *Session) GetFromContext(ctx context.Context) *entities.Session {
	contextV := ctx.Value(sessionContextKey)
	if contextV == nil {
		return &entities.Session{}
	}

	switch ses := contextV.(type) {
	case *entities.Session:
		return ses
	default:
		c.r.lg.Errorw("wrong type of session in context", nil)
		return &entities.Session{}
	}
}
