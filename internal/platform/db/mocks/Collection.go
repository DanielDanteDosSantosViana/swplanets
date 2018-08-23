// Code generated by mockery v1.0.0
package mocks

import mgo "gopkg.in/mgo.v2"
import "github.com/stretchr/testify/mock"

// Collection is an autogenerated mock type for the Collection type
type Collection struct {
	mock.Mock
}

// Count provides a mock function with given fields:
func (_m *Collection) Count() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: query
func (_m *Collection) Find(query interface{}) *mgo.Query {
	ret := _m.Called(query)

	var r0 *mgo.Query
	if rf, ok := ret.Get(0).(func(interface{}) *mgo.Query); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mgo.Query)
		}
	}

	return r0
}

// Insert provides a mock function with given fields: docs
func (_m *Collection) Insert(docs ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, docs...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(docs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remove provides a mock function with given fields: selector
func (_m *Collection) Remove(selector interface{}) error {
	ret := _m.Called(selector)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(selector)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: selector, update
func (_m *Collection) Update(selector interface{}, update interface{}) error {
	ret := _m.Called(selector, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) error); ok {
		r0 = rf(selector, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: selector, update
func (_m *Collection) One(result interface{}) error {
	ret := _m.Called(result)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: selector, update
func (_m *Collection) EnsureIndex(index mgo.Index) error {
	ret := _m.Called(index)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(index)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
