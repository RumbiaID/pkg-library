// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/RumbiaID/pkg-library/app/pkg/pending/domain"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// PendingRepository is an autogenerated mock type for the PendingRepository type
type PendingRepository struct {
	mock.Mock
}

// CreatePending provides a mock function with given fields: ctx, tx, model
func (_m *PendingRepository) CreatePending(ctx context.Context, tx *gorm.DB, model *domain.Pending) error {
	ret := _m.Called(ctx, tx, model)

	if len(ret) == 0 {
		panic("no return value specified for CreatePending")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *domain.Pending) error); ok {
		r0 = rf(ctx, tx, model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePending provides a mock function with given fields: ctx, tx, id
func (_m *PendingRepository) DeletePending(ctx context.Context, tx *gorm.DB, id int) error {
	ret := _m.Called(ctx, tx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeletePending")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, int) error); ok {
		r0 = rf(ctx, tx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPending provides a mock function with given fields: ctx, tx, tenantcode, tablename
func (_m *PendingRepository) GetPending(ctx context.Context, tx *gorm.DB, tenantcode string, tablename string) (*[]domain.Pending, error) {
	ret := _m.Called(ctx, tx, tenantcode, tablename)

	if len(ret) == 0 {
		panic("no return value specified for GetPending")
	}

	var r0 *[]domain.Pending
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, string, string) (*[]domain.Pending, error)); ok {
		return rf(ctx, tx, tenantcode, tablename)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, string, string) *[]domain.Pending); ok {
		r0 = rf(ctx, tx, tenantcode, tablename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Pending)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, string, string) error); ok {
		r1 = rf(ctx, tx, tenantcode, tablename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePending provides a mock function with given fields: ctx, tx, model
func (_m *PendingRepository) UpdatePending(ctx context.Context, tx *gorm.DB, model *domain.Pending) error {
	ret := _m.Called(ctx, tx, model)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePending")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *domain.Pending) error); ok {
		r0 = rf(ctx, tx, model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPendingRepository creates a new instance of PendingRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPendingRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PendingRepository {
	mock := &PendingRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
