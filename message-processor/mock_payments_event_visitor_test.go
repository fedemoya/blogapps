// Code generated by mockery v2.14.0. DO NOT EDIT.

package main

import mock "github.com/stretchr/testify/mock"

// MockPaymentsEventVisitor is an autogenerated mock type for the PaymentsEventVisitor type
type MockPaymentsEventVisitor struct {
	mock.Mock
}

type MockPaymentsEventVisitor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPaymentsEventVisitor) EXPECT() *MockPaymentsEventVisitor_Expecter {
	return &MockPaymentsEventVisitor_Expecter{mock: &_m.Mock}
}

// VisitWithdrawalCreated provides a mock function with given fields: withdrawalCreated
func (_m *MockPaymentsEventVisitor) VisitWithdrawalCreated(withdrawalCreated WithdrawalCreated) error {
	ret := _m.Called(withdrawalCreated)

	var r0 error
	if rf, ok := ret.Get(0).(func(WithdrawalCreated) error); ok {
		r0 = rf(withdrawalCreated)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPaymentsEventVisitor_VisitWithdrawalCreated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VisitWithdrawalCreated'
type MockPaymentsEventVisitor_VisitWithdrawalCreated_Call struct {
	*mock.Call
}

// VisitWithdrawalCreated is a helper method to define mock.On call
//  - withdrawalCreated WithdrawalCreated
func (_e *MockPaymentsEventVisitor_Expecter) VisitWithdrawalCreated(withdrawalCreated interface{}) *MockPaymentsEventVisitor_VisitWithdrawalCreated_Call {
	return &MockPaymentsEventVisitor_VisitWithdrawalCreated_Call{Call: _e.mock.On("VisitWithdrawalCreated", withdrawalCreated)}
}

func (_c *MockPaymentsEventVisitor_VisitWithdrawalCreated_Call) Run(run func(withdrawalCreated WithdrawalCreated)) *MockPaymentsEventVisitor_VisitWithdrawalCreated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(WithdrawalCreated))
	})
	return _c
}

func (_c *MockPaymentsEventVisitor_VisitWithdrawalCreated_Call) Return(_a0 error) *MockPaymentsEventVisitor_VisitWithdrawalCreated_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockPaymentsEventVisitor interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPaymentsEventVisitor creates a new instance of MockPaymentsEventVisitor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPaymentsEventVisitor(t mockConstructorTestingTNewMockPaymentsEventVisitor) *MockPaymentsEventVisitor {
	mock := &MockPaymentsEventVisitor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}