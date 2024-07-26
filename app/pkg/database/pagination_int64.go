package database

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
	"math"
)

type Paginate64 struct {
	Limit      int64 `json:"limit,omitempty"`
	Page       int64 `json:"page,omitempty"`
	TotalRows  int64 `json:"total_rows,omitempty"`
	TotalPages int64 `json:"total_pages,omitempty"`
}

func (mp *Paginate64) GetMongoPaginatedOpts() *options.FindOptions {
	l := mp.Limit
	skip := mp.Page*mp.Limit - mp.Limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}

	return &fOpt
}

func (mp *Paginate64) InitiateMongoTotal(count int64) {
	if mp.Limit == 0 {
		mp.Limit = count
	}
	if mp.Page == 0 {
		mp.Page = 1
	}
	mp.TotalRows = count
	mp.TotalPages = int64(math.Ceil(float64(count) / float64(mp.Limit)))
}

func NewSQLPaginate(limit, page int64) *Paginate64 {
	return &Paginate64{
		Limit: limit, Page: page,
	}
}

func (p *Paginate64) PaginatedSQLResult(value interface{}, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	if p.Limit == 0 {
		p.Limit = totalRows
	}
	if p.Page == 0 {
		p.Page = 1
	}
	offset := (p.Page - 1) * p.Limit

	p.TotalRows = totalRows
	p.TotalPages = int64(math.Ceil(float64(totalRows) / float64(p.Limit)))
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(int(p.Limit)).Offset(int(offset))
	}
}
