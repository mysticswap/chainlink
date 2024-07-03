// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	keystore "github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	mock "github.com/stretchr/testify/mock"
)

// Master is an autogenerated mock type for the Master type
type Master struct {
	mock.Mock
}

// Aptos provides a mock function with given fields:
func (_m *Master) Aptos() keystore.Aptos {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Aptos")
	}

	var r0 keystore.Aptos
	if rf, ok := ret.Get(0).(func() keystore.Aptos); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.Aptos)
		}
	}

	return r0
}

// CSA provides a mock function with given fields:
func (_m *Master) CSA() keystore.CSA {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CSA")
	}

	var r0 keystore.CSA
	if rf, ok := ret.Get(0).(func() keystore.CSA); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.CSA)
		}
	}

	return r0
}

// Cosmos provides a mock function with given fields:
func (_m *Master) Cosmos() keystore.Cosmos {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Cosmos")
	}

	var r0 keystore.Cosmos
	if rf, ok := ret.Get(0).(func() keystore.Cosmos); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.Cosmos)
		}
	}

	return r0
}

// Eth provides a mock function with given fields:
func (_m *Master) Eth() keystore.Eth {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Eth")
	}

	var r0 keystore.Eth
	if rf, ok := ret.Get(0).(func() keystore.Eth); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.Eth)
		}
	}

	return r0
}

// IsEmpty provides a mock function with given fields: ctx
func (_m *Master) IsEmpty(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for IsEmpty")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (bool, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OCR provides a mock function with given fields:
func (_m *Master) OCR() keystore.OCR {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OCR")
	}

	var r0 keystore.OCR
	if rf, ok := ret.Get(0).(func() keystore.OCR); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.OCR)
		}
	}

	return r0
}

// OCR2 provides a mock function with given fields:
func (_m *Master) OCR2() keystore.OCR2 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OCR2")
	}

	var r0 keystore.OCR2
	if rf, ok := ret.Get(0).(func() keystore.OCR2); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.OCR2)
		}
	}

	return r0
}

// P2P provides a mock function with given fields:
func (_m *Master) P2P() keystore.P2P {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for P2P")
	}

	var r0 keystore.P2P
	if rf, ok := ret.Get(0).(func() keystore.P2P); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.P2P)
		}
	}

	return r0
}

// Solana provides a mock function with given fields:
func (_m *Master) Solana() keystore.Solana {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Solana")
	}

	var r0 keystore.Solana
	if rf, ok := ret.Get(0).(func() keystore.Solana); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.Solana)
		}
	}

	return r0
}

// StarkNet provides a mock function with given fields:
func (_m *Master) StarkNet() keystore.StarkNet {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StarkNet")
	}

	var r0 keystore.StarkNet
	if rf, ok := ret.Get(0).(func() keystore.StarkNet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.StarkNet)
		}
	}

	return r0
}

// Unlock provides a mock function with given fields: ctx, password
func (_m *Master) Unlock(ctx context.Context, password string) error {
	ret := _m.Called(ctx, password)

	if len(ret) == 0 {
		panic("no return value specified for Unlock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VRF provides a mock function with given fields:
func (_m *Master) VRF() keystore.VRF {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for VRF")
	}

	var r0 keystore.VRF
	if rf, ok := ret.Get(0).(func() keystore.VRF); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(keystore.VRF)
		}
	}

	return r0
}

// NewMaster creates a new instance of Master. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMaster(t interface {
	mock.TestingT
	Cleanup(func())
}) *Master {
	mock := &Master{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
