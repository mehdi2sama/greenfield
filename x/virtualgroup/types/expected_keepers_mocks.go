// Code generated by MockGen. DO NOT EDIT.
// Source: x/virtualgroup/types/expected_keepers.go

// Package types is a generated GoMock package.
package types

import (
	reflect "reflect"

	math "cosmossdk.io/math"
	types "github.com/bnb-chain/greenfield/x/sp/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/x/auth/types"
	gomock "github.com/golang/mock/gomock"
)

// MockSpKeeper is a mock of SpKeeper interface.
type MockSpKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockSpKeeperMockRecorder
}

// MockSpKeeperMockRecorder is the mock recorder for MockSpKeeper.
type MockSpKeeperMockRecorder struct {
	mock *MockSpKeeper
}

// NewMockSpKeeper creates a new mock instance.
func NewMockSpKeeper(ctrl *gomock.Controller) *MockSpKeeper {
	mock := &MockSpKeeper{ctrl: ctrl}
	mock.recorder = &MockSpKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpKeeper) EXPECT() *MockSpKeeperMockRecorder {
	return m.recorder
}

// DepositDenomForSP mocks base method.
func (m *MockSpKeeper) DepositDenomForSP(ctx types0.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DepositDenomForSP", ctx)
	ret0, _ := ret[0].(string)
	return ret0
}

// DepositDenomForSP indicates an expected call of DepositDenomForSP.
func (mr *MockSpKeeperMockRecorder) DepositDenomForSP(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepositDenomForSP", reflect.TypeOf((*MockSpKeeper)(nil).DepositDenomForSP), ctx)
}

// Exit mocks base method.
func (m *MockSpKeeper) Exit(ctx types0.Context, sp *types.StorageProvider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exit", ctx, sp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exit indicates an expected call of Exit.
func (mr *MockSpKeeperMockRecorder) Exit(ctx, sp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exit", reflect.TypeOf((*MockSpKeeper)(nil).Exit), ctx, sp)
}

// GetStorageProvider mocks base method.
func (m *MockSpKeeper) GetStorageProvider(ctx types0.Context, id uint32) (*types.StorageProvider, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageProvider", ctx, id)
	ret0, _ := ret[0].(*types.StorageProvider)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStorageProvider indicates an expected call of GetStorageProvider.
func (mr *MockSpKeeperMockRecorder) GetStorageProvider(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageProvider", reflect.TypeOf((*MockSpKeeper)(nil).GetStorageProvider), ctx, id)
}

// GetStorageProviderByFundingAddr mocks base method.
func (m *MockSpKeeper) GetStorageProviderByFundingAddr(ctx types0.Context, sealAddr types0.AccAddress) (*types.StorageProvider, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageProviderByFundingAddr", ctx, sealAddr)
	ret0, _ := ret[0].(*types.StorageProvider)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStorageProviderByFundingAddr indicates an expected call of GetStorageProviderByFundingAddr.
func (mr *MockSpKeeperMockRecorder) GetStorageProviderByFundingAddr(ctx, sealAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageProviderByFundingAddr", reflect.TypeOf((*MockSpKeeper)(nil).GetStorageProviderByFundingAddr), ctx, sealAddr)
}

// GetStorageProviderByOperatorAddr mocks base method.
func (m *MockSpKeeper) GetStorageProviderByOperatorAddr(ctx types0.Context, addr types0.AccAddress) (*types.StorageProvider, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageProviderByOperatorAddr", ctx, addr)
	ret0, _ := ret[0].(*types.StorageProvider)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetStorageProviderByOperatorAddr indicates an expected call of GetStorageProviderByOperatorAddr.
func (mr *MockSpKeeperMockRecorder) GetStorageProviderByOperatorAddr(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageProviderByOperatorAddr", reflect.TypeOf((*MockSpKeeper)(nil).GetStorageProviderByOperatorAddr), ctx, addr)
}

// SetStorageProvider mocks base method.
func (m *MockSpKeeper) SetStorageProvider(ctx types0.Context, sp *types.StorageProvider) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStorageProvider", ctx, sp)
}

// SetStorageProvider indicates an expected call of SetStorageProvider.
func (mr *MockSpKeeperMockRecorder) SetStorageProvider(ctx, sp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStorageProvider", reflect.TypeOf((*MockSpKeeper)(nil).SetStorageProvider), ctx, sp)
}

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
func (m *MockAccountKeeper) GetAccount(ctx types0.Context, addr types0.AccAddress) types1.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types1.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// GetModuleAccount mocks base method.
func (m *MockAccountKeeper) GetModuleAccount(ctx types0.Context, moduleName string) types1.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, moduleName)
	ret0, _ := ret[0].(types1.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccount(ctx, moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccount), ctx, moduleName)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(name string) types0.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", name)
	ret0, _ := ret[0].(types0.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), name)
}

