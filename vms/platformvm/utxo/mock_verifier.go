// Copyright (C) 2022-2023, Dijets Inc, All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lasthyphen/dijetsnodego/vms/platformvm/utxo (interfaces: Verifier)

// Package utxo is a generated GoMock package.
package utxo

import (
	reflect "reflect"

	ids "github.com/lasthyphen/dijetsnodego/ids"
	djtx "github.com/lasthyphen/dijetsnodego/vms/components/djtx"
	verify "github.com/lasthyphen/dijetsnodego/vms/components/verify"
	state "github.com/lasthyphen/dijetsnodego/vms/platformvm/state"
	txs "github.com/lasthyphen/dijetsnodego/vms/platformvm/txs"
	gomock "github.com/golang/mock/gomock"
)

// MockVerifier is a mock of Verifier interface.
type MockVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockVerifierMockRecorder
}

// MockVerifierMockRecorder is the mock recorder for MockVerifier.
type MockVerifierMockRecorder struct {
	mock *MockVerifier
}

// NewMockVerifier creates a new mock instance.
func NewMockVerifier(ctrl *gomock.Controller) *MockVerifier {
	mock := &MockVerifier{ctrl: ctrl}
	mock.recorder = &MockVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerifier) EXPECT() *MockVerifierMockRecorder {
	return m.recorder
}

// VerifySpend mocks base method.
func (m *MockVerifier) VerifySpend(arg0 txs.UnsignedTx, arg1 state.UTXOGetter, arg2 []*djtx.TransferableInput, arg3 []*djtx.TransferableOutput, arg4 []verify.Verifiable, arg5 map[ids.ID]uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySpend", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySpend indicates an expected call of VerifySpend.
func (mr *MockVerifierMockRecorder) VerifySpend(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySpend", reflect.TypeOf((*MockVerifier)(nil).VerifySpend), arg0, arg1, arg2, arg3, arg4, arg5)
}

// VerifySpendUTXOs mocks base method.
func (m *MockVerifier) VerifySpendUTXOs(arg0 txs.UnsignedTx, arg1 []*djtx.UTXO, arg2 []*djtx.TransferableInput, arg3 []*djtx.TransferableOutput, arg4 []verify.Verifiable, arg5 map[ids.ID]uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySpendUTXOs", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySpendUTXOs indicates an expected call of VerifySpendUTXOs.
func (mr *MockVerifierMockRecorder) VerifySpendUTXOs(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySpendUTXOs", reflect.TypeOf((*MockVerifier)(nil).VerifySpendUTXOs), arg0, arg1, arg2, arg3, arg4, arg5)
}
