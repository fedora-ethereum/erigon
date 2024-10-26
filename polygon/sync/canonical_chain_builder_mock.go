// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erigontech/erigon/polygon/sync (interfaces: CanonicalChainBuilder)
//
// Generated by this command:
//
//	mockgen -typed=true -destination=./canonical_chain_builder_mock.go -package=sync . CanonicalChainBuilder
//

// Package sync is a generated GoMock package.
package sync

import (
	context "context"
	reflect "reflect"

	common "github.com/erigontech/erigon-lib/common"
	types "github.com/erigontech/erigon/core/types"
	gomock "go.uber.org/mock/gomock"
)

// MockCanonicalChainBuilder is a mock of CanonicalChainBuilder interface.
type MockCanonicalChainBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockCanonicalChainBuilderMockRecorder
	isgomock struct{}
}

// MockCanonicalChainBuilderMockRecorder is the mock recorder for MockCanonicalChainBuilder.
type MockCanonicalChainBuilderMockRecorder struct {
	mock *MockCanonicalChainBuilder
}

// NewMockCanonicalChainBuilder creates a new mock instance.
func NewMockCanonicalChainBuilder(ctrl *gomock.Controller) *MockCanonicalChainBuilder {
	mock := &MockCanonicalChainBuilder{ctrl: ctrl}
	mock.recorder = &MockCanonicalChainBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCanonicalChainBuilder) EXPECT() *MockCanonicalChainBuilderMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockCanonicalChainBuilder) Connect(ctx context.Context, headers []*types.Header) ([]*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx, headers)
	ret0, _ := ret[0].([]*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connect indicates an expected call of Connect.
func (mr *MockCanonicalChainBuilderMockRecorder) Connect(ctx, headers any) *MockCanonicalChainBuilderConnectCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockCanonicalChainBuilder)(nil).Connect), ctx, headers)
	return &MockCanonicalChainBuilderConnectCall{Call: call}
}

// MockCanonicalChainBuilderConnectCall wrap *gomock.Call
type MockCanonicalChainBuilderConnectCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCanonicalChainBuilderConnectCall) Return(newConnectedHeaders []*types.Header, err error) *MockCanonicalChainBuilderConnectCall {
	c.Call = c.Call.Return(newConnectedHeaders, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCanonicalChainBuilderConnectCall) Do(f func(context.Context, []*types.Header) ([]*types.Header, error)) *MockCanonicalChainBuilderConnectCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCanonicalChainBuilderConnectCall) DoAndReturn(f func(context.Context, []*types.Header) ([]*types.Header, error)) *MockCanonicalChainBuilderConnectCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ContainsHash mocks base method.
func (m *MockCanonicalChainBuilder) ContainsHash(hash common.Hash) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainsHash", hash)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ContainsHash indicates an expected call of ContainsHash.
func (mr *MockCanonicalChainBuilderMockRecorder) ContainsHash(hash any) *MockCanonicalChainBuilderContainsHashCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsHash", reflect.TypeOf((*MockCanonicalChainBuilder)(nil).ContainsHash), hash)
	return &MockCanonicalChainBuilderContainsHashCall{Call: call}
}

// MockCanonicalChainBuilderContainsHashCall wrap *gomock.Call
type MockCanonicalChainBuilderContainsHashCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCanonicalChainBuilderContainsHashCall) Return(arg0 bool) *MockCanonicalChainBuilderContainsHashCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCanonicalChainBuilderContainsHashCall) Do(f func(common.Hash) bool) *MockCanonicalChainBuilderContainsHashCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCanonicalChainBuilderContainsHashCall) DoAndReturn(f func(common.Hash) bool) *MockCanonicalChainBuilderContainsHashCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HeadersInRange mocks base method.
func (m *MockCanonicalChainBuilder) HeadersInRange(start, count uint64) []*types.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadersInRange", start, count)
	ret0, _ := ret[0].([]*types.Header)
	return ret0
}

// HeadersInRange indicates an expected call of HeadersInRange.
func (mr *MockCanonicalChainBuilderMockRecorder) HeadersInRange(start, count any) *MockCanonicalChainBuilderHeadersInRangeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadersInRange", reflect.TypeOf((*MockCanonicalChainBuilder)(nil).HeadersInRange), start, count)
	return &MockCanonicalChainBuilderHeadersInRangeCall{Call: call}
}

// MockCanonicalChainBuilderHeadersInRangeCall wrap *gomock.Call
type MockCanonicalChainBuilderHeadersInRangeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCanonicalChainBuilderHeadersInRangeCall) Return(arg0 []*types.Header) *MockCanonicalChainBuilderHeadersInRangeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCanonicalChainBuilderHeadersInRangeCall) Do(f func(uint64, uint64) []*types.Header) *MockCanonicalChainBuilderHeadersInRangeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCanonicalChainBuilderHeadersInRangeCall) DoAndReturn(f func(uint64, uint64) []*types.Header) *MockCanonicalChainBuilderHeadersInRangeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LowestCommonAncestor mocks base method.
func (m *MockCanonicalChainBuilder) LowestCommonAncestor(a, b common.Hash) (*types.Header, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LowestCommonAncestor", a, b)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// LowestCommonAncestor indicates an expected call of LowestCommonAncestor.
func (mr *MockCanonicalChainBuilderMockRecorder) LowestCommonAncestor(a, b any) *MockCanonicalChainBuilderLowestCommonAncestorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LowestCommonAncestor", reflect.TypeOf((*MockCanonicalChainBuilder)(nil).LowestCommonAncestor), a, b)
	return &MockCanonicalChainBuilderLowestCommonAncestorCall{Call: call}
}

