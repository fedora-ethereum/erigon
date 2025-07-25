// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erigontech/erigon/cl/das (interfaces: PeerDas)
//
// Generated by this command:
//
//	mockgen -typed=true -destination=mock_services/peer_das_mock.go -package=mock_services . PeerDas
//

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	common "github.com/erigontech/erigon-lib/common"
	cltypes "github.com/erigontech/erigon/cl/cltypes"
	peerdasstate "github.com/erigontech/erigon/cl/das/state"
	gomock "go.uber.org/mock/gomock"
)

// MockPeerDas is a mock of PeerDas interface.
type MockPeerDas struct {
	ctrl     *gomock.Controller
	recorder *MockPeerDasMockRecorder
	isgomock struct{}
}

// MockPeerDasMockRecorder is the mock recorder for MockPeerDas.
type MockPeerDasMockRecorder struct {
	mock *MockPeerDas
}

// NewMockPeerDas creates a new mock instance.
func NewMockPeerDas(ctrl *gomock.Controller) *MockPeerDas {
	mock := &MockPeerDas{ctrl: ctrl}
	mock.recorder = &MockPeerDasMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeerDas) EXPECT() *MockPeerDasMockRecorder {
	return m.recorder
}

// DownloadColumnsAndRecoverBlobs mocks base method.
func (m *MockPeerDas) DownloadColumnsAndRecoverBlobs(ctx context.Context, blocks []*cltypes.SignedBeaconBlock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadColumnsAndRecoverBlobs", ctx, blocks)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadColumnsAndRecoverBlobs indicates an expected call of DownloadColumnsAndRecoverBlobs.
func (mr *MockPeerDasMockRecorder) DownloadColumnsAndRecoverBlobs(ctx, blocks any) *MockPeerDasDownloadColumnsAndRecoverBlobsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadColumnsAndRecoverBlobs", reflect.TypeOf((*MockPeerDas)(nil).DownloadColumnsAndRecoverBlobs), ctx, blocks)
	return &MockPeerDasDownloadColumnsAndRecoverBlobsCall{Call: call}
}

// MockPeerDasDownloadColumnsAndRecoverBlobsCall wrap *gomock.Call
type MockPeerDasDownloadColumnsAndRecoverBlobsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPeerDasDownloadColumnsAndRecoverBlobsCall) Return(arg0 error) *MockPeerDasDownloadColumnsAndRecoverBlobsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPeerDasDownloadColumnsAndRecoverBlobsCall) Do(f func(context.Context, []*cltypes.SignedBeaconBlock) error) *MockPeerDasDownloadColumnsAndRecoverBlobsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPeerDasDownloadColumnsAndRecoverBlobsCall) DoAndReturn(f func(context.Context, []*cltypes.SignedBeaconBlock) error) *MockPeerDasDownloadColumnsAndRecoverBlobsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DownloadOnlyCustodyColumns mocks base method.
func (m *MockPeerDas) DownloadOnlyCustodyColumns(ctx context.Context, blocks []*cltypes.SignedBeaconBlock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadOnlyCustodyColumns", ctx, blocks)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadOnlyCustodyColumns indicates an expected call of DownloadOnlyCustodyColumns.
func (mr *MockPeerDasMockRecorder) DownloadOnlyCustodyColumns(ctx, blocks any) *MockPeerDasDownloadOnlyCustodyColumnsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadOnlyCustodyColumns", reflect.TypeOf((*MockPeerDas)(nil).DownloadOnlyCustodyColumns), ctx, blocks)
	return &MockPeerDasDownloadOnlyCustodyColumnsCall{Call: call}
}

// MockPeerDasDownloadOnlyCustodyColumnsCall wrap *gomock.Call
type MockPeerDasDownloadOnlyCustodyColumnsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPeerDasDownloadOnlyCustodyColumnsCall) Return(arg0 error) *MockPeerDasDownloadOnlyCustodyColumnsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPeerDasDownloadOnlyCustodyColumnsCall) Do(f func(context.Context, []*cltypes.SignedBeaconBlock) error) *MockPeerDasDownloadOnlyCustodyColumnsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPeerDasDownloadOnlyCustodyColumnsCall) DoAndReturn(f func(context.Context, []*cltypes.SignedBeaconBlock) error) *MockPeerDasDownloadOnlyCustodyColumnsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsArchivedMode mocks base method.
func (m *MockPeerDas) IsArchivedMode() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsArchivedMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsArchivedMode indicates an expected call of IsArchivedMode.
func (mr *MockPeerDasMockRecorder) IsArchivedMode() *MockPeerDasIsArchivedModeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsArchivedMode", reflect.TypeOf((*MockPeerDas)(nil).IsArchivedMode))
	return &MockPeerDasIsArchivedModeCall{Call: call}
}

