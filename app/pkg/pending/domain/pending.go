package domain

import (
	"encoding/json"
	"os"
	"time"
)

const (
	PendingDataTableName = "pending_data"
)

type Pending struct {
	ID          int             `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	TenantCode  string          `gorm:"size:100" json:"tenant_code"`
	TblName     string          `gorm:"size:100;column:table_name" json:"table_name"`
	ObjectId    string          `json:"object_id"`
	ObjectName  string          `json:"object_name"`
	ActionType  string          `validate:"required,eq=I|eq=U|eq=D" gorm:"size:1" json:"action_type"`
	RowStatus   int             `json:"row_status"`
	ReturnNotes string          `json:"return_notes"`
	NewValue    json.RawMessage `gorm:"type:json" json:"new_value,omitempty"`
	PendingBy   string          `gorm:"size:100" json:"pending_by"`
	PendingTime *time.Time      `gorm:"autoCreateTime" json:"pending_time"`
	PendingHost string          `gorm:"size:256" json:"pending_host"`
}

func (model *Pending) TableName() string {
	return os.Getenv("DB_PREFIX") + PendingDataTableName
}