// MockCanonicalChainBuilderLowestCommonAncestorCall wrap *gomock.Call
type MockCanonicalChainBuilderLowestCommonAncestorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCanonicalChainBuilderLowestCommonAncestorCall) Return(arg0 *types.Header, arg1 bool) *MockCanonicalChainBuilderLowestCommonAncestorCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCanonicalChainBuilderLowestCommonAncestorCall) Do(f func(common.Hash, common.Hash) (*types.Header, bool)) *MockCanonicalChainBuilderLowestCommonAncestorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCanonicalChainBuilderLowestCommonAncestorCall) DoAndReturn(f func(common.Hash, common.Hash) (*types.Header, bool)) *MockCanonicalChainBuilderLowestCommonAncestorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PruneNode mocks base method.
func (m *MockCanonicalChainBuilder) PruneNode(hash common.Hash) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PruneNode", hash)
	ret0, _ := ret[0].(error)
	return ret0
}

// PruneNode indicates an expected call of PruneNode.
func (mr *MockCanonicalChainBuilderMockRecorder) PruneNode(hash any) *MockCanonicalChainBuilderPruneNodeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PruneNode", reflect.TypeOf((*MockCanonicalChainBuilder)(nil).PruneNode), hash)
	return &MockCanonicalChainBuilderPruneNodeCall{Call: call}
}

// MockCanonicalChainBuilderPruneNodeCall wrap *gomock.Call
type MockCanonicalChainBuilderPruneNodeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCanonicalChainBuilderPruneNodeCall) Return(arg0 error) *MockCanonicalChainBuilderPruneNodeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCanonicalChainBuilderPruneNodeCall) Do(f func(common.Hash) error) *MockCanonicalChainBuilderPruneNodeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCanonicalChainBuilderPruneNodeCall) DoAndReturn(f func(common.Hash) error) *MockCanonicalChainBuilderPruneNodeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PruneRoot mocks base method.
func (m *MockCanonicalChainBuilder) PruneRoot(newRootNum uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PruneRoot", newRootNum)
	ret0, _ := ret[0].(error)
	return ret0
}

// PruneRoot indicates an expected call of PruneRoot.
func (mr *MockCanonicalChainBuilderMockRecorder) PruneRoot(newRootNum any) *MockCanonicalChainBuilderPruneRootCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PruneRoot", reflect.TypeOf((*MockCanonicalChainBuilder)(nil).PruneRoot), newRootNum)
	return &MockCanonicalChainBuilderPruneRootCall{Call: call}
}

// MockCanonicalChainBuilderPruneRootCall wrap *gomock.Call
type MockCanonicalChainBuilderPruneRootCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCanonicalChainBuilderPruneRootCall) Return(arg0 error) *MockCanonicalChainBuilderPruneRootCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCanonicalChainBuilderPruneRootCall) Do(f func(uint64) error) *MockCanonicalChainBuilderPruneRootCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCanonicalChainBuilderPruneRootCall) DoAndReturn(f func(uint64) error) *MockCanonicalChainBuilderPruneRootCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Reset mocks base method.
func (m *MockCanonicalChainBuilder) Reset(root *types.Header) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", root)
}

// Reset indicates an expected call of Reset.
func (mr *MockCanonicalChainBuilderMockRecorder) Reset(root any) *MockCanonicalChainBuilderResetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockCanonicalChainBuilder)(nil).Reset), root)
	return &MockCanonicalChainBuilderResetCall{Call: call}
}

// MockCanonicalChainBuilderResetCall wrap *gomock.Call
type MockCanonicalChainBuilderResetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCanonicalChainBuilderResetCall) Return() *MockCanonicalChainBuilderResetCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCanonicalChainBuilderResetCall) Do(f func(*types.Header)) *MockCanonicalChainBuilderResetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCanonicalChainBuilderResetCall) DoAndReturn(f func(*types.Header)) *MockCanonicalChainBuilderResetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Root mocks base method.
func (m *MockCanonicalChainBuilder) Root() *types.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Root")
	ret0, _ := ret[0].(*types.Header)
	return ret0
}

// Root indicates an expected call of Root.
func (mr *MockCanonicalChainBuilderMockRecorder) Root() *MockCanonicalChainBuilderRootCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Root", reflect.TypeOf((*MockCanonicalChainBuilder)(nil).Root))
	return &MockCanonicalChainBuilderRootCall{Call: call}
}

// MockCanonicalChainBuilderRootCall wrap *gomock.Call
type MockCanonicalChainBuilderRootCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCanonicalChainBuilderRootCall) Return(arg0 *types.Header) *MockCanonicalChainBuilderRootCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCanonicalChainBuilderRootCall) Do(f func() *types.Header) *MockCanonicalChainBuilderRootCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCanonicalChainBuilderRootCall) DoAndReturn(f func() *types.Header) *MockCanonicalChainBuilderRootCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Tip mocks base method.
func (m *MockCanonicalChainBuilder) Tip() *types.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tip")
	ret0, _ := ret[0].(*types.Header)
	return ret0
}

// Tip indicates an expected call of Tip.
func (mr *MockCanonicalChainBuilderMockRecorder) Tip() *MockCanonicalChainBuilderTipCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tip", reflect.TypeOf((*MockCanonicalChainBuilder)(nil).Tip))
	return &MockCanonicalChainBuilderTipCall{Call: call}
}

// MockCanonicalChainBuilderTipCall wrap *gomock.Call
type MockCanonicalChainBuilderTipCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCanonicalChainBuilderTipCall) Return(arg0 *types.Header) *MockCanonicalChainBuilderTipCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCanonicalChainBuilderTipCall) Do(f func() *types.Header) *MockCanonicalChainBuilderTipCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCanonicalChainBuilderTipCall) DoAndReturn(f func() *types.Header) *MockCanonicalChainBuilderTipCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
