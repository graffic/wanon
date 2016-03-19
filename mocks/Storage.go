package mocks

import "github.com/stretchr/testify/mock"

import "github.com/mkilling/goejdb"

type Storage struct {
	mock.Mock
}

func (_m *Storage) CreateColl(colname string, opts *goejdb.EjCollOpts) (*goejdb.EjColl, *goejdb.EjdbError) {
	ret := _m.Called(colname, opts)

	var r0 *goejdb.EjColl
	if rf, ok := ret.Get(0).(func(string, *goejdb.EjCollOpts) *goejdb.EjColl); ok {
		r0 = rf(colname, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*goejdb.EjColl)
		}
	}

	var r1 *goejdb.EjdbError
	if rf, ok := ret.Get(1).(func(string, *goejdb.EjCollOpts) *goejdb.EjdbError); ok {
		r1 = rf(colname, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*goejdb.EjdbError)
		}
	}

	return r0, r1
}
func (_m *Storage) CreateQuery(query string, queries ...string) (*goejdb.EjQuery, *goejdb.EjdbError) {
	ret := _m.Called(query, queries)

	var r0 *goejdb.EjQuery
	if rf, ok := ret.Get(0).(func(string, ...string) *goejdb.EjQuery); ok {
		r0 = rf(query, queries...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*goejdb.EjQuery)
		}
	}

	var r1 *goejdb.EjdbError
	if rf, ok := ret.Get(1).(func(string, ...string) *goejdb.EjdbError); ok {
		r1 = rf(query, queries...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*goejdb.EjdbError)
		}
	}

	return r0, r1
}
func (_m *Storage) GetColl(colname string) (*goejdb.EjColl, *goejdb.EjdbError) {
	ret := _m.Called(colname)

	var r0 *goejdb.EjColl
	if rf, ok := ret.Get(0).(func(string) *goejdb.EjColl); ok {
		r0 = rf(colname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*goejdb.EjColl)
		}
	}

	var r1 *goejdb.EjdbError
	if rf, ok := ret.Get(1).(func(string) *goejdb.EjdbError); ok {
		r1 = rf(colname)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*goejdb.EjdbError)
		}
	}

	return r0, r1
}
func (_m *Storage) Meta() ([]byte, *goejdb.EjdbError) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 *goejdb.EjdbError
	if rf, ok := ret.Get(1).(func() *goejdb.EjdbError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*goejdb.EjdbError)
		}
	}

	return r0, r1
}
