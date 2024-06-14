package loggingdata

type InsertReturn struct {
	CreatedBy   string
	CreatedHost string
	TenantCode  string
}

type UpdateReturn struct {
	UpdatedBy   string
	UpdatedHost string
	TenantCode  string
}
