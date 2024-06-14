package loggingdata

import (
	"github.com/gin-gonic/gin"
)

func InsertData(ctx *gin.Context) *InsertReturn {
	header := ctx.Request.Header
	username := header.Get("X-Sso-Username")
	hostname := ctx.ClientIP()
	tenantCode := header.Get("X-Sso-Tenantcode")

	return &InsertReturn{
		CreatedBy:   username,
		CreatedHost: hostname,
		TenantCode:  tenantCode,
	}
}
