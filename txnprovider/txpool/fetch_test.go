// Copyright 2021 The Erigon Authors
// This file is part of Erigon.
//
// Erigon is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Erigon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Erigon. If not, see <http://www.gnu.org/licenses/>.

package txpool

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/erigontech/erigon-lib/common/u256"
	"github.com/erigontech/erigon-lib/direct"
	"github.com/erigontech/erigon-lib/gointerfaces"
	remote "github.com/erigontech/erigon-lib/gointerfaces/remoteproto"
	"github.com/erigontech/erigon-lib/gointerfaces/sentryproto"
	"github.com/erigontech/erigon-lib/gointerfaces/typesproto"
	"github.com/erigontech/erigon-lib/kv"
	"github.com/erigontech/erigon-lib/kv/memdb"
	"github.com/erigontech/erigon-lib/log/v3"
)

func TestFetch(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	remoteKvClient := remote.NewMockKVClient(ctrl)
	sentryServer := sentryproto.NewMockSentryServer(ctrl)
	pool := NewMockPool(ctrl)
	pool.EXPECT().Started().Return(true)

	m := NewMockSentry(ctx, sentryServer)
	sentryClient := direct.NewSentryClientDirect(direct.ETH67, m)
	var wg sync.WaitGroup
	fetch := NewFetch(ctx, []sentryproto.SentryClient{sentryClient}, pool, remoteKvClient, nil, *u256.N1, log.New(), WithP2PFetcherWg(&wg))
	m.StreamWg.Add(2)
	fetch.ConnectSentries()
	m.StreamWg.Wait()
	// Send one transaction id
	wg.Add(1)
	errs := m.Send(&sentryproto.InboundMessage{
		Id:     sentryproto.MessageId_NEW_POOLED_TRANSACTION_HASHES_66,
		Data:   decodeHex("e1a0595e27a835cd79729ff1eeacec3120eeb6ed1464a04ec727aaca734ead961328"),
		PeerId: peerID,
	})
	for i, err := range errs {
		if err != nil {
			t.Errorf("sending new pool txn hashes 66 (%d): %v", i, err)
		}
	}
	wg.Wait()
}

