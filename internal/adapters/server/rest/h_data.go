package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTypes"
	"github.com/rendau/dutchman/internal/domain/entities"
)

// @Router  /data [get]
// @Tags    data
// @Produce json
// @Success 200 {object} dopTypes.ListRep{results=[]entities.DataListSt}
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hDataList(c *gin.Context) {
	result, _, err := o.ucs.DataList(o.getRequestContext(c))
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, dopTypes.ListRep{
		Results: result,
	})
}

// @Router  /data [post]
// @Tags    data
// @Param   body body     entities.DataCUSt false "body"
// @Success 200  {object} dopTypes.CreateRep{id=string}
// @Failure 400  {object} dopTypes.ErrRep
func (o *St) hDataCreate(c *gin.Context) {
	reqObj := &entities.DataCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	result, err := o.ucs.DataCreate(o.getRequestContext(c), reqObj)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, dopTypes.CreateRep{Id: result})
}

// @Router  /data/:id [get]
// @Tags    data
// @Param   id path string true "id"
// @Produce json
// @Success 200 {object} entities.DataSt
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hDataGet(c *gin.Context) {
	id := c.Param("id")

	result, err := o.ucs.DataGet(o.getRequestContext(c), id)
	if dopHttps.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Router  /data/:id [put]
// @Tags    data
// @Param   id   path string            true  "id"
// @Param   body body entities.DataCUSt false "body"
// @Produce json
// @Success 200
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hDataUpdate(c *gin.Context) {
	id := c.Param("id")

	reqObj := &entities.DataCUSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	dopHttps.Error(c, o.ucs.DataUpdate(o.getRequestContext(c), id, reqObj))
}

// @Router  /data/:id [delete]
// @Tags    data
// @Param   id path string true "id"
// @Success 200
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hDataDelete(c *gin.Context) {
	id := c.Param("id")

	dopHttps.Error(c, o.ucs.DataDelete(o.getRequestContext(c), id))
}

// @Router  /data/deploy [POST]
// @Tags    data
// @Param   body body entities.DataDeployReqSt false "body"
// @Success 200
// @Failure 400 {object} dopTypes.ErrRep
func (o *St) hDataDeploy(c *gin.Context) {
	reqObj := &entities.DataDeployReqSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	dopHttps.Error(c, o.ucs.DataDeploy(o.getRequestContext(c), reqObj))
}
