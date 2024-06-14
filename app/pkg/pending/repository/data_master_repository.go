package repository

import (
	"context"
	"gorm.io/gorm"
	"pkg-library/app/pkg/pending/domain"
)

type PendingRepository interface {
	//Pending
	CreatePending(ctx context.Context, tx *gorm.DB, model *domain.Pending) error
	UpdatePending(ctx context.Context, tx *gorm.DB, model *domain.Pending) error
	DeletePending(ctx context.Context, tx *gorm.DB, id int) error
	GetPending(ctx context.Context, tx *gorm.DB, tenantcode, tablename string) (*[]domain.Pending, error)
}
