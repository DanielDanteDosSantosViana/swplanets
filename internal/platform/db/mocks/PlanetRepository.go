// Code generated by mockery v1.0.0
package mocks

import (
	"github.com/DanielDanteDosSantosViana/swplanets/internal/planet"
	"github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type PlanetRepository struct {
	mock.Mock
}

func (_m *PlanetRepository) Store(_a0 *planet.Planet) (*planet.Planet, error) {
	ret := _m.Called(_a0)

	var r0 *planet.Planet
	if rf, ok := ret.Get(0).(func(*planet.Planet) *planet.Planet); ok {
		r0 = rf(_a0)
	} else if ret.Get(0) != nil {
		r0 = ret.Get(0).(*planet.Planet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*planet.Planet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *PlanetRepository) Remove(_a0 string) error {
	ret := _m.Called(_a0)
	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}
	return r0
}

func (_m *PlanetRepository) List() ([]planet.Planet, error) {
	ret := _m.Called()

	var r0 []planet.Planet
	if rf, ok := ret.Get(0).(func() []planet.Planet); ok {
		r0 = rf()
	} else if ret.Get(0) != nil {
		r0 = ret.Get(0).([]planet.Planet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *PlanetRepository) GetById(id string) (*planet.Planet, error) {
	ret := _m.Called()

	var r0 *planet.Planet
	if rf, ok := ret.Get(0).(func() *planet.Planet); ok {
		r0 = rf()
	} else if ret.Get(0) != nil {
		r0 = ret.Get(0).(*planet.Planet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *PlanetRepository) GetByName(id string) ([]planet.Planet, error) {
	ret := _m.Called()

	var r0 []planet.Planet
	if rf, ok := ret.Get(0).(func() []planet.Planet); ok {
		r0 = rf()
	} else if ret.Get(0) != nil {
		r0 = ret.Get(0).([]planet.Planet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
