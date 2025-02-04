// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/utils/devices (interfaces: Devices)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/mock_utils/mock_devices/mock_devices_client.go github.com/netapp/trident/utils/devices Devices
//

// Package mock_devices is a generated GoMock package.
package mock_devices

import (
	context "context"
	reflect "reflect"
	time "time"

	models "github.com/netapp/trident/utils/models"
	gomock "go.uber.org/mock/gomock"
)

// MockDevices is a mock of Devices interface.
type MockDevices struct {
	ctrl     *gomock.Controller
	recorder *MockDevicesMockRecorder
}

// MockDevicesMockRecorder is the mock recorder for MockDevices.
type MockDevicesMockRecorder struct {
	mock *MockDevices
}

// NewMockDevices creates a new mock instance.
func NewMockDevices(ctrl *gomock.Controller) *MockDevices {
	mock := &MockDevices{ctrl: ctrl}
	mock.recorder = &MockDevicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDevices) EXPECT() *MockDevicesMockRecorder {
	return m.recorder
}

// CloseLUKSDevice mocks base method.
func (m *MockDevices) CloseLUKSDevice(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseLUKSDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseLUKSDevice indicates an expected call of CloseLUKSDevice.
func (mr *MockDevicesMockRecorder) CloseLUKSDevice(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseLUKSDevice", reflect.TypeOf((*MockDevices)(nil).CloseLUKSDevice), arg0, arg1)
}

// EnsureDeviceReadable mocks base method.
func (m *MockDevices) EnsureDeviceReadable(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureDeviceReadable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureDeviceReadable indicates an expected call of EnsureDeviceReadable.
func (mr *MockDevicesMockRecorder) EnsureDeviceReadable(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureDeviceReadable", reflect.TypeOf((*MockDevices)(nil).EnsureDeviceReadable), arg0, arg1)
}

// EnsureLUKSDeviceClosed mocks base method.
func (m *MockDevices) EnsureLUKSDeviceClosed(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureLUKSDeviceClosed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureLUKSDeviceClosed indicates an expected call of EnsureLUKSDeviceClosed.
func (mr *MockDevicesMockRecorder) EnsureLUKSDeviceClosed(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureLUKSDeviceClosed", reflect.TypeOf((*MockDevices)(nil).EnsureLUKSDeviceClosed), arg0, arg1)
}

// EnsureLUKSDeviceClosedWithMaxWaitLimit mocks base method.
func (m *MockDevices) EnsureLUKSDeviceClosedWithMaxWaitLimit(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureLUKSDeviceClosedWithMaxWaitLimit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureLUKSDeviceClosedWithMaxWaitLimit indicates an expected call of EnsureLUKSDeviceClosedWithMaxWaitLimit.
func (mr *MockDevicesMockRecorder) EnsureLUKSDeviceClosedWithMaxWaitLimit(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureLUKSDeviceClosedWithMaxWaitLimit", reflect.TypeOf((*MockDevices)(nil).EnsureLUKSDeviceClosedWithMaxWaitLimit), arg0, arg1)
}

// FindDevicesForMultipathDevice mocks base method.
func (m *MockDevices) FindDevicesForMultipathDevice(arg0 context.Context, arg1 string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDevicesForMultipathDevice", arg0, arg1)
	ret0, _ := ret[0].([]string)
	return ret0
}

// FindDevicesForMultipathDevice indicates an expected call of FindDevicesForMultipathDevice.
func (mr *MockDevicesMockRecorder) FindDevicesForMultipathDevice(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDevicesForMultipathDevice", reflect.TypeOf((*MockDevices)(nil).FindDevicesForMultipathDevice), arg0, arg1)
}

// FindMultipathDeviceForDevice mocks base method.
func (m *MockDevices) FindMultipathDeviceForDevice(arg0 context.Context, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMultipathDeviceForDevice", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// FindMultipathDeviceForDevice indicates an expected call of FindMultipathDeviceForDevice.
func (mr *MockDevicesMockRecorder) FindMultipathDeviceForDevice(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMultipathDeviceForDevice", reflect.TypeOf((*MockDevices)(nil).FindMultipathDeviceForDevice), arg0, arg1)
}

// FlushDevice mocks base method.
func (m *MockDevices) FlushDevice(arg0 context.Context, arg1 *models.ScsiDeviceInfo, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushDevice", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushDevice indicates an expected call of FlushDevice.
func (mr *MockDevicesMockRecorder) FlushDevice(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushDevice", reflect.TypeOf((*MockDevices)(nil).FlushDevice), arg0, arg1, arg2)
}

// FlushOneDevice mocks base method.
func (m *MockDevices) FlushOneDevice(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushOneDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushOneDevice indicates an expected call of FlushOneDevice.
func (mr *MockDevicesMockRecorder) FlushOneDevice(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushOneDevice", reflect.TypeOf((*MockDevices)(nil).FlushOneDevice), arg0, arg1)
}

// GetDeviceFSType mocks base method.
func (m *MockDevices) GetDeviceFSType(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceFSType", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceFSType indicates an expected call of GetDeviceFSType.
func (mr *MockDevicesMockRecorder) GetDeviceFSType(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceFSType", reflect.TypeOf((*MockDevices)(nil).GetDeviceFSType), arg0, arg1)
}

// GetDiskSize mocks base method.
func (m *MockDevices) GetDiskSize(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskSize", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskSize indicates an expected call of GetDiskSize.
func (mr *MockDevicesMockRecorder) GetDiskSize(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskSize", reflect.TypeOf((*MockDevices)(nil).GetDiskSize), arg0, arg1)
}

// GetLUKSDeviceForMultipathDevice mocks base method.
func (m *MockDevices) GetLUKSDeviceForMultipathDevice(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLUKSDeviceForMultipathDevice", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLUKSDeviceForMultipathDevice indicates an expected call of GetLUKSDeviceForMultipathDevice.
func (mr *MockDevicesMockRecorder) GetLUKSDeviceForMultipathDevice(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLUKSDeviceForMultipathDevice", reflect.TypeOf((*MockDevices)(nil).GetLUKSDeviceForMultipathDevice), arg0)
}

// GetLunSerial mocks base method.
func (m *MockDevices) GetLunSerial(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLunSerial", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLunSerial indicates an expected call of GetLunSerial.
func (mr *MockDevicesMockRecorder) GetLunSerial(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLunSerial", reflect.TypeOf((*MockDevices)(nil).GetLunSerial), arg0, arg1)
}

// GetMultipathDeviceUUID mocks base method.
func (m *MockDevices) GetMultipathDeviceUUID(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMultipathDeviceUUID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultipathDeviceUUID indicates an expected call of GetMultipathDeviceUUID.
func (mr *MockDevicesMockRecorder) GetMultipathDeviceUUID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultipathDeviceUUID", reflect.TypeOf((*MockDevices)(nil).GetMultipathDeviceUUID), arg0)
}

// IsDeviceUnformatted mocks base method.
func (m *MockDevices) IsDeviceUnformatted(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDeviceUnformatted", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsDeviceUnformatted indicates an expected call of IsDeviceUnformatted.
func (mr *MockDevicesMockRecorder) IsDeviceUnformatted(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDeviceUnformatted", reflect.TypeOf((*MockDevices)(nil).IsDeviceUnformatted), arg0, arg1)
}

// ListAllDevices mocks base method.
func (m *MockDevices) ListAllDevices(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ListAllDevices", arg0)
}

// ListAllDevices indicates an expected call of ListAllDevices.
func (mr *MockDevicesMockRecorder) ListAllDevices(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllDevices", reflect.TypeOf((*MockDevices)(nil).ListAllDevices), arg0)
}

// MultipathFlushDevice mocks base method.
func (m *MockDevices) MultipathFlushDevice(arg0 context.Context, arg1 *models.ScsiDeviceInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultipathFlushDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MultipathFlushDevice indicates an expected call of MultipathFlushDevice.
func (mr *MockDevicesMockRecorder) MultipathFlushDevice(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultipathFlushDevice", reflect.TypeOf((*MockDevices)(nil).MultipathFlushDevice), arg0, arg1)
}

// RemoveDevice mocks base method.
func (m *MockDevices) RemoveDevice(arg0 context.Context, arg1 []string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDevice", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDevice indicates an expected call of RemoveDevice.
func (mr *MockDevicesMockRecorder) RemoveDevice(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDevice", reflect.TypeOf((*MockDevices)(nil).RemoveDevice), arg0, arg1, arg2)
}

// RemoveMultipathDeviceMapping mocks base method.
func (m *MockDevices) RemoveMultipathDeviceMapping(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMultipathDeviceMapping", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMultipathDeviceMapping indicates an expected call of RemoveMultipathDeviceMapping.
func (mr *MockDevicesMockRecorder) RemoveMultipathDeviceMapping(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMultipathDeviceMapping", reflect.TypeOf((*MockDevices)(nil).RemoveMultipathDeviceMapping), arg0, arg1)
}

// RemoveMultipathDeviceMappingWithRetries mocks base method.
func (m *MockDevices) RemoveMultipathDeviceMappingWithRetries(arg0 context.Context, arg1 string, arg2 uint64, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMultipathDeviceMappingWithRetries", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMultipathDeviceMappingWithRetries indicates an expected call of RemoveMultipathDeviceMappingWithRetries.
func (mr *MockDevicesMockRecorder) RemoveMultipathDeviceMappingWithRetries(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMultipathDeviceMappingWithRetries", reflect.TypeOf((*MockDevices)(nil).RemoveMultipathDeviceMappingWithRetries), arg0, arg1, arg2, arg3)
}

// ScanTargetLUN mocks base method.
func (m *MockDevices) ScanTargetLUN(arg0 context.Context, arg1 []models.ScsiDeviceAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanTargetLUN", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanTargetLUN indicates an expected call of ScanTargetLUN.
func (mr *MockDevicesMockRecorder) ScanTargetLUN(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanTargetLUN", reflect.TypeOf((*MockDevices)(nil).ScanTargetLUN), arg0, arg1)
}

// VerifyMultipathDevice mocks base method.
func (m *MockDevices) VerifyMultipathDevice(arg0 context.Context, arg1 *models.VolumePublishInfo, arg2 []models.VolumePublishInfo, arg3 *models.ScsiDeviceInfo) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyMultipathDevice", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyMultipathDevice indicates an expected call of VerifyMultipathDevice.
func (mr *MockDevicesMockRecorder) VerifyMultipathDevice(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyMultipathDevice", reflect.TypeOf((*MockDevices)(nil).VerifyMultipathDevice), arg0, arg1, arg2, arg3)
}

// VerifyMultipathDeviceSize mocks base method.
func (m *MockDevices) VerifyMultipathDeviceSize(arg0 context.Context, arg1, arg2 string) (int64, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyMultipathDeviceSize", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VerifyMultipathDeviceSize indicates an expected call of VerifyMultipathDeviceSize.
func (mr *MockDevicesMockRecorder) VerifyMultipathDeviceSize(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyMultipathDeviceSize", reflect.TypeOf((*MockDevices)(nil).VerifyMultipathDeviceSize), arg0, arg1, arg2)
}

// WaitForDevice mocks base method.
func (m *MockDevices) WaitForDevice(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForDevice", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForDevice indicates an expected call of WaitForDevice.
func (mr *MockDevicesMockRecorder) WaitForDevice(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForDevice", reflect.TypeOf((*MockDevices)(nil).WaitForDevice), arg0, arg1)
}

// WaitForDevicesRemoval mocks base method.
func (m *MockDevices) WaitForDevicesRemoval(arg0 context.Context, arg1 string, arg2 []string, arg3 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForDevicesRemoval", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForDevicesRemoval indicates an expected call of WaitForDevicesRemoval.
func (mr *MockDevicesMockRecorder) WaitForDevicesRemoval(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForDevicesRemoval", reflect.TypeOf((*MockDevices)(nil).WaitForDevicesRemoval), arg0, arg1, arg2, arg3)
}
