package manage

import (
	"github.com/graffic/goejdb"
	"github.com/stretchr/testify/mock"
)

// This is an autogenerated mock type for the Storage type
type MockStorage struct {
	mock.Mock
}

// CreateColl provides a mock function with given fields: colname, opts
func (_m *MockStorage) CreateColl(colname string, opts *goejdb.EjCollOpts) (*goejdb.EjColl, *goejdb.EjdbError) {
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

// CreateQuery provides a mock function with given fields: query, queries
func (_m *MockStorage) CreateQuery(query string, queries ...string) (*goejdb.EjQuery, *goejdb.EjdbError) {
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

// GetColl provides a mock function with given fields: colname
func (_m *MockStorage) GetColl(colname string) (*goejdb.EjColl, *goejdb.EjdbError) {
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

// Meta provides a mock function with given fields:
func (_m *MockStorage) Meta() ([]byte, *goejdb.EjdbError) {
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

// RmColl provides a mock function with given fields: colname, unlinkfile
func (_m *MockStorage) RmColl(colname string, unlinkfile bool) (bool, *goejdb.EjdbError) {
	ret := _m.Called(colname, unlinkfile)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, bool) bool); ok {
		r0 = rf(colname, unlinkfile)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *goejdb.EjdbError
	if rf, ok := ret.Get(1).(func(string, bool) *goejdb.EjdbError); ok {
		r1 = rf(colname, unlinkfile)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*goejdb.EjdbError)
		}
	}

	return r0, r1
}

// Sync provides a mock function with given fields:
func (_m *MockStorage) Sync() (bool, *goejdb.EjdbError) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
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
