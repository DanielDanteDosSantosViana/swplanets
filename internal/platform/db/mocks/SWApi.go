package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

type Client struct {
	mock.Mock
}

func (_m *Client) GetNumberOfAppearancesByPlanetName(_a0 string) (int, error) {

	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(search string) int); ok {
		r0 = rf(_a0)
	} else if ret.Get(0) != nil {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(search string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
