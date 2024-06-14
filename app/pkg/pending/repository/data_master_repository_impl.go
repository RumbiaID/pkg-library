package repository

import (
	"context"
	"gorm.io/gorm"
	"pkg-library/app/pkg/getfilter"
	"pkg-library/app/pkg/pending/domain"
)

type DataMasterRepositoryImpl struct {
	db *gorm.DB
}

func NewDataMasterRepositoryImpl(db *gorm.DB) DataMasterRepository {
	return &DataMasterRepositoryImpl{db: db}
}

func (r *DataMasterRepositoryImpl) CreatePending(ctx context.Context, tx *gorm.DB, model *domain.Pending) error {

	query := tx.WithContext(ctx).Model(&domain.Pending{})
	if err := query.Create(&model).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *DataMasterRepositoryImpl) UpdatePending(ctx context.Context, tx *gorm.DB, model *domain.Pending) error {
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

func (r *DataMasterRepositoryImpl) DeletePending(ctx context.Context, tx *gorm.DB, id int) error {
	query := tx.WithContext(ctx)

	if err := query.
		Delete(&domain.Pending{ID: id}).Error; err != nil {
		return err
	}
	return nil
}

func (r *DataMasterRepositoryImpl) GetPending(
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

func (r *DataMasterRepositoryImpl) queryGoals(
	tx *gorm.DB, arrQuery []getfilter.FilterItem, arrSort []getfilter.FilterSort,
) *gorm.DB {
	for _, filter := range arrQuery {
		keyList := getfilter.GenerateWhere(filter)
		switch filter.Field {
		case "id", "goal_type", "is_approve", "hex_color", "created_at", "created_by", "created_nik", "updated_at", "updated_nik", "pending_id", "sys_row_status", "tenant_code":
			if filter.Operator == "is" {
				switch filter.Value {
				case "null":
					tx = tx.Where(filter.Field + " IS NULL")
				case "not null":
					tx = tx.Where(filter.Field + " IS NOT NULL")
				}
			} else {
				tx = tx.Where(filter.Field+" "+filter.Operator+" ?", keyList)
			}
		}
	}
	for _, filter := range arrSort {
		switch filter.Field {
		case "id", "goal_type", "is_approve", "hex_color", "created_at", "created_by", "created_nik", "updated_at", "updated_nik", "pending_id", "sys_row_status", "tenant_code":
			switch filter.Value {
			case "asc", "desc":
				tx = tx.Order(filter.Field + " " + filter.Value)
			}
		}
	}
	//tx.Where("deleted_at", nil)
	return tx
}
