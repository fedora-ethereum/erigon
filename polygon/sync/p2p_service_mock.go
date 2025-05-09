// Code generated by MockGen. DO NOT EDIT.
// Source: ./p2p_service.go
//
// Generated by this command:
//
//	mockgen -typed=true -source=./p2p_service.go -destination=./p2p_service_mock.go -package=sync . p2pservice
//

// Package sync is a generated GoMock package.
package sync

import (
	context "context"
	big "math/big"
	reflect "reflect"

	common "github.com/erigontech/erigon-lib/common"
	types "github.com/erigontech/erigon-lib/types"
	p2p "github.com/erigontech/erigon/polygon/p2p"
	gomock "go.uber.org/mock/gomock"
)

// Mockp2pService is a mock of p2pService interface.
type Mockp2pService struct {
	ctrl     *gomock.Controller
	recorder *Mockp2pServiceMockRecorder
	isgomock struct{}
}

// Mockp2pServiceMockRecorder is the mock recorder for Mockp2pService.
type Mockp2pServiceMockRecorder struct {
	mock *Mockp2pService
}

// NewMockp2pService creates a new mock instance.
func NewMockp2pService(ctrl *gomock.Controller) *Mockp2pService {
	mock := &Mockp2pService{ctrl: ctrl}
	mock.recorder = &Mockp2pServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockp2pService) EXPECT() *Mockp2pServiceMockRecorder {
	return m.recorder
}

// FetchBlocksBackwardsByHash mocks base method.
func (m *Mockp2pService) FetchBlocksBackwardsByHash(ctx context.Context, hash common.Hash, amount uint64, peerId *p2p.PeerId, opts ...p2p.FetcherOption) (p2p.FetcherResponse[[]*types.Block], error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, hash, amount, peerId}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchBlocksBackwardsByHash", varargs...)
	ret0, _ := ret[0].(p2p.FetcherResponse[[]*types.Block])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBlocksBackwardsByHash indicates an expected call of FetchBlocksBackwardsByHash.
func (mr *Mockp2pServiceMockRecorder) FetchBlocksBackwardsByHash(ctx, hash, amount, peerId any, opts ...any) *Mockp2pServiceFetchBlocksBackwardsByHashCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, hash, amount, peerId}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBlocksBackwardsByHash", reflect.TypeOf((*Mockp2pService)(nil).FetchBlocksBackwardsByHash), varargs...)
	return &Mockp2pServiceFetchBlocksBackwardsByHashCall{Call: call}
}

// Mockp2pServiceFetchBlocksBackwardsByHashCall wrap *gomock.Call
type Mockp2pServiceFetchBlocksBackwardsByHashCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *Mockp2pServiceFetchBlocksBackwardsByHashCall) Return(arg0 p2p.FetcherResponse[[]*types.Block], arg1 error) *Mockp2pServiceFetchBlocksBackwardsByHashCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *Mockp2pServiceFetchBlocksBackwardsByHashCall) Do(f func(context.Context, common.Hash, uint64, *p2p.PeerId, ...p2p.FetcherOption) (p2p.FetcherResponse[[]*types.Block], error)) *Mockp2pServiceFetchBlocksBackwardsByHashCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *Mockp2pServiceFetchBlocksBackwardsByHashCall) DoAndReturn(f func(context.Context, common.Hash, uint64, *p2p.PeerId, ...p2p.FetcherOption) (p2p.FetcherResponse[[]*types.Block], error)) *Mockp2pServiceFetchBlocksBackwardsByHashCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FetchBodies mocks base method.
func (m *Mockp2pService) FetchBodies(ctx context.Context, headers []*types.Header, peerId *p2p.PeerId, opts ...p2p.FetcherOption) (p2p.FetcherResponse[[]*types.Body], error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, headers, peerId}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchBodies", varargs...)
	ret0, _ := ret[0].(p2p.FetcherResponse[[]*types.Body])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBodies indicates an expected call of FetchBodies.
func (mr *Mockp2pServiceMockRecorder) FetchBodies(ctx, headers, peerId any, opts ...any) *Mockp2pServiceFetchBodiesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, headers, peerId}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBodies", reflect.TypeOf((*Mockp2pService)(nil).FetchBodies), varargs...)
	return &Mockp2pServiceFetchBodiesCall{Call: call}
}

