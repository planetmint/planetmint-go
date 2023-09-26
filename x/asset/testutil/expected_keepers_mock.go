// Code generated by MockGen. DO NOT EDIT.
// Source: x/asset/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	types1 "github.com/planetmint/planetmint-go/x/machine/types"
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx types.Context, addr types.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}

// MockMachineKeeper is a mock of MachineKeeper interface.
type MockMachineKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockMachineKeeperMockRecorder
}

// MockMachineKeeperMockRecorder is the mock recorder for MockMachineKeeper.
type MockMachineKeeperMockRecorder struct {
	mock *MockMachineKeeper
}

// NewMockMachineKeeper creates a new mock instance.
func NewMockMachineKeeper(ctrl *gomock.Controller) *MockMachineKeeper {
	mock := &MockMachineKeeper{ctrl: ctrl}
	mock.recorder = &MockMachineKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineKeeper) EXPECT() *MockMachineKeeperMockRecorder {
	return m.recorder
}

// GetMachine mocks base method.
func (m *MockMachineKeeper) GetMachine(ctx types.Context, index types1.MachineIndex) (types1.Machine, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachine", ctx, index)
	ret0, _ := ret[0].(types1.Machine)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetMachine indicates an expected call of GetMachine.
func (mr *MockMachineKeeperMockRecorder) GetMachine(ctx, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachine", reflect.TypeOf((*MockMachineKeeper)(nil).GetMachine), ctx, index)
}

// GetMachineIndexByPubKey mocks base method.
func (m *MockMachineKeeper) GetMachineIndexByPubKey(ctx types.Context, pubKey string) (types1.MachineIndex, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineIndexByPubKey", ctx, pubKey)
	ret0, _ := ret[0].(types1.MachineIndex)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetMachineIndexByPubKey indicates an expected call of GetMachineIndexByPubKey.
func (mr *MockMachineKeeperMockRecorder) GetMachineIndexByPubKey(ctx, pubKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineIndexByPubKey", reflect.TypeOf((*MockMachineKeeper)(nil).GetMachineIndexByPubKey), ctx, pubKey)
}

// GetMachineIndexByPubKey indicates an expected call of GetMachineIndexByPubKey.
func (mr *MockMachineKeeperMockRecorder) GetMachineIndexByAddress(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineIndexByAddress", reflect.TypeOf((*MockMachineKeeper)(nil).GetMachineIndexByAddress), ctx, address)
}