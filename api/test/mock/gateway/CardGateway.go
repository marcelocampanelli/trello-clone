// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	entity "github.com/marcelocampanelli/trello-clone/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// CardGateway is an autogenerated mock type for the CardGateway type
type CardGateway struct {
	mock.Mock
}

// Create provides a mock function with given fields: card
func (_m *CardGateway) Create(card *entity.Card) (*entity.Card, error) {
	ret := _m.Called(card)

	var r0 *entity.Card
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Card) (*entity.Card, error)); ok {
		return rf(card)
	}
	if rf, ok := ret.Get(0).(func(*entity.Card) *entity.Card); ok {
		r0 = rf(card)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Card)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Card) error); ok {
		r1 = rf(card)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *CardGateway) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllByList provides a mock function with given fields: listID
func (_m *CardGateway) FindAllByList(listID string) ([]*entity.Card, error) {
	ret := _m.Called(listID)

	var r0 []*entity.Card
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*entity.Card, error)); ok {
		return rf(listID)
	}
	if rf, ok := ret.Get(0).(func(string) []*entity.Card); ok {
		r0 = rf(listID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Card)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(listID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *CardGateway) FindByID(id string) (*entity.Card, error) {
	ret := _m.Called(id)

	var r0 *entity.Card
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Card, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Card); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Card)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, card
func (_m *CardGateway) Update(id string, card *entity.Card) (*entity.Card, error) {
	ret := _m.Called(id, card)

	var r0 *entity.Card
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *entity.Card) (*entity.Card, error)); ok {
		return rf(id, card)
	}
	if rf, ok := ret.Get(0).(func(string, *entity.Card) *entity.Card); ok {
		r0 = rf(id, card)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Card)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *entity.Card) error); ok {
		r1 = rf(id, card)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCardGateway creates a new instance of CardGateway. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCardGateway(t interface {
	mock.TestingT
	Cleanup(func())
}) *CardGateway {
	mock := &CardGateway{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
