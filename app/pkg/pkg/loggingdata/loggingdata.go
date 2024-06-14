package loggingdata

import "cpm-portfolio-service/app/internal/handlers/app"

func InsertData(ctx *app.Context) *InsertReturn {
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

func UpdateData(ctx *app.Context) *UpdateReturn {
	header := ctx.Request.Header
	username := header.Get("X-Sso-Username")
	hostname := header.Get("X-Real-Ip")
	tenantCode := header.Get("X-Sso-Tenantcode")

	return &UpdateReturn{
		UpdatedBy:   username,
		UpdatedHost: hostname,
		TenantCode:  tenantCode,
	}
}
