// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	models "github.com/corentings/stouuf-bot/models"
	mock "github.com/stretchr/testify/mock"
)

// IKarmaRepository is an autogenerated mock type for the IKarmaRepository type
type IKarmaRepository struct {
	mock.Mock
}

// CreateKarma provides a mock function with given fields: karma
func (_m *IKarmaRepository) CreateKarma(karma models.Karma) error {
	ret := _m.Called(karma)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Karma) error); ok {
		r0 = rf(karma)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetKarma provides a mock function with given fields: userID, guildID
func (_m *IKarmaRepository) GetKarma(userID string, guildID string) (*models.Karma, error) {
	ret := _m.Called(userID, guildID)

	var r0 *models.Karma
	if rf, ok := ret.Get(0).(func(string, string) *models.Karma); ok {
		r0 = rf(userID, guildID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Karma)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, guildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateKarma provides a mock function with given fields: karma
func (_m *IKarmaRepository) UpdateKarma(karma models.Karma) error {
	ret := _m.Called(karma)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Karma) error); ok {
		r0 = rf(karma)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIKarmaRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIKarmaRepository creates a new instance of IKarmaRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIKarmaRepository(t mockConstructorTestingTNewIKarmaRepository) *IKarmaRepository {
	mock := &IKarmaRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
