// Code generated by mockery (devel). DO NOT EDIT.

package errs

import mock "github.com/stretchr/testify/mock"

// HttpError is an autogenerated mock type for the HttpError type
type HttpError struct {
	mock.Mock
}

// Error provides a mock function with given fields:
func (_m *HttpError) Error() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// StatusCode provides a mock function with given fields:
func (_m *HttpError) StatusCode() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
