// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/utils/fcp (interfaces: FcpReconcileUtils)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/mock_utils/mock_fcp/mock_reconcile_utils.go github.com/netapp/trident/utils/fcp FcpReconcileUtils
//

// Package mock_fcp is a generated GoMock package.
package mock_fcp

import (
	context "context"
	reflect "reflect"

	models "github.com/netapp/trident/utils/models"
	gomock "go.uber.org/mock/gomock"
)

// MockFcpReconcileUtils is a mock of FcpReconcileUtils interface.
type MockFcpReconcileUtils struct {
	ctrl     *gomock.Controller
	recorder *MockFcpReconcileUtilsMockRecorder
}

// MockFcpReconcileUtilsMockRecorder is the mock recorder for MockFcpReconcileUtils.
type MockFcpReconcileUtilsMockRecorder struct {
	mock *MockFcpReconcileUtils
}

// NewMockFcpReconcileUtils creates a new mock instance.
func NewMockFcpReconcileUtils(ctrl *gomock.Controller) *MockFcpReconcileUtils {
	mock := &MockFcpReconcileUtils{ctrl: ctrl}
	mock.recorder = &MockFcpReconcileUtilsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFcpReconcileUtils) EXPECT() *MockFcpReconcileUtilsMockRecorder {
	return m.recorder
}

// GetDevicesForLUN mocks base method.
func (m *MockFcpReconcileUtils) GetDevicesForLUN(arg0 []string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesForLUN", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesForLUN indicates an expected call of GetDevicesForLUN.
func (mr *MockFcpReconcileUtilsMockRecorder) GetDevicesForLUN(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesForLUN", reflect.TypeOf((*MockFcpReconcileUtils)(nil).GetDevicesForLUN), arg0)
}

// GetFCPHostSessionMapForTarget mocks base method.
func (m *MockFcpReconcileUtils) GetFCPHostSessionMapForTarget(arg0 context.Context, arg1 string) []map[string]int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFCPHostSessionMapForTarget", arg0, arg1)
	ret0, _ := ret[0].([]map[string]int)
	return ret0
}

// GetFCPHostSessionMapForTarget indicates an expected call of GetFCPHostSessionMapForTarget.
func (mr *MockFcpReconcileUtilsMockRecorder) GetFCPHostSessionMapForTarget(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFCPHostSessionMapForTarget", reflect.TypeOf((*MockFcpReconcileUtils)(nil).GetFCPHostSessionMapForTarget), arg0, arg1)
}

// GetSysfsBlockDirsForLUN mocks base method.
func (m *MockFcpReconcileUtils) GetSysfsBlockDirsForLUN(arg0 int, arg1 []map[string]int) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSysfsBlockDirsForLUN", arg0, arg1)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetSysfsBlockDirsForLUN indicates an expected call of GetSysfsBlockDirsForLUN.
func (mr *MockFcpReconcileUtilsMockRecorder) GetSysfsBlockDirsForLUN(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSysfsBlockDirsForLUN", reflect.TypeOf((*MockFcpReconcileUtils)(nil).GetSysfsBlockDirsForLUN), arg0, arg1)
}

// ReconcileFCPVolumeInfo mocks base method.
func (m *MockFcpReconcileUtils) ReconcileFCPVolumeInfo(arg0 context.Context, arg1 *models.VolumeTrackingInfo) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileFCPVolumeInfo", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileFCPVolumeInfo indicates an expected call of ReconcileFCPVolumeInfo.
func (mr *MockFcpReconcileUtilsMockRecorder) ReconcileFCPVolumeInfo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileFCPVolumeInfo", reflect.TypeOf((*MockFcpReconcileUtils)(nil).ReconcileFCPVolumeInfo), arg0, arg1)
}
