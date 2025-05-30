// Code generated by MockGen. DO NOT EDIT.
// Source: types/interfaces/binary_parser.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	definitions "github.com/Peersyst/xrpl-go/binary-codec/definitions"
	gomock "github.com/golang/mock/gomock"
)

// MockBinaryParser is a mock of BinaryParser interface.
type MockBinaryParser struct {
	ctrl     *gomock.Controller
	recorder *MockBinaryParserMockRecorder
}

// MockBinaryParserMockRecorder is the mock recorder for MockBinaryParser.
type MockBinaryParserMockRecorder struct {
	mock *MockBinaryParser
}

// NewMockBinaryParser creates a new mock instance.
func NewMockBinaryParser(ctrl *gomock.Controller) *MockBinaryParser {
	mock := &MockBinaryParser{ctrl: ctrl}
	mock.recorder = &MockBinaryParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBinaryParser) EXPECT() *MockBinaryParserMockRecorder {
	return m.recorder
}

// HasMore mocks base method.
func (m *MockBinaryParser) HasMore() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasMore")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasMore indicates an expected call of HasMore.
func (mr *MockBinaryParserMockRecorder) HasMore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasMore", reflect.TypeOf((*MockBinaryParser)(nil).HasMore))
}

// Peek mocks base method.
func (m *MockBinaryParser) Peek() (byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peek")
	ret0, _ := ret[0].(byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Peek indicates an expected call of Peek.
func (mr *MockBinaryParserMockRecorder) Peek() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peek", reflect.TypeOf((*MockBinaryParser)(nil).Peek))
}

// ReadByte mocks base method.
func (m *MockBinaryParser) ReadByte() (byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByte")
	ret0, _ := ret[0].(byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByte indicates an expected call of ReadByte.
func (mr *MockBinaryParserMockRecorder) ReadByte() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByte", reflect.TypeOf((*MockBinaryParser)(nil).ReadByte))
}

// ReadBytes mocks base method.
func (m *MockBinaryParser) ReadBytes(n int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadBytes", n)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadBytes indicates an expected call of ReadBytes.
func (mr *MockBinaryParserMockRecorder) ReadBytes(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadBytes", reflect.TypeOf((*MockBinaryParser)(nil).ReadBytes), n)
}

// ReadField mocks base method.
func (m *MockBinaryParser) ReadField() (*definitions.FieldInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadField")
	ret0, _ := ret[0].(*definitions.FieldInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadField indicates an expected call of ReadField.
func (mr *MockBinaryParserMockRecorder) ReadField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadField", reflect.TypeOf((*MockBinaryParser)(nil).ReadField))
}

// ReadVariableLength mocks base method.
func (m *MockBinaryParser) ReadVariableLength() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadVariableLength")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadVariableLength indicates an expected call of ReadVariableLength.
func (mr *MockBinaryParserMockRecorder) ReadVariableLength() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadVariableLength", reflect.TypeOf((*MockBinaryParser)(nil).ReadVariableLength))
}
