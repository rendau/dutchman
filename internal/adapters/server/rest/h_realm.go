package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTypes"
	"github.com/rendau/dutchman/internal/domain/entities"
)

// @Router		/realm [get]
// @Tags		realm
// @Param		query	query	entities.RealmListParsSt	false	"query"
// @Produce	json
// @Success	200	{object}	dopTypes.PaginatedListRep{results=[]entities.RealmSt}
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hRealmList(c *gin.Context) {
	pars := &entities.RealmListParsSt{}
	if !dopHttps.BindQuery(c, pars) {
		return
	}

	result, tCount, err := o.ucs.RealmList(o.getRequestContext(c), pars)
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

// @Router		/realm [post]
// @Tags		realm
// @Param		body	body		entities.RealmCUSt	false	"body"
// @Success	200		{object}	dopTypes.CreateRep{id=string}
// @Failure	400		{object}	dopTypes.ErrRep
func (o *St) hRealmCreate(c *gin.Context) {
	reqObj := &entities.RealmCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.RealmCreate(o.getRequestContext(c), reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, dopTypes.CreateRep{Id: result})
}

// @Router		/realm/:id [get]
// @Tags		realm
// @Param		id	path	string	true	"id"
// @Produce	json
// @Success	200	{object}	entities.RealmSt
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hRealmGet(c *gin.Context) {
	id := c.Param("id")

	result, err := o.ucs.RealmGet(o.getRequestContext(c), id)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router		/realm/:id [put]
// @Tags		realm
// @Param		id		path	string				true	"id"
// @Param		body	body	entities.RealmCUSt	false	"body"
// @Produce	json
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hRealmUpdate(c *gin.Context) {
	id := c.Param("id")

	reqObj := &entities.RealmCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	dopHttps.Error(c, o.ucs.RealmUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router		/realm/:id [delete]
// @Tags		realm
// @Param		id	path	string	true	"id"
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hRealmDelete(c *gin.Context) {
	id := c.Param("id")

	dopHttps.Error(c, o.ucs.RealmDelete(o.getRequestContext(c), id))
}

// @Router		/realm/:id/preview_conf [GET]
// @Tags		realm
// @Param		id	path		string	true	"id"
// @Success	200	{object}	entities.KrakendSt
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hRealmPreviewConf(c *gin.Context) {
	id := c.Param("id")

	result, err := o.ucs.RealmPreviewConf(o.getRequestContext(c), id)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router		/realm/:id/deploy [POST]
// @Tags		realm
// @Param		id	path	string	true	"id"
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hRealmDeploy(c *gin.Context) {
	id := c.Param("id")

	dopHttps.Error(c, o.ucs.RealmDeploy(o.getRequestContext(c), id))
}

// @Router		/realm/:id/import_conf [POST]
// @Tags		realm
// @Param		id		path	string				true	"id"
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hRealmImportConf(c *gin.Context) {
	id := c.Param("id")

	reqObj := &entities.KrakendSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	dopHttps.Error(c, o.ucs.RealmImportConf(o.getRequestContext(c), id, reqObj))
}
