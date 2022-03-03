// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	fs "io/fs"
	os "os"

	mock "github.com/stretchr/testify/mock"
)

// OSInterface is an autogenerated mock type for the OSInterface type
type OSInterface struct {
	mock.Mock
}

// IsNotExist provides a mock function with given fields: _a0
func (_m *OSInterface) IsNotExist(_a0 error) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(error) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Mkdir provides a mock function with given fields: _a0, _a1
func (_m *OSInterface) Mkdir(_a0 string, _a1 fs.FileMode) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, fs.FileMode) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OpenFile provides a mock function with given fields: _a0, _a1, _a2
func (_m *OSInterface) OpenFile(_a0 string, _a1 int, _a2 fs.FileMode) (*os.File, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *os.File
	if rf, ok := ret.Get(0).(func(string, int, fs.FileMode) *os.File); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, fs.FileMode) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stat provides a mock function with given fields: _a0
func (_m *OSInterface) Stat(_a0 string) (fs.FileInfo, error) {
	ret := _m.Called(_a0)

	var r0 fs.FileInfo
	if rf, ok := ret.Get(0).(func(string) fs.FileInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fs.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserHomeDir provides a mock function with given fields:
func (_m *OSInterface) UserHomeDir() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
