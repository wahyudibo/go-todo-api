// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// StorageAdapter is an autogenerated mock type for the StorageAdapter type
type StorageAdapter struct {
	mock.Mock
}

// GenerateDownloadURL provides a mock function with given fields: ctx, objectKey
func (_m *StorageAdapter) GenerateDownloadURL(ctx context.Context, objectKey string) (string, error) {
	ret := _m.Called(ctx, objectKey)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, objectKey)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, objectKey)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, objectKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upload provides a mock function with given fields: ctx, directory, objectKey, body
func (_m *StorageAdapter) Upload(ctx context.Context, directory string, objectKey string, body io.Reader) (string, error) {
	ret := _m.Called(ctx, directory, objectKey, body)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader) (string, error)); ok {
		return rf(ctx, directory, objectKey, body)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader) string); ok {
		r0 = rf(ctx, directory, objectKey, body)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, io.Reader) error); ok {
		r1 = rf(ctx, directory, objectKey, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStorageAdapter interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorageAdapter creates a new instance of StorageAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorageAdapter(t mockConstructorTestingTNewStorageAdapter) *StorageAdapter {
	mock := &StorageAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
