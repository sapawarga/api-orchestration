// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sapawarga/api-orchestration/repository (interfaces: PhonebookI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	phonebook "github.com/sapawarga/api-orchestration/model/phonebook"
	reflect "reflect"
)

// MockPhonebookI is a mock of PhonebookI interface
type MockPhonebookI struct {
	ctrl     *gomock.Controller
	recorder *MockPhonebookIMockRecorder
}

// MockPhonebookIMockRecorder is the mock recorder for MockPhonebookI
type MockPhonebookIMockRecorder struct {
	mock *MockPhonebookI
}

// NewMockPhonebookI creates a new mock instance
func NewMockPhonebookI(ctrl *gomock.Controller) *MockPhonebookI {
	mock := &MockPhonebookI{ctrl: ctrl}
	mock.recorder = &MockPhonebookIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPhonebookI) EXPECT() *MockPhonebookIMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockPhonebookI) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPhonebookIMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPhonebookI)(nil).Delete), arg0, arg1)
}

// GetDetail mocks base method
func (m *MockPhonebookI) GetDetail(arg0 context.Context, arg1 int64) (*phonebook.PhoneDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetail", arg0, arg1)
	ret0, _ := ret[0].(*phonebook.PhoneDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetail indicates an expected call of GetDetail
func (mr *MockPhonebookIMockRecorder) GetDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetail", reflect.TypeOf((*MockPhonebookI)(nil).GetDetail), arg0, arg1)
}

// GetList mocks base method
func (m *MockPhonebookI) GetList(arg0 context.Context, arg1 *phonebook.GetListRequest) (*phonebook.GetListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", arg0, arg1)
	ret0, _ := ret[0].(*phonebook.GetListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList
func (mr *MockPhonebookIMockRecorder) GetList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockPhonebookI)(nil).GetList), arg0, arg1)
}

// Insert mocks base method
func (m *MockPhonebookI) Insert(arg0 context.Context, arg1 *phonebook.AddPhonebookRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockPhonebookIMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPhonebookI)(nil).Insert), arg0, arg1)
}

// Update mocks base method
func (m *MockPhonebookI) Update(arg0 context.Context, arg1 *phonebook.UpdatePhonebookRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockPhonebookIMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPhonebookI)(nil).Update), arg0, arg1)
}
