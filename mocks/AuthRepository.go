// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/qvtec/go-app/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// AuthRepository is an autogenerated mock type for the AuthRepository type
type AuthRepository struct {
	mock.Mock
}

// GetPasswordByEmail provides a mock function with given fields: email
func (_m *AuthRepository) GetPasswordByEmail(email string) (string, error) {
	ret := _m.Called(email)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *AuthRepository) GetUserByEmail(email string) (*domain.User, error) {
	ret := _m.Called(email)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePassword provides a mock function with given fields: email, password
func (_m *AuthRepository) UpdatePassword(email string, password string) error {
	ret := _m.Called(email, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAuthRepository creates a new instance of AuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthRepository {
	mock := &AuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}