package keeper

import (
	context "context"
	reflect "reflect"

	address "cosmossdk.io/core/address"
	types "github.com/cosmos/cosmos-sdk/types"
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

// AddressCodec mocks base method.
func (m *MockAccountKeeper) AddressCodec() address.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddressCodec")
	ret0, _ := ret[0].(address.Codec)
	return ret0
}

// AddressCodec indicates an expected call of AddressCodec.
func (mr *MockAccountKeeperMockRecorder) AddressCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressCodec", reflect.TypeOf((*MockAccountKeeper)(nil).AddressCodec))
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx context.Context, addr types.AccAddress) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// NewAccountWithAddress mocks base method.
func (m *MockAccountKeeper) NewAccountWithAddress(ctx context.Context, addr types.AccAddress) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccountWithAddress", ctx, addr)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// NewAccountWithAddress indicates an expected call of NewAccountWithAddress.
func (mr *MockAccountKeeperMockRecorder) NewAccountWithAddress(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccountWithAddress", reflect.TypeOf((*MockAccountKeeper)(nil).NewAccountWithAddress), ctx, addr)
}

// SetAccount mocks base method.
func (m *MockAccountKeeper) SetAccount(ctx context.Context, acc types.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", ctx, acc)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountKeeperMockRecorder) SetAccount(ctx, acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetAccount), ctx, acc)
}