// Mockp2pServiceFetchBodiesCall wrap *gomock.Call
type Mockp2pServiceFetchBodiesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *Mockp2pServiceFetchBodiesCall) Return(arg0 p2p.FetcherResponse[[]*types.Body], arg1 error) *Mockp2pServiceFetchBodiesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *Mockp2pServiceFetchBodiesCall) Do(f func(context.Context, []*types.Header, *p2p.PeerId, ...p2p.FetcherOption) (p2p.FetcherResponse[[]*types.Body], error)) *Mockp2pServiceFetchBodiesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *Mockp2pServiceFetchBodiesCall) DoAndReturn(f func(context.Context, []*types.Header, *p2p.PeerId, ...p2p.FetcherOption) (p2p.FetcherResponse[[]*types.Body], error)) *Mockp2pServiceFetchBodiesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FetchHeaders mocks base method.
func (m *Mockp2pService) FetchHeaders(ctx context.Context, start, end uint64, peerId *p2p.PeerId, opts ...p2p.FetcherOption) (p2p.FetcherResponse[[]*types.Header], error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, start, end, peerId}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchHeaders", varargs...)
	ret0, _ := ret[0].(p2p.FetcherResponse[[]*types.Header])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchHeaders indicates an expected call of FetchHeaders.
func (mr *Mockp2pServiceMockRecorder) FetchHeaders(ctx, start, end, peerId any, opts ...any) *Mockp2pServiceFetchHeadersCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, start, end, peerId}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchHeaders", reflect.TypeOf((*Mockp2pService)(nil).FetchHeaders), varargs...)
	return &Mockp2pServiceFetchHeadersCall{Call: call}
}

// Mockp2pServiceFetchHeadersCall wrap *gomock.Call
type Mockp2pServiceFetchHeadersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *Mockp2pServiceFetchHeadersCall) Return(arg0 p2p.FetcherResponse[[]*types.Header], arg1 error) *Mockp2pServiceFetchHeadersCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *Mockp2pServiceFetchHeadersCall) Do(f func(context.Context, uint64, uint64, *p2p.PeerId, ...p2p.FetcherOption) (p2p.FetcherResponse[[]*types.Header], error)) *Mockp2pServiceFetchHeadersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *Mockp2pServiceFetchHeadersCall) DoAndReturn(f func(context.Context, uint64, uint64, *p2p.PeerId, ...p2p.FetcherOption) (p2p.FetcherResponse[[]*types.Header], error)) *Mockp2pServiceFetchHeadersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListPeersMayHaveBlockNum mocks base method.
func (m *Mockp2pService) ListPeersMayHaveBlockNum(blockNum uint64) []*p2p.PeerId {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPeersMayHaveBlockNum", blockNum)
	ret0, _ := ret[0].([]*p2p.PeerId)
	return ret0
}

// ListPeersMayHaveBlockNum indicates an expected call of ListPeersMayHaveBlockNum.
func (mr *Mockp2pServiceMockRecorder) ListPeersMayHaveBlockNum(blockNum any) *Mockp2pServiceListPeersMayHaveBlockNumCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPeersMayHaveBlockNum", reflect.TypeOf((*Mockp2pService)(nil).ListPeersMayHaveBlockNum), blockNum)
	return &Mockp2pServiceListPeersMayHaveBlockNumCall{Call: call}
}

// Mockp2pServiceListPeersMayHaveBlockNumCall wrap *gomock.Call
type Mockp2pServiceListPeersMayHaveBlockNumCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *Mockp2pServiceListPeersMayHaveBlockNumCall) Return(arg0 []*p2p.PeerId) *Mockp2pServiceListPeersMayHaveBlockNumCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *Mockp2pServiceListPeersMayHaveBlockNumCall) Do(f func(uint64) []*p2p.PeerId) *Mockp2pServiceListPeersMayHaveBlockNumCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *Mockp2pServiceListPeersMayHaveBlockNumCall) DoAndReturn(f func(uint64) []*p2p.PeerId) *Mockp2pServiceListPeersMayHaveBlockNumCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MaxPeers mocks base method.
func (m *Mockp2pService) MaxPeers() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxPeers")
	ret0, _ := ret[0].(int)
	return ret0
}

// MaxPeers indicates an expected call of MaxPeers.
func (mr *Mockp2pServiceMockRecorder) MaxPeers() *Mockp2pServiceMaxPeersCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxPeers", reflect.TypeOf((*Mockp2pService)(nil).MaxPeers))
	return &Mockp2pServiceMaxPeersCall{Call: call}
}

