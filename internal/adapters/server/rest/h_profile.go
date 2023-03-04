package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dutchman/internal/domain/entities"
)

// @Router  /profile [get]
// @Tags    profile
// @Produce json
// @Success 200 {object} entities.ProfileSt
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hProfileGet(c *gin.Context) {
	repObj, err := o.ucs.ProfileGet(o.getRequestContext(c))
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, repObj)
}

// @Router  /profile/auth [post]
// @Tags    profile
// @Param   body body entities.ProfileAuthReqSt false "body"
// @Produce json
// @Success 200 {object} entities.ProfileAuthRepSt
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hProfileAuth(c *gin.Context) {
	reqObj := &entities.ProfileAuthReqSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.ProfileAuth(context.Background(), reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router  /profile/auth/token [post]
// @Tags    profile
// @Param   body body entities.ProfileAuthByRefreshTokenReqSt false "body"
// @Produce json
// @Success 200 {object} entities.ProfileAuthByRefreshTokenRepSt
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hProfileAuthByRefreshToken(c *gin.Context) {
	reqObj := &entities.ProfileAuthByRefreshTokenReqSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.ProfileAuthByRefreshToken(context.Background(), reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}
