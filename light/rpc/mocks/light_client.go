// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	time "time"

	types "github.com/line/ostracon/types"
)

// LightClient is an autogenerated mock type for the LightClient type
type LightClient struct {
	mock.Mock
}

// ChainID provides a mock function with given fields:
func (_m *LightClient) ChainID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TrustedLightBlock provides a mock function with given fields: height
func (_m *LightClient) TrustedLightBlock(height int64) (*types.LightBlock, error) {
	ret := _m.Called(height)

	var r0 *types.LightBlock
	if rf, ok := ret.Get(0).(func(int64) *types.LightBlock); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.LightBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, now
func (_m *LightClient) Update(ctx context.Context, now time.Time) (*types.LightBlock, error) {
	ret := _m.Called(ctx, now)

	var r0 *types.LightBlock
	if rf, ok := ret.Get(0).(func(context.Context, time.Time) *types.LightBlock); ok {
		r0 = rf(ctx, now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.LightBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time) error); ok {
		r1 = rf(ctx, now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyLightBlockAtHeight provides a mock function with given fields: ctx, height, now
func (_m *LightClient) VerifyLightBlockAtHeight(ctx context.Context, height int64, now time.Time) (*types.LightBlock, error) {
	ret := _m.Called(ctx, height, now)

	var r0 *types.LightBlock
	if rf, ok := ret.Get(0).(func(context.Context, int64, time.Time) *types.LightBlock); ok {
		r0 = rf(ctx, height, now)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.LightBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, time.Time) error); ok {
		r1 = rf(ctx, height, now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}