// Code generated by MockGen. DO NOT EDIT.
// Source: serdes/interfaces/definitions.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	definitions "github.com/Peersyst/xrpl-go/binary-codec/definitions"
	gomock "github.com/golang/mock/gomock"
)

// MockDefinitions is a mock of Definitions interface.
type MockDefinitions struct {
	ctrl     *gomock.Controller
	recorder *MockDefinitionsMockRecorder
}

// MockDefinitionsMockRecorder is the mock recorder for MockDefinitions.
type MockDefinitionsMockRecorder struct {
	mock *MockDefinitions
}

// NewMockDefinitions creates a new mock instance.
func NewMockDefinitions(ctrl *gomock.Controller) *MockDefinitions {
	mock := &MockDefinitions{ctrl: ctrl}
	mock.recorder = &MockDefinitionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDefinitions) EXPECT() *MockDefinitionsMockRecorder {
	return m.recorder
}

// CreateFieldHeader mocks base method.
func (m *MockDefinitions) CreateFieldHeader(typecode, fieldcode int32) definitions.FieldHeader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFieldHeader", typecode, fieldcode)
	ret0, _ := ret[0].(definitions.FieldHeader)
	return ret0
}

// CreateFieldHeader indicates an expected call of CreateFieldHeader.
func (mr *MockDefinitionsMockRecorder) CreateFieldHeader(typecode, fieldcode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFieldHeader", reflect.TypeOf((*MockDefinitions)(nil).CreateFieldHeader), typecode, fieldcode)
}

// GetFieldHeaderByFieldName mocks base method.
func (m *MockDefinitions) GetFieldHeaderByFieldName(fieldName string) (*definitions.FieldHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFieldHeaderByFieldName", fieldName)
	ret0, _ := ret[0].(*definitions.FieldHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFieldHeaderByFieldName indicates an expected call of GetFieldHeaderByFieldName.
func (mr *MockDefinitionsMockRecorder) GetFieldHeaderByFieldName(fieldName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFieldHeaderByFieldName", reflect.TypeOf((*MockDefinitions)(nil).GetFieldHeaderByFieldName), fieldName)
}

// GetFieldInstanceByFieldName mocks base method.
func (m *MockDefinitions) GetFieldInstanceByFieldName(fieldName string) (*definitions.FieldInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFieldInstanceByFieldName", fieldName)
	ret0, _ := ret[0].(*definitions.FieldInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFieldInstanceByFieldName indicates an expected call of GetFieldInstanceByFieldName.
func (mr *MockDefinitionsMockRecorder) GetFieldInstanceByFieldName(fieldName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFieldInstanceByFieldName", reflect.TypeOf((*MockDefinitions)(nil).GetFieldInstanceByFieldName), fieldName)
}

// GetFieldNameByFieldHeader mocks base method.
func (m *MockDefinitions) GetFieldNameByFieldHeader(fh definitions.FieldHeader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFieldNameByFieldHeader", fh)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFieldNameByFieldHeader indicates an expected call of GetFieldNameByFieldHeader.
func (mr *MockDefinitionsMockRecorder) GetFieldNameByFieldHeader(fh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFieldNameByFieldHeader", reflect.TypeOf((*MockDefinitions)(nil).GetFieldNameByFieldHeader), fh)
}