// MockPeerDasIsArchivedModeCall wrap *gomock.Call
type MockPeerDasIsArchivedModeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPeerDasIsArchivedModeCall) Return(arg0 bool) *MockPeerDasIsArchivedModeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPeerDasIsArchivedModeCall) Do(f func() bool) *MockPeerDasIsArchivedModeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPeerDasIsArchivedModeCall) DoAndReturn(f func() bool) *MockPeerDasIsArchivedModeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsBlobAlreadyRecovered mocks base method.
func (m *MockPeerDas) IsBlobAlreadyRecovered(blockRoot common.Hash) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBlobAlreadyRecovered", blockRoot)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBlobAlreadyRecovered indicates an expected call of IsBlobAlreadyRecovered.
func (mr *MockPeerDasMockRecorder) IsBlobAlreadyRecovered(blockRoot any) *MockPeerDasIsBlobAlreadyRecoveredCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBlobAlreadyRecovered", reflect.TypeOf((*MockPeerDas)(nil).IsBlobAlreadyRecovered), blockRoot)
	return &MockPeerDasIsBlobAlreadyRecoveredCall{Call: call}
}

// MockPeerDasIsBlobAlreadyRecoveredCall wrap *gomock.Call
type MockPeerDasIsBlobAlreadyRecoveredCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPeerDasIsBlobAlreadyRecoveredCall) Return(arg0 bool) *MockPeerDasIsBlobAlreadyRecoveredCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPeerDasIsBlobAlreadyRecoveredCall) Do(f func(common.Hash) bool) *MockPeerDasIsBlobAlreadyRecoveredCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPeerDasIsBlobAlreadyRecoveredCall) DoAndReturn(f func(common.Hash) bool) *MockPeerDasIsBlobAlreadyRecoveredCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsColumnOverHalf mocks base method.
func (m *MockPeerDas) IsColumnOverHalf(blockRoot common.Hash) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsColumnOverHalf", blockRoot)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsColumnOverHalf indicates an expected call of IsColumnOverHalf.
func (mr *MockPeerDasMockRecorder) IsColumnOverHalf(blockRoot any) *MockPeerDasIsColumnOverHalfCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsColumnOverHalf", reflect.TypeOf((*MockPeerDas)(nil).IsColumnOverHalf), blockRoot)
	return &MockPeerDasIsColumnOverHalfCall{Call: call}
}

// MockPeerDasIsColumnOverHalfCall wrap *gomock.Call
type MockPeerDasIsColumnOverHalfCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPeerDasIsColumnOverHalfCall) Return(arg0 bool) *MockPeerDasIsColumnOverHalfCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPeerDasIsColumnOverHalfCall) Do(f func(common.Hash) bool) *MockPeerDasIsColumnOverHalfCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPeerDasIsColumnOverHalfCall) DoAndReturn(f func(common.Hash) bool) *MockPeerDasIsColumnOverHalfCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsDataAvailable mocks base method.
func (m *MockPeerDas) IsDataAvailable(blockRoot common.Hash) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDataAvailable", blockRoot)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsDataAvailable indicates an expected call of IsDataAvailable.
func (mr *MockPeerDasMockRecorder) IsDataAvailable(blockRoot any) *MockPeerDasIsDataAvailableCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDataAvailable", reflect.TypeOf((*MockPeerDas)(nil).IsDataAvailable), blockRoot)
	return &MockPeerDasIsDataAvailableCall{Call: call}
}

