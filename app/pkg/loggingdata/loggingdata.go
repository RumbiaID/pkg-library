package loggingdata

import (
	"github.com/gin-gonic/gin"
)

func GetRequestHeader(ctx *gin.Context) *RequestHeader {
	header := ctx.Request.Header
	username := header.Get("X-Sso-Username")
	hostname := ctx.ClientIP()
	tenantCode := header.Get("X-Sso-Tenantcode")

	return &RequestHeader{
		CreatedBy:   username,
		CreatedHost: hostname,
		TenantCode:  tenantCode,
	}
}
