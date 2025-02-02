// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"
	bindings "razor/pkg/bindings"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"
)

// CoinUtils is an autogenerated mock type for the CoinUtils type
type CoinUtils struct {
	mock.Mock
}

// BalanceOf provides a mock function with given fields: _a0, _a1, _a2
func (_m *CoinUtils) BalanceOf(_a0 *bindings.RAZOR, _a1 *bind.CallOpts, _a2 common.Address) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bindings.RAZOR, *bind.CallOpts, common.Address) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bindings.RAZOR, *bind.CallOpts, common.Address) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
