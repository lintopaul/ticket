// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lintopaul/ticket/proto (interfaces: TicketServiceClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	proto "github.com/lintopaul/ticket/proto"
	grpc "google.golang.org/grpc"
)

// MockTicketServiceClient is a mock of TicketServiceClient interface.
type MockTicketServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTicketServiceClientMockRecorder
}

// MockTicketServiceClientMockRecorder is the mock recorder for MockTicketServiceClient.
type MockTicketServiceClientMockRecorder struct {
	mock *MockTicketServiceClient
}

// NewMockTicketServiceClient creates a new mock instance.
func NewMockTicketServiceClient(ctrl *gomock.Controller) *MockTicketServiceClient {
	mock := &MockTicketServiceClient{ctrl: ctrl}
	mock.recorder = &MockTicketServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketServiceClient) EXPECT() *MockTicketServiceClientMockRecorder {
	return m.recorder
}

// GetReceipt mocks base method.
func (m *MockTicketServiceClient) GetReceipt(arg0 context.Context, arg1 *proto.ReceiptRequest, arg2 ...grpc.CallOption) (*proto.ReceiptResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReceipt", varargs...)
	ret0, _ := ret[0].(*proto.ReceiptResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceipt indicates an expected call of GetReceipt.
func (mr *MockTicketServiceClientMockRecorder) GetReceipt(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipt", reflect.TypeOf((*MockTicketServiceClient)(nil).GetReceipt), varargs...)
}

// GetUsersBySection mocks base method.
func (m *MockTicketServiceClient) GetUsersBySection(arg0 context.Context, arg1 *proto.SectionRequest, arg2 ...grpc.CallOption) (*proto.SectionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsersBySection", varargs...)
	ret0, _ := ret[0].(*proto.SectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersBySection indicates an expected call of GetUsersBySection.
func (mr *MockTicketServiceClientMockRecorder) GetUsersBySection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersBySection", reflect.TypeOf((*MockTicketServiceClient)(nil).GetUsersBySection), varargs...)
}

// ModifySeat mocks base method.
func (m *MockTicketServiceClient) ModifySeat(arg0 context.Context, arg1 *proto.ModifySeatRequest, arg2 ...grpc.CallOption) (*proto.ModifySeatResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ModifySeat", varargs...)
	ret0, _ := ret[0].(*proto.ModifySeatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifySeat indicates an expected call of ModifySeat.
func (mr *MockTicketServiceClientMockRecorder) ModifySeat(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifySeat", reflect.TypeOf((*MockTicketServiceClient)(nil).ModifySeat), varargs...)
}

// PurchaseTicket mocks base method.
func (m *MockTicketServiceClient) PurchaseTicket(arg0 context.Context, arg1 *proto.PurchaseRequest, arg2 ...grpc.CallOption) (*proto.PurchaseResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PurchaseTicket", varargs...)
	ret0, _ := ret[0].(*proto.PurchaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PurchaseTicket indicates an expected call of PurchaseTicket.
func (mr *MockTicketServiceClientMockRecorder) PurchaseTicket(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurchaseTicket", reflect.TypeOf((*MockTicketServiceClient)(nil).PurchaseTicket), varargs...)
}

// RemoveUser mocks base method.
func (m *MockTicketServiceClient) RemoveUser(arg0 context.Context, arg1 *proto.RemoveUserRequest, arg2 ...grpc.CallOption) (*proto.RemoveUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveUser", varargs...)
	ret0, _ := ret[0].(*proto.RemoveUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockTicketServiceClientMockRecorder) RemoveUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockTicketServiceClient)(nil).RemoveUser), varargs...)
}
