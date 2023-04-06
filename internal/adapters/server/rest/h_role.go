package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTypes"
	"github.com/rendau/dutchman/internal/domain/entities"
)

// @Router		/role [get]
// @Tags		role
// @Param		query	query	entities.RoleListParsSt	false	"query"
// @Produce	json
// @Success	200	{object}	dopTypes.PaginatedListRep{results=[]entities.RoleSt}
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hRoleList(c *gin.Context) {
	pars := &entities.RoleListParsSt{}
	if !dopHttps.BindQuery(c, pars) {
		return
	}

	result, tCount, err := o.ucs.RoleList(o.getRequestContext(c), pars)
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

// @Router		/role [post]
// @Tags		role
// @Param		body	body		entities.RoleCUSt	false	"body"
// @Success	200		{object}	dopTypes.CreateRep{id=string}
// @Failure	400		{object}	dopTypes.ErrRep
func (o *St) hRoleCreate(c *gin.Context) {
	reqObj := &entities.RoleCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.RoleCreate(o.getRequestContext(c), reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, dopTypes.CreateRep{Id: result})
}

// @Router		/role/:id [get]
// @Tags		role
// @Param		id	path	string	true	"id"
// @Produce	json
// @Success	200	{object}	entities.RoleSt
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hRoleGet(c *gin.Context) {
	id := c.Param("id")

	result, err := o.ucs.RoleGet(o.getRequestContext(c), id)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router		/role/:id [put]
// @Tags		role
// @Param		id		path	string				true	"id"
// @Param		body	body	entities.RoleCUSt	false	"body"
// @Produce	json
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hRoleUpdate(c *gin.Context) {
	id := c.Param("id")

	reqObj := &entities.RoleCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	dopHttps.Error(c, o.ucs.RoleUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router		/role/:id [delete]
// @Tags		role
// @Param		id	path	string	true	"id"
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hRoleDelete(c *gin.Context) {
	id := c.Param("id")

	dopHttps.Error(c, o.ucs.RoleDelete(o.getRequestContext(c), id))
}

// @Router		/role/fetch_remote_uri [post]
// @Tags		role
// @Param		body	body		entities.RoleFetchRemoteReqSt	false	"body"
// @Success	200		{array}		entities.RoleFetchRemoteRepItemSt
// @Failure	400		{object}	dopTypes.ErrRep
func (o *St) hRoleFetchRemoteUri(c *gin.Context) {
	reqObj := &entities.RoleFetchRemoteReqSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result := o.ucs.RoleFetchRemoteUri(o.getRequestContext(c), reqObj.Uri, reqObj.Path)

	c.JSON(http.StatusOK, result)
}
