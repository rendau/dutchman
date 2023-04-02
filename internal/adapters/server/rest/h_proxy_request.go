package rest

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rendau/dop/adapters/client/httpc"
	"github.com/rendau/dop/adapters/client/httpc/httpclient"
	dopHttps "github.com/rendau/dop/adapters/server/https"
)

// @Router		/proxy_request [POST]
// @Tags		proxy_request
// @Param		body	body	entities.ProxyRequestSendReqSt	false	"body"
// @Success	200
// @Failure	400	{object}	dopTypes.ErrRep
func (o *St) hProxyRequest(c *gin.Context) {
	const timeOut = 5 * time.Second

	reqObj := map[string]string{}
	if !dopHttps.BindJSON(c, &reqObj) {
		return
	}

	targetUrl := reqObj["url"]

	hClient := httpclient.New(o.lg, &httpc.OptionsSt{
		Client: &http.Client{
			Timeout:   timeOut,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		},
		LogFlags: httpc.NoLogError,
	})

	resp, err := hClient.Send(&httpc.OptionsSt{
		Method:    "GET",
		Uri:       targetUrl,
		RepStream: true,
	})
	if dopHttps.Error(c, err) {
		return
	}
	defer resp.Stream.Close()

	_, _ = io.Copy(c.Writer, resp.Stream)
}