// IterateAccounts mocks base method.
func (m *MockAccountKeeper) IterateAccounts(ctx types0.Context, process func(types1.AccountI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateAccounts", ctx, process)
}

// IterateAccounts indicates an expected call of IterateAccounts.
func (mr *MockAccountKeeperMockRecorder) IterateAccounts(ctx, process interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateAccounts", reflect.TypeOf((*MockAccountKeeper)(nil).IterateAccounts), ctx, process)
}

// SetModuleAccount mocks base method.
func (m *MockAccountKeeper) SetModuleAccount(arg0 types0.Context, arg1 types1.ModuleAccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetModuleAccount", arg0, arg1)
}

// SetModuleAccount indicates an expected call of SetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) SetModuleAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetModuleAccount), arg0, arg1)
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

// GetAllBalances mocks base method.
func (m *MockBankKeeper) GetAllBalances(ctx types0.Context, addr types0.AccAddress) types0.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
	ret0, _ := ret[0].(types0.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankKeeperMockRecorder) GetAllBalances(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), ctx, addr)
}

// GetBalance mocks base method.
func (m *MockBankKeeper) GetBalance(ctx types0.Context, addr types0.AccAddress, denom string) types0.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, addr, denom)
	ret0, _ := ret[0].(types0.Coin)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockBankKeeperMockRecorder) GetBalance(ctx, addr, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBankKeeper)(nil).GetBalance), ctx, addr, denom)
}

// LockedCoins mocks base method.
func (m *MockBankKeeper) LockedCoins(ctx types0.Context, addr types0.AccAddress) types0.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockedCoins", ctx, addr)
	ret0, _ := ret[0].(types0.Coins)
	return ret0
}

// LockedCoins indicates an expected call of LockedCoins.
func (mr *MockBankKeeperMockRecorder) LockedCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockedCoins", reflect.TypeOf((*MockBankKeeper)(nil).LockedCoins), ctx, addr)
}

// SendCoins mocks base method.
func (m *MockBankKeeper) SendCoins(ctx types0.Context, fromAddr, toAddr types0.AccAddress, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", ctx, fromAddr, toAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockBankKeeperMockRecorder) SendCoins(ctx, fromAddr, toAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockBankKeeper)(nil).SendCoins), ctx, fromAddr, toAddr, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx types0.Context, senderAddr types0.AccAddress, recipientModule string, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx types0.Context, senderModule string, recipientAddr types0.AccAddress, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx types0.Context, addr types0.AccAddress) types0.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types0.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}

// MockPaymentKeeper is a mock of PaymentKeeper interface.
type MockPaymentKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentKeeperMockRecorder
}

// MockPaymentKeeperMockRecorder is the mock recorder for MockPaymentKeeper.
type MockPaymentKeeperMockRecorder struct {
	mock *MockPaymentKeeper
}

// NewMockPaymentKeeper creates a new mock instance.
func NewMockPaymentKeeper(ctrl *gomock.Controller) *MockPaymentKeeper {
	mock := &MockPaymentKeeper{ctrl: ctrl}
	mock.recorder = &MockPaymentKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentKeeper) EXPECT() *MockPaymentKeeperMockRecorder {
	return m.recorder
}

// IsEmptyNetFlow mocks base method.
func (m *MockPaymentKeeper) IsEmptyNetFlow(ctx types0.Context, account types0.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmptyNetFlow", ctx, account)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmptyNetFlow indicates an expected call of IsEmptyNetFlow.
func (mr *MockPaymentKeeperMockRecorder) IsEmptyNetFlow(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmptyNetFlow", reflect.TypeOf((*MockPaymentKeeper)(nil).IsEmptyNetFlow), ctx, account)
}

// QueryDynamicBalance mocks base method.
func (m *MockPaymentKeeper) QueryDynamicBalance(ctx types0.Context, addr types0.AccAddress) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryDynamicBalance", ctx, addr)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryDynamicBalance indicates an expected call of QueryDynamicBalance.
func (mr *MockPaymentKeeperMockRecorder) QueryDynamicBalance(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryDynamicBalance", reflect.TypeOf((*MockPaymentKeeper)(nil).QueryDynamicBalance), ctx, addr)
}

// Withdraw mocks base method.
func (m *MockPaymentKeeper) Withdraw(ctx types0.Context, fromAddr, toAddr types0.AccAddress, amount math.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", ctx, fromAddr, toAddr, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockPaymentKeeperMockRecorder) Withdraw(ctx, fromAddr, toAddr, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockPaymentKeeper)(nil).Withdraw), ctx, fromAddr, toAddr, amount)
}
