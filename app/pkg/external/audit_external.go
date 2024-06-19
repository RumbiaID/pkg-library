package external

import (
	"context"
	"encoding/json"
	"github.com/RumbiaID/pkg-library/app/pkg/constants"
	"github.com/RumbiaID/pkg-library/app/pkg/exception"
	"github.com/google/uuid"
	"time"
)

type AuditSvcExternal interface {
	CreateAudit(ctx context.Context, model *RequestAuditCreate) *exception.Exception
}

type RequestAuditCreate struct {
	RequestUUID    string      `validate:"uuid" json:"request_uuid"`
	TenantCode     string      `gorm:"size:100" json:"tenant_code"`
	Tablename      string      `gorm:"size:100" json:"table_name"`
	TraceID        string      `gorm:"size:50" json:"trace_id"`
	AuditBeginTime *time.Time  `gorm:"autoCreateTime" json:"audit_begin_time"`
	AuditEndTime   *time.Time  `json:"audit_end_time"`
	AuditHost      string      `gorm:"size:256" json:"audit_host"`
	AuditUser      string      `gorm:"size:100" json:"audit_user"`
	ObjectName     string      `gorm:"size:256" json:"object_name"`
	ActionType     string      `validate:"required,eq=I|eq=U|eq=D|eq=L|eq=V|eq=S" gorm:"size:1" json:"action_type"`
	ApprovalStatus int         `validate:"required,eq=0|eq=1|eq=2|eq=3|eq=4" gorm:"size:2" json:"approval_status"`
	OldValue       interface{} `json:"old_value"`
	NewValue       interface{} `json:"new_value"`
	Status         int         `validate:"required,eq=0|eq=1|eq=2|eq=3" gorm:"size:2" json:"status"`
	ErrorMessage   string      `gorm:"size:1024" json:"error_message"`
}

func NewRequestAuditCreate(
	tenant string, tablename string, auditHost string, auditUser string,
	objectName string, actionType string, approvalStatus int, status int,
) *RequestAuditCreate {
	now := time.Now()
	return &RequestAuditCreate{
		TenantCode:  tenant,
		RequestUUID: uuid.NewString(), AuditBeginTime: &now, Tablename: tablename, AuditHost: auditHost,
		AuditUser: auditUser, ObjectName: objectName, ActionType: actionType, ApprovalStatus: approvalStatus,
		Status: status,
	}
}

func (model *RequestAuditCreate) DeclareAuditEndTime() {
	end := time.Now()
	model.AuditEndTime = &end

}

func (model *RequestAuditCreate) DeclareAuditTraceID(id string) {
	model.TraceID = id

}

func (model *RequestAuditCreate) DeclareAuditNewValue(new interface{}) {
	if new != nil {
		newValue, err := json.Marshal(new)
		if err == nil {
			_ = json.Unmarshal(newValue, &model.NewValue)
		}
	}
}

func (model *RequestAuditCreate) DeclareAuditOldValue(old interface{}) {
	if old != nil {
		oldValue, err := json.Marshal(old)
		if err == nil {
			_ = json.Unmarshal(oldValue, &model.OldValue)
		}
	}
}

func (model *RequestAuditCreate) DeclareAuditError(message string, err error) {
	model.Status = constants.AUDIT_STATUS_FAILED
	//model.ErrorMessage = message + " error:" + err.Error()
	if err != nil {
		model.ErrorMessage = message + " error:" + err.Error()
	} else if err == nil {
		model.ErrorMessage = message
	}

}
