// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/simimpact/srsim/pkg/engine/attribute (interfaces: Modifier)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	info "github.com/simimpact/srsim/pkg/engine/info"
	key "github.com/simimpact/srsim/pkg/key"
)

// MockAttribute is a mock of Modifier interface.
type MockAttribute struct {
	ctrl     *gomock.Controller
	recorder *MockAttributeMockRecorder
}

// MockAttributeMockRecorder is the mock recorder for MockAttribute.
type MockAttributeMockRecorder struct {
	mock *MockAttribute
}

// NewMockAttribute creates a new mock instance.
func NewMockAttribute(ctrl *gomock.Controller) *MockAttribute {
	mock := &MockAttribute{ctrl: ctrl}
	mock.recorder = &MockAttributeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAttribute) EXPECT() *MockAttributeMockRecorder {
	return m.recorder
}

// AddTarget mocks base method.
func (m *MockAttribute) AddTarget(arg0 key.TargetID, arg1 info.Attributes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTarget", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTarget indicates an expected call of AddTarget.
func (mr *MockAttributeMockRecorder) AddTarget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTarget", reflect.TypeOf((*MockAttribute)(nil).AddTarget), arg0, arg1)
}

// Energy mocks base method.
func (m *MockAttribute) Energy(arg0 key.TargetID) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Energy", arg0)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Energy indicates an expected call of Energy.
func (mr *MockAttributeMockRecorder) Energy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Energy", reflect.TypeOf((*MockAttribute)(nil).Energy), arg0)
}

// EnergyRatio mocks base method.
func (m *MockAttribute) EnergyRatio(arg0 key.TargetID) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnergyRatio", arg0)
	ret0, _ := ret[0].(float64)
	return ret0
}

// EnergyRatio indicates an expected call of EnergyRatio.
func (mr *MockAttributeMockRecorder) EnergyRatio(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnergyRatio", reflect.TypeOf((*MockAttribute)(nil).EnergyRatio), arg0)
}

// HPRatio mocks base method.
func (m *MockAttribute) HPRatio(arg0 key.TargetID) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HPRatio", arg0)
	ret0, _ := ret[0].(float64)
	return ret0
}

// HPRatio indicates an expected call of HPRatio.
func (mr *MockAttributeMockRecorder) HPRatio(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HPRatio", reflect.TypeOf((*MockAttribute)(nil).HPRatio), arg0)
}

// IsAlive mocks base method.
func (m *MockAttribute) IsAlive(arg0 key.TargetID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAlive", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAlive indicates an expected call of IsAlive.
func (mr *MockAttributeMockRecorder) IsAlive(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAlive", reflect.TypeOf((*MockAttribute)(nil).IsAlive), arg0)
}

// MaxEnergy mocks base method.
func (m *MockAttribute) MaxEnergy(arg0 key.TargetID) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxEnergy", arg0)
	ret0, _ := ret[0].(float64)
	return ret0
}

// MaxEnergy indicates an expected call of MaxEnergy.
func (mr *MockAttributeMockRecorder) MaxEnergy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxEnergy", reflect.TypeOf((*MockAttribute)(nil).MaxEnergy), arg0)
}

// ModifyEnergy mocks base method.
func (m *MockAttribute) ModifyEnergy(arg0 key.TargetID, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyEnergy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyEnergy indicates an expected call of ModifyEnergy.
func (mr *MockAttributeMockRecorder) ModifyEnergy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyEnergy", reflect.TypeOf((*MockAttribute)(nil).ModifyEnergy), arg0, arg1)
}

// ModifyEnergyFixed mocks base method.
func (m *MockAttribute) ModifyEnergyFixed(arg0 key.TargetID, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyEnergyFixed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyEnergyFixed indicates an expected call of ModifyEnergyFixed.
func (mr *MockAttributeMockRecorder) ModifyEnergyFixed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyEnergyFixed", reflect.TypeOf((*MockAttribute)(nil).ModifyEnergyFixed), arg0, arg1)
}

// ModifyHPByAmount mocks base method.
func (m *MockAttribute) ModifyHPByAmount(arg0, arg1 key.TargetID, arg2 float64, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyHPByAmount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyHPByAmount indicates an expected call of ModifyHPByAmount.
func (mr *MockAttributeMockRecorder) ModifyHPByAmount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyHPByAmount", reflect.TypeOf((*MockAttribute)(nil).ModifyHPByAmount), arg0, arg1, arg2, arg3)
}

// ModifyHPByRatio mocks base method.
func (m *MockAttribute) ModifyHPByRatio(arg0, arg1 key.TargetID, arg2 info.ModifyHPByRatio, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyHPByRatio", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyHPByRatio indicates an expected call of ModifyHPByRatio.
func (mr *MockAttributeMockRecorder) ModifyHPByRatio(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyHPByRatio", reflect.TypeOf((*MockAttribute)(nil).ModifyHPByRatio), arg0, arg1, arg2, arg3)
}

// ModifyStance mocks base method.
func (m *MockAttribute) ModifyStance(arg0, arg1 key.TargetID, arg2 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyStance", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyStance indicates an expected call of ModifyStance.
func (mr *MockAttributeMockRecorder) ModifyStance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyStance", reflect.TypeOf((*MockAttribute)(nil).ModifyStance), arg0, arg1, arg2)
}

// SetEnergy mocks base method.
func (m *MockAttribute) SetEnergy(arg0 key.TargetID, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEnergy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEnergy indicates an expected call of SetEnergy.
func (mr *MockAttributeMockRecorder) SetEnergy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEnergy", reflect.TypeOf((*MockAttribute)(nil).SetEnergy), arg0, arg1)
}

// SetHP mocks base method.
func (m *MockAttribute) SetHP(arg0, arg1 key.TargetID, arg2 float64, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHP", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHP indicates an expected call of SetHP.
func (mr *MockAttributeMockRecorder) SetHP(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHP", reflect.TypeOf((*MockAttribute)(nil).SetHP), arg0, arg1, arg2, arg3)
}

// SetStance mocks base method.
func (m *MockAttribute) SetStance(arg0, arg1 key.TargetID, arg2 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStance", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStance indicates an expected call of SetStance.
func (mr *MockAttributeMockRecorder) SetStance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStance", reflect.TypeOf((*MockAttribute)(nil).SetStance), arg0, arg1, arg2)
}

// Stance mocks base method.
func (m *MockAttribute) Stance(arg0 key.TargetID) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stance", arg0)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Stance indicates an expected call of Stance.
func (mr *MockAttributeMockRecorder) Stance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stance", reflect.TypeOf((*MockAttribute)(nil).Stance), arg0)
}

// Stats mocks base method.
func (m *MockAttribute) Stats(arg0 key.TargetID) *info.Stats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats", arg0)
	ret0, _ := ret[0].(*info.Stats)
	return ret0
}

// Stats indicates an expected call of Stats.
func (mr *MockAttributeMockRecorder) Stats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockAttribute)(nil).Stats), arg0)
}