func TestSendTxnPropagate(t *testing.T) {
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()
	t.Run("few remote byHash", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		sentryServer := sentryproto.NewMockSentryServer(ctrl)

		times := 2
		requests := make([]*sentryproto.SendMessageToRandomPeersRequest, 0, times)
		sentryServer.EXPECT().
			SendMessageToRandomPeers(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, r *sentryproto.SendMessageToRandomPeersRequest) (*sentryproto.SentPeers, error) {
				requests = append(requests, r)
				return nil, nil
			}).
			Times(times)

		sentryServer.EXPECT().PeerById(gomock.Any(), gomock.Any()).
			DoAndReturn(
				func(_ context.Context, r *sentryproto.PeerByIdRequest) (*sentryproto.PeerByIdReply, error) {
					return &sentryproto.PeerByIdReply{
						Peer: &typesproto.PeerInfo{
							Id:   r.PeerId.String(),
							Caps: []string{"eth/68"},
						}}, nil
				}).AnyTimes()

		m := NewMockSentry(ctx, sentryServer)
		send := NewSend(ctx, []sentryproto.SentryClient{direct.NewSentryClientDirect(direct.ETH68, m)}, log.New())
		send.BroadcastPooledTxns(testRlps(2), 100)
		send.AnnouncePooledTxns([]byte{0, 1}, []uint32{10, 15}, toHashes(1, 42), 100)

		require.Len(t, requests, 2)

		txnsMessage := requests[0].Data
		assert.Equal(t, sentryproto.MessageId_TRANSACTIONS_66, txnsMessage.Id)
		assert.Len(t, txnsMessage.Data, 3)

		txnHashesMessage := requests[1].Data
		assert.Equal(t, sentryproto.MessageId_NEW_POOLED_TRANSACTION_HASHES_68, txnHashesMessage.Id)
		assert.Len(t, txnHashesMessage.Data, 76)
	})

	t.Run("much remote byHash", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		sentryServer := sentryproto.NewMockSentryServer(ctrl)

		times := 2
		requests := make([]*sentryproto.SendMessageToRandomPeersRequest, 0, times)
		sentryServer.EXPECT().
			SendMessageToRandomPeers(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, r *sentryproto.SendMessageToRandomPeersRequest) (*sentryproto.SentPeers, error) {
				requests = append(requests, r)
				return nil, nil
			}).
			Times(times)

		m := NewMockSentry(ctx, sentryServer)
		send := NewSend(ctx, []sentryproto.SentryClient{direct.NewSentryClientDirect(direct.ETH68, m)}, log.New())
		list := make(Hashes, p2pTxPacketLimit*3)
		for i := 0; i < len(list); i += 32 {
			b := []byte(fmt.Sprintf("%x", i))
			copy(list[i:i+32], b)
		}
		send.BroadcastPooledTxns(testRlps(len(list)/32), 100)
		send.AnnouncePooledTxns([]byte{0, 1, 2}, []uint32{10, 12, 14}, list, 100)

		require.Len(t, requests, 2)

		txnsMessage := requests[0].Data
		require.Equal(t, sentryproto.MessageId_TRANSACTIONS_66, txnsMessage.Id)
		require.Positive(t, len(txnsMessage.Data))

		txnHashesMessage := requests[1].Data
		require.Equal(t, sentryproto.MessageId_NEW_POOLED_TRANSACTION_HASHES_68, txnHashesMessage.Id)
		require.Positive(t, len(txnHashesMessage.Data))
	})

	t.Run("few local byHash", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		sentryServer := sentryproto.NewMockSentryServer(ctrl)

		times := 2
		requests := make([]*sentryproto.SendMessageToRandomPeersRequest, 0, times)
		sentryServer.EXPECT().
			SendMessageToRandomPeers(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, r *sentryproto.SendMessageToRandomPeersRequest) (*sentryproto.SentPeers, error) {
				requests = append(requests, r)
				return nil, nil
			}).
			Times(times)

		m := NewMockSentry(ctx, sentryServer)
		send := NewSend(ctx, []sentryproto.SentryClient{direct.NewSentryClientDirect(direct.ETH68, m)}, log.New())
		send.BroadcastPooledTxns(testRlps(2), 100)
		send.AnnouncePooledTxns([]byte{0, 1}, []uint32{10, 15}, toHashes(1, 42), 100)

		require.Len(t, requests, 2)

		txnsMessage := requests[0].Data
		assert.Equal(t, sentryproto.MessageId_TRANSACTIONS_66, txnsMessage.Id)
		assert.Positive(t, len(txnsMessage.Data))

		txnHashesMessage := requests[1].Data
		assert.Equal(t, sentryproto.MessageId_NEW_POOLED_TRANSACTION_HASHES_68, txnHashesMessage.Id)
		assert.Len(t, txnHashesMessage.Data, 76)
	})

	t.Run("sync with new peer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		sentryServer := sentryproto.NewMockSentryServer(ctrl)

		times := 3
		requests := make([]*sentryproto.SendMessageByIdRequest, 0, times)
		sentryServer.EXPECT().
			SendMessageById(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, r *sentryproto.SendMessageByIdRequest) (*sentryproto.SentPeers, error) {
				requests = append(requests, r)
				return nil, nil
			}).
			Times(times)

		sentryServer.EXPECT().PeerById(gomock.Any(), gomock.Any()).
			DoAndReturn(
				func(_ context.Context, r *sentryproto.PeerByIdRequest) (*sentryproto.PeerByIdReply, error) {
					return &sentryproto.PeerByIdReply{
						Peer: &typesproto.PeerInfo{
							Id:   r.PeerId.String(),
							Caps: []string{"eth/68"},
						}}, nil
				}).AnyTimes()

		m := NewMockSentry(ctx, sentryServer)
		send := NewSend(ctx, []sentryproto.SentryClient{direct.NewSentryClientDirect(direct.ETH68, m)}, log.New())
		expectPeers := toPeerIDs(1, 2, 42)
		send.PropagatePooledTxnsToPeersList(expectPeers, []byte{0, 1}, []uint32{10, 15}, toHashes(1, 42))

		require.Len(t, requests, 3)
		for i, req := range requests {
			assert.Equal(t, expectPeers[i], PeerID(req.PeerId))
			assert.Equal(t, sentryproto.MessageId_NEW_POOLED_TRANSACTION_HASHES_68, req.Data.Id)
			assert.Positive(t, len(req.Data.Data))
		}
	})
}

func decodeHex(in string) []byte {
	payload, err := hex.DecodeString(in)
	if err != nil {
		panic(err)
	}
	return payload
}

