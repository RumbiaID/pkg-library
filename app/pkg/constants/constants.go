package constants

const (
	ACTION_TYPE_INSERT           = "I"
	ACTION_TYPE_UPDATE           = "U"
	ACTION_TYPE_DELETE           = "D"
	ACTION_TYPE_LIST             = "L"
	ACTION_TYPE_VIEW             = "V"
	ACTION_TYPE_SEARCH           = "S"
	SYSROW_STATUS_PENDING_INSERT = 0
	SYSROW_STATUS_ACTIVE         = 1
	SYSROW_STATUS_PENDING_UPDATE = 2
	SYSROW_STATUS_PENDING_DELETE = 3
	SYSROW_STATUS_RETURN_INSERT  = 4
	SYSROW_STATUS_RETURN_UPDATE  = 5
	APPROVAL_STATUS_PENDING      = 0
	APPROVAL_STATUS_APPROVE      = 1
	APPROVAL_STATUS_REJECT       = 2
	APPROVAL_STATUS_RETURN       = 3
	APPROVAL_STATUS_RETRY        = 4
	APPROVAL_STATUS_DIRECT       = 5
	AUDIT_STATUS_STARTED         = 0
	AUDIT_STATUS_SUCCESS         = 1
	AUDIT_STATUS_FAILED          = 2
	AUDIT_STATUS_UNKNOWN         = 3
	DATA_STATUS_DELETED          = "Deleted"
	DATA_STATUS_ACTIVE           = "Active"
	DATA_STATUS_PENDING_INSERT   = "Pending Insert"
	DATA_STATUS_PENDING_UPDATE   = "Pending Update"
	DATA_STATUS_PENDING_DELETE   = "Pending Delete"
	DATA_STATUS_PENDING_RETURN   = "Pending Return"
)

var (
	FILTER_PENDING         = []int{SYSROW_STATUS_PENDING_INSERT, SYSROW_STATUS_PENDING_DELETE}
	TYPE_INCOME_EXPENSE    = []string{"Income", "Expense"}
	TYPE_ASSET_LIABILITIES = []string{"Asset", "Liabilities"}
	PENDING_INSERT_FIELD   = []string{
		"sys_row_status", "sys_created_by", "sys_created_host", "sys_last_pending_by",
		"sys_last_pending_host", "sys_last_pending_time",
	}
	PENDING_UPDATE_FIELD = []string{
		"sys_row_status", "sys_last_pending_by", "sys_last_approve_by",
		"sys_last_approve_host", "sys_last_approval_notes", "sys_last_pending_host",
		"sys_last_pending_time", "sys_last_approve_time", "pending_id",
	}
	PENDING_DELETE_FIELD = []string{
		"sys_row_status", "sys_last_pending_by", "sys_last_approve_by",
		"sys_last_approve_host", "sys_last_approval_notes", "sys_last_pending_host",
		"sys_last_pending_time", "sys_last_approve_time", "pending_id",
	}
	APPROVE_UPSERT_FIELD = []string{
		"sys_row_status", "sys_last_approve_by", "sys_last_approve_host",
		"sys_last_approval_notes", "sys_last_pending_time", "sys_last_approve_time",
		"pending_id",
	}
	RETURN_UPSERT_FIELD = []string{
		"sys_row_status", "sys_last_approval_notes",
	}
	REJECT_DELUP_FIELD = []string{
		"sys_row_status", "sys_last_approve_by", "sys_last_approve_time",
		"sys_last_approve_host", "sys_last_approval_notes", "pending_id",
	}
	RETRY_INSERT_FIELD = []string{
		"sys_row_status", "sys_created_by", "sys_created_host",
		"sys_last_pending_by", "sys_last_pending_time", "sys_last_pending_host",
	}
)
