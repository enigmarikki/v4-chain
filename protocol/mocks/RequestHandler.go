// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RequestHandler is an autogenerated mock type for the RequestHandler type
type RequestHandler struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, url
func (_m *RequestHandler) Get(ctx context.Context, url string) (*http.Response, error) {
	ret := _m.Called(ctx, url)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(context.Context, string) *http.Response); ok {
		r0 = rf(ctx, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRequestHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequestHandler creates a new instance of RequestHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequestHandler(t mockConstructorTestingTNewRequestHandler) *RequestHandler {
	mock := &RequestHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}