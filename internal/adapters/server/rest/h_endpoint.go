package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTypes"
	"github.com/rendau/dutchman/internal/domain/entities"
)

// @Router  /endpoint [get]
// @Tags    endpoint
// @Param   query query entities.EndpointListParsSt false "query"
// @Produce json
// @Success 200 {object} dopTypes.PaginatedListRep{results=[]entities.EndpointSt}
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hEndpointList(c *gin.Context) {
	pars := &entities.EndpointListParsSt{}
	if !dopHttps.BindQuery(c, pars) {
		return
	}

	result, tCount, err := o.ucs.EndpointList(o.getRequestContext(c), pars)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, dopTypes.PaginatedListRep{
		Page:       pars.Page,
		PageSize:   pars.PageSize,
		TotalCount: tCount,
		Results:    result,
	})
}

// @Router  /endpoint [post]
// @Tags    endpoint
// @Param   body body     entities.EndpointCUSt false "body"
// @Success 200  {object} dopTypes.CreateRep{id=string}
// @Failure 400  {object} dopTypes.ErrRep
func (o *St) hEndpointCreate(c *gin.Context) {
	reqObj := &entities.EndpointCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.EndpointCreate(o.getRequestContext(c), reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, dopTypes.CreateRep{Id: result})
}

// @Router  /endpoint/:id [get]
// @Tags    endpoint
// @Param   id    path  string                     true  "id"
// @Param   query query entities.EndpointGetParsSt false "query"
// @Produce json
// @Success 200 {object} entities.EndpointSt
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hEndpointGet(c *gin.Context) {
	id := c.Param("id")

	pars := &entities.EndpointGetParsSt{}
	if !dopHttps.BindQuery(c, pars) {
		return
	}

	result, err := o.ucs.EndpointGet(o.getRequestContext(c), id, pars)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router  /endpoint/:id [put]
// @Tags    endpoint
// @Param   id   path string                true  "id"
// @Param   body body entities.EndpointCUSt false "body"
// @Produce json
// @Success 200
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hEndpointUpdate(c *gin.Context) {
	id := c.Param("id")

	reqObj := &entities.EndpointCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	dopHttps.Error(c, o.ucs.EndpointUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router  /endpoint/:id [delete]
// @Tags    endpoint
// @Param   id path string true "id"
// @Success 200
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hEndpointDelete(c *gin.Context) {
	id := c.Param("id")

	dopHttps.Error(c, o.ucs.EndpointDelete(o.getRequestContext(c), id))
}
