// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bnb-chain/greenfield/x/bridge/types (interfaces: BankKeeper,StakingKeeper,CrossChainKeeper)

// Package types is a generated GoMock package.
package types

import (
	big "math/big"
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockBankKeeper is a mock of BankKeeper interface
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// SendCoinsFromAccountToModule mocks base method
func (m *MockBankKeeper) SendCoinsFromAccountToModule(arg0 types.Context, arg1 types.AccAddress, arg2 string, arg3 types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), arg0, arg1, arg2, arg3)
}

// SendCoinsFromModuleToAccount mocks base method
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(arg0 types.Context, arg1 string, arg2 types.AccAddress, arg3 types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), arg0, arg1, arg2, arg3)
}

// SendCoinsFromModuleToModule mocks base method
func (m *MockBankKeeper) SendCoinsFromModuleToModule(arg0 types.Context, arg1, arg2 string, arg3 types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToModule", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToModule indicates an expected call of SendCoinsFromModuleToModule
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToModule(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToModule), arg0, arg1, arg2, arg3)
}

// SpendableCoins mocks base method
func (m *MockBankKeeper) SpendableCoins(arg0 types.Context, arg1 types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", arg0, arg1)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins
func (mr *MockBankKeeperMockRecorder) SpendableCoins(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), arg0, arg1)
}

// MockStakingKeeper is a mock of StakingKeeper interface
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// BondDenom mocks base method
func (m *MockStakingKeeper) BondDenom(arg0 types.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BondDenom", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// BondDenom indicates an expected call of BondDenom
func (mr *MockStakingKeeperMockRecorder) BondDenom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BondDenom", reflect.TypeOf((*MockStakingKeeper)(nil).BondDenom), arg0)
}

// MockCrossChainKeeper is a mock of CrossChainKeeper interface
type MockCrossChainKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockCrossChainKeeperMockRecorder
}

// MockCrossChainKeeperMockRecorder is the mock recorder for MockCrossChainKeeper
type MockCrossChainKeeperMockRecorder struct {
	mock *MockCrossChainKeeper
}

// NewMockCrossChainKeeper creates a new mock instance
func NewMockCrossChainKeeper(ctrl *gomock.Controller) *MockCrossChainKeeper {
	mock := &MockCrossChainKeeper{ctrl: ctrl}
	mock.recorder = &MockCrossChainKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCrossChainKeeper) EXPECT() *MockCrossChainKeeperMockRecorder {
	return m.recorder
}

// CreateRawIBCPackageWithFee mocks base method
func (m *MockCrossChainKeeper) CreateRawIBCPackageWithFee(arg0 types.Context, arg1 types.ChannelID, arg2 types.CrossChainPackageType, arg3 []byte, arg4, arg5 *big.Int) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRawIBCPackageWithFee", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRawIBCPackageWithFee indicates an expected call of CreateRawIBCPackageWithFee
func (mr *MockCrossChainKeeperMockRecorder) CreateRawIBCPackageWithFee(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRawIBCPackageWithFee", reflect.TypeOf((*MockCrossChainKeeper)(nil).CreateRawIBCPackageWithFee), arg0, arg1, arg2, arg3, arg4, arg5)
}

// RegisterChannel mocks base method
func (m *MockCrossChainKeeper) RegisterChannel(arg0 string, arg1 types.ChannelID, arg2 types.CrossChainApplication) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterChannel", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterChannel indicates an expected call of RegisterChannel
func (mr *MockCrossChainKeeperMockRecorder) RegisterChannel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterChannel", reflect.TypeOf((*MockCrossChainKeeper)(nil).RegisterChannel), arg0, arg1, arg2)
}
