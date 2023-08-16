// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	perpetualstypes "github.com/dydxprotocol/v4/x/perpetuals/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// PreparePerpetualsKeeper is an autogenerated mock type for the PreparePerpetualsKeeper type
type PreparePerpetualsKeeper struct {
	mock.Mock
}

// GetAddPremiumVotes provides a mock function with given fields: ctx
func (_m *PreparePerpetualsKeeper) GetAddPremiumVotes(ctx types.Context) *perpetualstypes.MsgAddPremiumVotes {
	ret := _m.Called(ctx)

	var r0 *perpetualstypes.MsgAddPremiumVotes
	if rf, ok := ret.Get(0).(func(types.Context) *perpetualstypes.MsgAddPremiumVotes); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*perpetualstypes.MsgAddPremiumVotes)
		}
	}

	return r0
}

type mockConstructorTestingTNewPreparePerpetualsKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewPreparePerpetualsKeeper creates a new instance of PreparePerpetualsKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPreparePerpetualsKeeper(t mockConstructorTestingTNewPreparePerpetualsKeeper) *PreparePerpetualsKeeper {
	mock := &PreparePerpetualsKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}