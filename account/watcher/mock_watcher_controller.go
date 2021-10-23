// Code generated by MockGen. DO NOT EDIT.
// Source: ../account/watcher/watcher_controller.go

// Package watcher is a generated GoMock package.
package watcher

import (
	context "context"
	reflect "reflect"

	btcec "github.com/btcsuite/btcd/btcec"
	chainhash "github.com/btcsuite/btcd/chaincfg/chainhash"
	wire "github.com/btcsuite/btcd/wire"
	gomock "github.com/golang/mock/gomock"
	chainntnfs "github.com/lightningnetwork/lnd/chainntnfs"
)

// MockChainNotifierClient is a mock of ChainNotifierClient interface.
type MockChainNotifierClient struct {
	ctrl     *gomock.Controller
	recorder *MockChainNotifierClientMockRecorder
}

// MockChainNotifierClientMockRecorder is the mock recorder for MockChainNotifierClient.
type MockChainNotifierClientMockRecorder struct {
	mock *MockChainNotifierClient
}

// NewMockChainNotifierClient creates a new mock instance.
func NewMockChainNotifierClient(ctrl *gomock.Controller) *MockChainNotifierClient {
	mock := &MockChainNotifierClient{ctrl: ctrl}
	mock.recorder = &MockChainNotifierClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainNotifierClient) EXPECT() *MockChainNotifierClientMockRecorder {
	return m.recorder
}

// RegisterBlockEpochNtfn mocks base method.
func (m *MockChainNotifierClient) RegisterBlockEpochNtfn(ctx context.Context) (chan int32, chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterBlockEpochNtfn", ctx)
	ret0, _ := ret[0].(chan int32)
	ret1, _ := ret[1].(chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RegisterBlockEpochNtfn indicates an expected call of RegisterBlockEpochNtfn.
func (mr *MockChainNotifierClientMockRecorder) RegisterBlockEpochNtfn(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterBlockEpochNtfn", reflect.TypeOf((*MockChainNotifierClient)(nil).RegisterBlockEpochNtfn), ctx)
}

// RegisterConfirmationsNtfn mocks base method.
func (m *MockChainNotifierClient) RegisterConfirmationsNtfn(ctx context.Context, txid *chainhash.Hash, pkScript []byte, numConfs, heightHint int32) (chan *chainntnfs.TxConfirmation, chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterConfirmationsNtfn", ctx, txid, pkScript, numConfs, heightHint)
	ret0, _ := ret[0].(chan *chainntnfs.TxConfirmation)
	ret1, _ := ret[1].(chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RegisterConfirmationsNtfn indicates an expected call of RegisterConfirmationsNtfn.
func (mr *MockChainNotifierClientMockRecorder) RegisterConfirmationsNtfn(ctx, txid, pkScript, numConfs, heightHint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterConfirmationsNtfn", reflect.TypeOf((*MockChainNotifierClient)(nil).RegisterConfirmationsNtfn), ctx, txid, pkScript, numConfs, heightHint)
}

// RegisterSpendNtfn mocks base method.
func (m *MockChainNotifierClient) RegisterSpendNtfn(ctx context.Context, outpoint *wire.OutPoint, pkScript []byte, heightHint int32) (chan *chainntnfs.SpendDetail, chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterSpendNtfn", ctx, outpoint, pkScript, heightHint)
	ret0, _ := ret[0].(chan *chainntnfs.SpendDetail)
	ret1, _ := ret[1].(chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RegisterSpendNtfn indicates an expected call of RegisterSpendNtfn.
func (mr *MockChainNotifierClientMockRecorder) RegisterSpendNtfn(ctx, outpoint, pkScript, heightHint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterSpendNtfn", reflect.TypeOf((*MockChainNotifierClient)(nil).RegisterSpendNtfn), ctx, outpoint, pkScript, heightHint)
}

// MockWatcherController is a mock of WatcherController interface.
type MockWatcherController struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherControllerMockRecorder
}

// MockWatcherControllerMockRecorder is the mock recorder for MockWatcherController.
type MockWatcherControllerMockRecorder struct {
	mock *MockWatcherController
}

// NewMockWatcherController creates a new mock instance.
func NewMockWatcherController(ctrl *gomock.Controller) *MockWatcherController {
	mock := &MockWatcherController{ctrl: ctrl}
	mock.recorder = &MockWatcherControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatcherController) EXPECT() *MockWatcherControllerMockRecorder {
	return m.recorder
}

// CancelAccountConf mocks base method.
func (m *MockWatcherController) CancelAccountConf(traderKey *btcec.PublicKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelAccountConf", traderKey)
}

// CancelAccountConf indicates an expected call of CancelAccountConf.
func (mr *MockWatcherControllerMockRecorder) CancelAccountConf(traderKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelAccountConf", reflect.TypeOf((*MockWatcherController)(nil).CancelAccountConf), traderKey)
}

// CancelAccountSpend mocks base method.
func (m *MockWatcherController) CancelAccountSpend(traderKey *btcec.PublicKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelAccountSpend", traderKey)
}

// CancelAccountSpend indicates an expected call of CancelAccountSpend.
func (mr *MockWatcherControllerMockRecorder) CancelAccountSpend(traderKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelAccountSpend", reflect.TypeOf((*MockWatcherController)(nil).CancelAccountSpend), traderKey)
}

// Start mocks base method.
func (m *MockWatcherController) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockWatcherControllerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockWatcherController)(nil).Start))
}

// Stop mocks base method.
func (m *MockWatcherController) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockWatcherControllerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockWatcherController)(nil).Stop))
}

// WatchAccountConf mocks base method.
func (m *MockWatcherController) WatchAccountConf(traderKey *btcec.PublicKey, txHash chainhash.Hash, script []byte, numConfs, heightHint uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchAccountConf", traderKey, txHash, script, numConfs, heightHint)
	ret0, _ := ret[0].(error)
	return ret0
}

// WatchAccountConf indicates an expected call of WatchAccountConf.
func (mr *MockWatcherControllerMockRecorder) WatchAccountConf(traderKey, txHash, script, numConfs, heightHint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAccountConf", reflect.TypeOf((*MockWatcherController)(nil).WatchAccountConf), traderKey, txHash, script, numConfs, heightHint)
}

// WatchAccountExpiration mocks base method.
func (m *MockWatcherController) WatchAccountExpiration(traderKey *btcec.PublicKey, expiry uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchAccountExpiration", traderKey, expiry)
	ret0, _ := ret[0].(error)
	return ret0
}

// WatchAccountExpiration indicates an expected call of WatchAccountExpiration.
func (mr *MockWatcherControllerMockRecorder) WatchAccountExpiration(traderKey, expiry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAccountExpiration", reflect.TypeOf((*MockWatcherController)(nil).WatchAccountExpiration), traderKey, expiry)
}

// WatchAccountSpend mocks base method.
func (m *MockWatcherController) WatchAccountSpend(traderKey *btcec.PublicKey, accountPoint wire.OutPoint, script []byte, heightHint uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchAccountSpend", traderKey, accountPoint, script, heightHint)
	ret0, _ := ret[0].(error)
	return ret0
}

// WatchAccountSpend indicates an expected call of WatchAccountSpend.
func (mr *MockWatcherControllerMockRecorder) WatchAccountSpend(traderKey, accountPoint, script, heightHint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAccountSpend", reflect.TypeOf((*MockWatcherController)(nil).WatchAccountSpend), traderKey, accountPoint, script, heightHint)
}