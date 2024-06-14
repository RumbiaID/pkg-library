package repository

import (
	"context"
	"github.com/RumbiaID/pkg-library/app/pkg/pending/domain"
	"gorm.io/gorm"
)

type PendingRepositoryImpl struct {
	db *gorm.DB
}

func NewPendingRepository(db *gorm.DB) PendingRepository {
	return &PendingRepositoryImpl{db: db}
}

func (r *PendingRepositoryImpl) CreatePending(ctx context.Context, tx *gorm.DB, model *domain.Pending) error {

	query := tx.WithContext(ctx).Model(&domain.Pending{})
	if err := query.Create(&model).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *PendingRepositoryImpl) UpdatePending(ctx context.Context, tx *gorm.DB, model *domain.Pending) error {
	query := tx.WithContext(ctx)
	if err := query.
		Model(&domain.Pending{ID: model.ID}).
		Select("new_value").
		Updates(model).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *PendingRepositoryImpl) DeletePending(ctx context.Context, tx *gorm.DB, id int) error {
	query := tx.WithContext(ctx)

	if err := query.
		Delete(&domain.Pending{ID: id}).Error; err != nil {
		return err
	}
	return nil
}

func (r *PendingRepositoryImpl) GetPending(
	ctx context.Context, tx *gorm.DB, tenantcode, tablename string,
) (*[]domain.Pending, error) {
	var models *[]domain.Pending
	if err := tx.
		WithContext(ctx).
		Model(&domain.Pending{}).Where("tenant_code = ?", tenantcode).Where("table_name = ?", tablename).
		Find(&models).
		Error; err != nil {
		return nil, err
	}
	return models, nil
}
