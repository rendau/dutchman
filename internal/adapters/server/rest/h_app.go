package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTypes"
	"github.com/rendau/dutchman/internal/domain/entities"
)

// @Router  /app [get]
// @Tags    app
// @Param   query query entities.AppListParsSt false "query"
// @Produce json
// @Success 200 {object} dopTypes.PaginatedListRep{results=[]entities.AppSt}
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hAppList(c *gin.Context) {
	pars := &entities.AppListParsSt{}
	if !dopHttps.BindQuery(c, pars) {
		return
	}

	result, tCount, err := o.ucs.AppList(o.getRequestContext(c), pars)
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

// @Router  /app [post]
// @Tags    app
// @Param   body body     entities.AppCUSt false "body"
// @Success 200  {object} dopTypes.CreateRep{id=string}
// @Failure 400  {object} dopTypes.ErrRep
func (o *St) hAppCreate(c *gin.Context) {
	reqObj := &entities.AppCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.AppCreate(o.getRequestContext(c), reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, dopTypes.CreateRep{Id: result})
}

// @Router  /app/:id [get]
// @Tags    app
// @Param   id path string true "id"
// @Produce json
// @Success 200 {object} entities.AppSt
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hAppGet(c *gin.Context) {
	id := c.Param("id")

	result, err := o.ucs.AppGet(o.getRequestContext(c), id)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router  /app/:id [put]
// @Tags    app
// @Param   id   path string           true  "id"
// @Param   body body entities.AppCUSt false "body"
// @Produce json
// @Success 200
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hAppUpdate(c *gin.Context) {
	id := c.Param("id")

	reqObj := &entities.AppCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	dopHttps.Error(c, o.ucs.AppUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router  /app/:id [delete]
// @Tags    app
// @Param   id path string true "id"
// @Success 200
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hAppDelete(c *gin.Context) {
	id := c.Param("id")

	dopHttps.Error(c, o.ucs.AppDelete(o.getRequestContext(c), id))
}
