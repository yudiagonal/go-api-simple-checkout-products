// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/yudiagonal/go-api-simple-checkout-products/internal/app/product/dto"
	mock "github.com/stretchr/testify/mock"

	model "github.com/yudiagonal/go-api-simple-checkout-products/internal/model"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, payload
func (_m *ProductRepository) Create(ctx context.Context, payload dto.InsertProductDto) (*model.Product, error) {
	ret := _m.Called(ctx, payload)

	var r0 *model.Product
	if rf, ok := ret.Get(0).(func(context.Context, dto.InsertProductDto) *model.Product); ok {
		r0 = rf(ctx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.InsertProductDto) error); ok {
		r1 = rf(ctx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProduct provides a mock function with given fields: ctx, filter
func (_m *ProductRepository) GetProduct(ctx context.Context, filter dto.FilterProductDto) ([]dto.GetProduct, error) {
	ret := _m.Called(ctx, filter)

	var r0 []dto.GetProduct
	if rf, ok := ret.Get(0).(func(context.Context, dto.FilterProductDto) []dto.GetProduct); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.GetProduct)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.FilterProductDto) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductRepository(t mockConstructorTestingTNewProductRepository) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
