// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	http "net/http"

	jwt "github.com/0xTatsu/g-api/jwt"
	mock "github.com/stretchr/testify/mock"
)

// JWT is an autogenerated mock type for the JWT type
type JWT struct {
	mock.Mock
}

// CreateAccessToken provides a mock function with given fields: c
func (_m *JWT) CreateAccessToken(c jwt.AccessClaims) (string, error) {
	ret := _m.Called(c)

	var r0 string
	if rf, ok := ret.Get(0).(func(jwt.AccessClaims) string); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(jwt.AccessClaims) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRefreshToken provides a mock function with given fields: c
func (_m *JWT) CreateRefreshToken(c jwt.RefreshClaims) (string, error) {
	ret := _m.Called(c)

	var r0 string
	if rf, ok := ret.Get(0).(func(jwt.RefreshClaims) string); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(jwt.RefreshClaims) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTokenPair provides a mock function with given fields: accessClaims, refreshClaims
func (_m *JWT) CreateTokenPair(accessClaims jwt.AccessClaims, refreshClaims jwt.RefreshClaims) (string, string, error) {
	ret := _m.Called(accessClaims, refreshClaims)

	var r0 string
	if rf, ok := ret.Get(0).(func(jwt.AccessClaims, jwt.RefreshClaims) string); ok {
		r0 = rf(accessClaims, refreshClaims)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(jwt.AccessClaims, jwt.RefreshClaims) string); ok {
		r1 = rf(accessClaims, refreshClaims)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(jwt.AccessClaims, jwt.RefreshClaims) error); ok {
		r2 = rf(accessClaims, refreshClaims)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Verifier provides a mock function with given fields:
func (_m *JWT) Verifier() func(http.Handler) http.Handler {
	ret := _m.Called()

	var r0 func(http.Handler) http.Handler
	if rf, ok := ret.Get(0).(func() func(http.Handler) http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(http.Handler) http.Handler)
		}
	}

	return r0
}