// MockPeerDasIsDataAvailableCall wrap *gomock.Call
type MockPeerDasIsDataAvailableCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPeerDasIsDataAvailableCall) Return(arg0 bool, arg1 error) *MockPeerDasIsDataAvailableCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPeerDasIsDataAvailableCall) Do(f func(common.Hash) (bool, error)) *MockPeerDasIsDataAvailableCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPeerDasIsDataAvailableCall) DoAndReturn(f func(common.Hash) (bool, error)) *MockPeerDasIsDataAvailableCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Prune mocks base method.
func (m *MockPeerDas) Prune(keepSlotDistance uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prune", keepSlotDistance)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prune indicates an expected call of Prune.
func (mr *MockPeerDasMockRecorder) Prune(keepSlotDistance any) *MockPeerDasPruneCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prune", reflect.TypeOf((*MockPeerDas)(nil).Prune), keepSlotDistance)
	return &MockPeerDasPruneCall{Call: call}
}

// MockPeerDasPruneCall wrap *gomock.Call
type MockPeerDasPruneCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPeerDasPruneCall) Return(arg0 error) *MockPeerDasPruneCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPeerDasPruneCall) Do(f func(uint64) error) *MockPeerDasPruneCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPeerDasPruneCall) DoAndReturn(f func(uint64) error) *MockPeerDasPruneCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StateReader mocks base method.
func (m *MockPeerDas) StateReader() peerdasstate.PeerDasStateReader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateReader")
	ret0, _ := ret[0].(peerdasstate.PeerDasStateReader)
	return ret0
}

// StateReader indicates an expected call of StateReader.
func (mr *MockPeerDasMockRecorder) StateReader() *MockPeerDasStateReaderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateReader", reflect.TypeOf((*MockPeerDas)(nil).StateReader))
	return &MockPeerDasStateReaderCall{Call: call}
}

// MockPeerDasStateReaderCall wrap *gomock.Call
type MockPeerDasStateReaderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPeerDasStateReaderCall) Return(arg0 peerdasstate.PeerDasStateReader) *MockPeerDasStateReaderCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPeerDasStateReaderCall) Do(f func() peerdasstate.PeerDasStateReader) *MockPeerDasStateReaderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPeerDasStateReaderCall) DoAndReturn(f func() peerdasstate.PeerDasStateReader) *MockPeerDasStateReaderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// TryScheduleRecover mocks base method.
func (m *MockPeerDas) TryScheduleRecover(slot uint64, blockRoot common.Hash) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryScheduleRecover", slot, blockRoot)
	ret0, _ := ret[0].(error)
	return ret0
}

// TryScheduleRecover indicates an expected call of TryScheduleRecover.
func (mr *MockPeerDasMockRecorder) TryScheduleRecover(slot, blockRoot any) *MockPeerDasTryScheduleRecoverCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryScheduleRecover", reflect.TypeOf((*MockPeerDas)(nil).TryScheduleRecover), slot, blockRoot)
	return &MockPeerDasTryScheduleRecoverCall{Call: call}
}

// MockPeerDasTryScheduleRecoverCall wrap *gomock.Call
type MockPeerDasTryScheduleRecoverCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPeerDasTryScheduleRecoverCall) Return(arg0 error) *MockPeerDasTryScheduleRecoverCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPeerDasTryScheduleRecoverCall) Do(f func(uint64, common.Hash) error) *MockPeerDasTryScheduleRecoverCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPeerDasTryScheduleRecoverCall) DoAndReturn(f func(uint64, common.Hash) error) *MockPeerDasTryScheduleRecoverCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateValidatorsCustody mocks base method.
func (m *MockPeerDas) UpdateValidatorsCustody(cgc uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateValidatorsCustody", cgc)
}

// UpdateValidatorsCustody indicates an expected call of UpdateValidatorsCustody.
func (mr *MockPeerDasMockRecorder) UpdateValidatorsCustody(cgc any) *MockPeerDasUpdateValidatorsCustodyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateValidatorsCustody", reflect.TypeOf((*MockPeerDas)(nil).UpdateValidatorsCustody), cgc)
	return &MockPeerDasUpdateValidatorsCustodyCall{Call: call}
}

// MockPeerDasUpdateValidatorsCustodyCall wrap *gomock.Call
type MockPeerDasUpdateValidatorsCustodyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPeerDasUpdateValidatorsCustodyCall) Return() *MockPeerDasUpdateValidatorsCustodyCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPeerDasUpdateValidatorsCustodyCall) Do(f func(uint64)) *MockPeerDasUpdateValidatorsCustodyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPeerDasUpdateValidatorsCustodyCall) DoAndReturn(f func(uint64)) *MockPeerDasUpdateValidatorsCustodyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