func TestOnNewBlock(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, db := memdb.NewTestDB(t, kv.ChainDB), memdb.NewTestDB(t, kv.TxPoolDB)
	ctrl := gomock.NewController(t)

	stream := remote.NewMockKV_StateChangesClient[*remote.StateChangeBatch](ctrl)
	i := 0
	stream.EXPECT().
		Recv().
		DoAndReturn(func() (*remote.StateChangeBatch, error) {
			if i > 0 {
				return nil, io.EOF
			}
			i++
			return &remote.StateChangeBatch{
				StateVersionId: 1,
				ChangeBatch: []*remote.StateChange{
					{
						Txs: [][]byte{
							decodeHex(TxnParseMainnetTests[0].PayloadStr),
							decodeHex(TxnParseMainnetTests[1].PayloadStr),
							decodeHex(TxnParseMainnetTests[2].PayloadStr),
						},
						BlockHeight: 1,
						BlockHash:   gointerfaces.ConvertHashToH256([32]byte{}),
					},
				},
			}, nil
		}).
		AnyTimes()

	stateChanges := remote.NewMockKVClient(ctrl)
	stateChanges.
		EXPECT().
		StateChanges(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, _ *remote.StateChangeRequest, _ ...grpc.CallOption) (remote.KV_StateChangesClient, error) {
			return stream, nil
		})

	pool := NewMockPool(ctrl)

	pool.EXPECT().
		ValidateSerializedTxn(gomock.Any()).
		DoAndReturn(func(_ []byte) error {
			return nil
		}).
		Times(3)

	var minedTxns TxnSlots
	pool.EXPECT().
		OnNewBlock(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, _ *remote.StateChangeBatch, _ TxnSlots, _ TxnSlots, minedTxnsArg TxnSlots) error {
			minedTxns = minedTxnsArg
			return nil
		}).
		Times(1)

	fetch := NewFetch(ctx, nil, pool, stateChanges, db, *u256.N1, log.New())
	err := fetch.handleStateChanges(ctx, stateChanges)
	assert.ErrorIs(t, io.EOF, err)
	assert.Len(t, minedTxns.Txns, 3)
}

type MockSentry struct {
	ctx context.Context
	*sentryproto.MockSentryServer
	streams      map[sentryproto.MessageId][]sentryproto.Sentry_MessagesServer
	peersStreams []sentryproto.Sentry_PeerEventsServer
	StreamWg     sync.WaitGroup
	lock         sync.RWMutex
}

func NewMockSentry(ctx context.Context, sentryServer *sentryproto.MockSentryServer) *MockSentry {
	return &MockSentry{
		ctx:              ctx,
		MockSentryServer: sentryServer,
	}
}

var peerID PeerID = gointerfaces.ConvertHashToH512([64]byte{0x12, 0x34, 0x50}) // "12345"

func (ms *MockSentry) Send(req *sentryproto.InboundMessage) (errs []error) {
	ms.lock.RLock()
	defer ms.lock.RUnlock()
	for _, stream := range ms.streams[req.Id] {
		if err := stream.Send(req); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (ms *MockSentry) SetStatus(context.Context, *sentryproto.StatusData) (*sentryproto.SetStatusReply, error) {
	return &sentryproto.SetStatusReply{}, nil
}
func (ms *MockSentry) HandShake(context.Context, *emptypb.Empty) (*sentryproto.HandShakeReply, error) {
	return &sentryproto.HandShakeReply{Protocol: sentryproto.Protocol_ETH68}, nil
}
func (ms *MockSentry) Messages(req *sentryproto.MessagesRequest, stream sentryproto.Sentry_MessagesServer) error {
	ms.lock.Lock()
	if ms.streams == nil {
		ms.streams = map[sentryproto.MessageId][]sentryproto.Sentry_MessagesServer{}
	}
	for _, id := range req.Ids {
		ms.streams[id] = append(ms.streams[id], stream)
	}
	ms.lock.Unlock()
	ms.StreamWg.Done()
	select {
	case <-ms.ctx.Done():
		return nil
	case <-stream.Context().Done():
		return nil
	}
}

func (ms *MockSentry) PeerEvents(_ *sentryproto.PeerEventsRequest, stream sentryproto.Sentry_PeerEventsServer) error {
	ms.lock.Lock()
	ms.peersStreams = append(ms.peersStreams, stream)
	ms.lock.Unlock()
	ms.StreamWg.Done()
	select {
	case <-ms.ctx.Done():
		return nil
	case <-stream.Context().Done():
		return nil
	}
}

func testRlps(num int) [][]byte {
	rlps := make([][]byte, num)
	for i := 0; i < num; i++ {
		rlps[i] = []byte{1}
	}
	return rlps
}

func toPeerIDs(h ...byte) (out []PeerID) {
	for i := range h {
		hash := [64]byte{h[i]}
		out = append(out, gointerfaces.ConvertHashToH512(hash))
	}
	return out
}
