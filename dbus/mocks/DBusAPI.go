// Copyright 2020 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	dbus "github.com/mendersoftware/mender/dbus"
	mock "github.com/stretchr/testify/mock"
)

// DBusAPI is an autogenerated mock type for the DBusAPI type
type DBusAPI struct {
	mock.Mock
}

// BusGet provides a mock function with given fields: _a0
func (_m *DBusAPI) BusGet(_a0 uint) (dbus.Pointer, error) {
	ret := _m.Called(_a0)

	var r0 dbus.Pointer
	if rf, ok := ret.Get(0).(func(uint) dbus.Pointer); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(dbus.Pointer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BusOwnNameOnConnection provides a mock function with given fields: _a0, _a1, _a2
func (_m *DBusAPI) BusOwnNameOnConnection(_a0 dbus.Pointer, _a1 string, _a2 uint) (uint, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint
	if rf, ok := ret.Get(0).(func(dbus.Pointer, string, uint) uint); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dbus.Pointer, string, uint) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BusRegisterInterface provides a mock function with given fields: _a0, _a1, _a2
func (_m *DBusAPI) BusRegisterInterface(_a0 dbus.Pointer, _a1 string, _a2 string) (uint, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint
	if rf, ok := ret.Get(0).(func(dbus.Pointer, string, string) uint); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dbus.Pointer, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateGUID provides a mock function with given fields:
func (_m *DBusAPI) GenerateGUID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsGUID provides a mock function with given fields: _a0
func (_m *DBusAPI) IsGUID(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MainLoopNew provides a mock function with given fields:
func (_m *DBusAPI) MainLoopNew() dbus.Pointer {
	ret := _m.Called()

	var r0 dbus.Pointer
	if rf, ok := ret.Get(0).(func() dbus.Pointer); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(dbus.Pointer)
	}

	return r0
}

// MainLoopQuit provides a mock function with given fields: _a0
func (_m *DBusAPI) MainLoopQuit(_a0 dbus.Pointer) {
	_m.Called(_a0)
}

// MainLoopRun provides a mock function with given fields: _a0
func (_m *DBusAPI) MainLoopRun(_a0 dbus.Pointer) {
	_m.Called(_a0)
}

// RegisterMethodCallCallback provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *DBusAPI) RegisterMethodCallCallback(_a0 string, _a1 string, _a2 string, _a3 func(string, string, string) (interface{}, error)) {
	_m.Called(_a0, _a1, _a2, _a3)
}
