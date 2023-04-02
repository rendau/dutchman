package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTypes"
	"github.com/rendau/dutchman/internal/domain/entities"
)

// @Router   /perm [get]
// @Tags     perm
// @Param    query  query  entities.PermListParsSt  false  "query"
// @Produce  json
// @Success  200  {object}  dopTypes.PaginatedListRep{results=[]entities.PermSt}
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hPermList(c *gin.Context) {
	pars := &entities.PermListParsSt{}
	if !dopHttps.BindQuery(c, pars) {
		return
	}

	result, tCount, err := o.ucs.PermList(o.getRequestContext(c), pars)
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

// @Router   /perm [post]
// @Tags     perm
// @Param    body  body  entities.PermCUSt  false  "body"
// @Success  200  {object} dopTypes.CreateRep{id=string}
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hPermCreate(c *gin.Context) {
	reqObj := &entities.PermCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.PermCreate(o.getRequestContext(c), reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, dopTypes.CreateRep{Id: result})
}

// @Router   /perm/:id [get]
// @Tags     perm
// @Param    id path string true "id"
// @Produce  json
// @Success  200  {object}  entities.PermSt
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hPermGet(c *gin.Context) {
	id := c.Param("id")

	result, err := o.ucs.PermGet(o.getRequestContext(c), id)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router   /perm/:id [put]
// @Tags     perm
// @Param    id path string true "id"
// @Param    body  body  entities.PermCUSt  false  "body"
// @Produce  json
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hPermUpdate(c *gin.Context) {
	id := c.Param("id")

	reqObj := &entities.PermCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	dopHttps.Error(c, o.ucs.PermUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router   /perm/:id [delete]
// @Tags     perm
// @Param    id path string true "id"
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (o *St) hPermDelete(c *gin.Context) {
	id := c.Param("id")

	dopHttps.Error(c, o.ucs.PermDelete(o.getRequestContext(c), id))
}