// Mockp2pServiceMaxPeersCall wrap *gomock.Call
type Mockp2pServiceMaxPeersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *Mockp2pServiceMaxPeersCall) Return(arg0 int) *Mockp2pServiceMaxPeersCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *Mockp2pServiceMaxPeersCall) Do(f func() int) *Mockp2pServiceMaxPeersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *Mockp2pServiceMaxPeersCall) DoAndReturn(f func() int) *Mockp2pServiceMaxPeersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Penalize mocks base method.
func (m *Mockp2pService) Penalize(ctx context.Context, peerId *p2p.PeerId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Penalize", ctx, peerId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Penalize indicates an expected call of Penalize.
func (mr *Mockp2pServiceMockRecorder) Penalize(ctx, peerId any) *Mockp2pServicePenalizeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Penalize", reflect.TypeOf((*Mockp2pService)(nil).Penalize), ctx, peerId)
	return &Mockp2pServicePenalizeCall{Call: call}
}

// Mockp2pServicePenalizeCall wrap *gomock.Call
type Mockp2pServicePenalizeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *Mockp2pServicePenalizeCall) Return(arg0 error) *Mockp2pServicePenalizeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *Mockp2pServicePenalizeCall) Do(f func(context.Context, *p2p.PeerId) error) *Mockp2pServicePenalizeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *Mockp2pServicePenalizeCall) DoAndReturn(f func(context.Context, *p2p.PeerId) error) *Mockp2pServicePenalizeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PublishNewBlock mocks base method.
func (m *Mockp2pService) PublishNewBlock(block *types.Block, td *big.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishNewBlock", block, td)
}

// PublishNewBlock indicates an expected call of PublishNewBlock.
func (mr *Mockp2pServiceMockRecorder) PublishNewBlock(block, td any) *Mockp2pServicePublishNewBlockCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishNewBlock", reflect.TypeOf((*Mockp2pService)(nil).PublishNewBlock), block, td)
	return &Mockp2pServicePublishNewBlockCall{Call: call}
}

// Mockp2pServicePublishNewBlockCall wrap *gomock.Call
type Mockp2pServicePublishNewBlockCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *Mockp2pServicePublishNewBlockCall) Return() *Mockp2pServicePublishNewBlockCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *Mockp2pServicePublishNewBlockCall) Do(f func(*types.Block, *big.Int)) *Mockp2pServicePublishNewBlockCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *Mockp2pServicePublishNewBlockCall) DoAndReturn(f func(*types.Block, *big.Int)) *Mockp2pServicePublishNewBlockCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PublishNewBlockHashes mocks base method.
func (m *Mockp2pService) PublishNewBlockHashes(block *types.Block) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishNewBlockHashes", block)
}

// PublishNewBlockHashes indicates an expected call of PublishNewBlockHashes.
func (mr *Mockp2pServiceMockRecorder) PublishNewBlockHashes(block any) *Mockp2pServicePublishNewBlockHashesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishNewBlockHashes", reflect.TypeOf((*Mockp2pService)(nil).PublishNewBlockHashes), block)
	return &Mockp2pServicePublishNewBlockHashesCall{Call: call}
}

// Mockp2pServicePublishNewBlockHashesCall wrap *gomock.Call
type Mockp2pServicePublishNewBlockHashesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *Mockp2pServicePublishNewBlockHashesCall) Return() *Mockp2pServicePublishNewBlockHashesCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *Mockp2pServicePublishNewBlockHashesCall) Do(f func(*types.Block)) *Mockp2pServicePublishNewBlockHashesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *Mockp2pServicePublishNewBlockHashesCall) DoAndReturn(f func(*types.Block)) *Mockp2pServicePublishNewBlockHashesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Run mocks base method.
func (m *Mockp2pService) Run(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *Mockp2pServiceMockRecorder) Run(ctx any) *Mockp2pServiceRunCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*Mockp2pService)(nil).Run), ctx)
	return &Mockp2pServiceRunCall{Call: call}
}

// Mockp2pServiceRunCall wrap *gomock.Call
type Mockp2pServiceRunCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *Mockp2pServiceRunCall) Return(arg0 error) *Mockp2pServiceRunCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *Mockp2pServiceRunCall) Do(f func(context.Context) error) *Mockp2pServiceRunCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *Mockp2pServiceRunCall) DoAndReturn(f func(context.Context) error) *Mockp2pServiceRunCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
