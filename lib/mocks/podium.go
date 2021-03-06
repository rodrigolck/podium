// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/topfreegames/podium/lib (interfaces: PodiumInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	lib "github.com/topfreegames/podium/lib"
	reflect "reflect"
)

// MockPodiumInterface is a mock of PodiumInterface interface
type MockPodiumInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPodiumInterfaceMockRecorder
}

// MockPodiumInterfaceMockRecorder is the mock recorder for MockPodiumInterface
type MockPodiumInterfaceMockRecorder struct {
	mock *MockPodiumInterface
}

// NewMockPodiumInterface creates a new mock instance
func NewMockPodiumInterface(ctrl *gomock.Controller) *MockPodiumInterface {
	mock := &MockPodiumInterface{ctrl: ctrl}
	mock.recorder = &MockPodiumInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPodiumInterface) EXPECT() *MockPodiumInterfaceMockRecorder {
	return m.recorder
}

// DeleteLeaderboard mocks base method
func (m *MockPodiumInterface) DeleteLeaderboard(arg0 context.Context, arg1 string) (*lib.Response, error) {
	ret := m.ctrl.Call(m, "DeleteLeaderboard", arg0, arg1)
	ret0, _ := ret[0].(*lib.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLeaderboard indicates an expected call of DeleteLeaderboard
func (mr *MockPodiumInterfaceMockRecorder) DeleteLeaderboard(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLeaderboard", reflect.TypeOf((*MockPodiumInterface)(nil).DeleteLeaderboard), arg0, arg1)
}

// GetCount mocks base method
func (m *MockPodiumInterface) GetCount(arg0 context.Context, arg1 string) (int, error) {
	ret := m.ctrl.Call(m, "GetCount", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount
func (mr *MockPodiumInterfaceMockRecorder) GetCount(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockPodiumInterface)(nil).GetCount), arg0, arg1)
}

// GetMember mocks base method
func (m *MockPodiumInterface) GetMember(arg0 context.Context, arg1, arg2 string) (*lib.Member, error) {
	ret := m.ctrl.Call(m, "GetMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(*lib.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMember indicates an expected call of GetMember
func (mr *MockPodiumInterfaceMockRecorder) GetMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMember", reflect.TypeOf((*MockPodiumInterface)(nil).GetMember), arg0, arg1, arg2)
}

// GetMembers mocks base method
func (m *MockPodiumInterface) GetMembers(arg0 context.Context, arg1 string, arg2 []string) (*lib.MemberList, error) {
	ret := m.ctrl.Call(m, "GetMembers", arg0, arg1, arg2)
	ret0, _ := ret[0].(*lib.MemberList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembers indicates an expected call of GetMembers
func (mr *MockPodiumInterfaceMockRecorder) GetMembers(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembers", reflect.TypeOf((*MockPodiumInterface)(nil).GetMembers), arg0, arg1, arg2)
}

// GetMembersAroundMember mocks base method
func (m *MockPodiumInterface) GetMembersAroundMember(arg0 context.Context, arg1, arg2 string, arg3 int) (*lib.MemberList, error) {
	ret := m.ctrl.Call(m, "GetMembersAroundMember", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*lib.MemberList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembersAroundMember indicates an expected call of GetMembersAroundMember
func (mr *MockPodiumInterfaceMockRecorder) GetMembersAroundMember(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembersAroundMember", reflect.TypeOf((*MockPodiumInterface)(nil).GetMembersAroundMember), arg0, arg1, arg2, arg3)
}

// GetTop mocks base method
func (m *MockPodiumInterface) GetTop(arg0 context.Context, arg1 string, arg2, arg3 int) (*lib.MemberList, error) {
	ret := m.ctrl.Call(m, "GetTop", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*lib.MemberList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTop indicates an expected call of GetTop
func (mr *MockPodiumInterfaceMockRecorder) GetTop(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTop", reflect.TypeOf((*MockPodiumInterface)(nil).GetTop), arg0, arg1, arg2, arg3)
}

// GetTopPercent mocks base method
func (m *MockPodiumInterface) GetTopPercent(arg0 context.Context, arg1 string, arg2 int) (*lib.MemberList, error) {
	ret := m.ctrl.Call(m, "GetTopPercent", arg0, arg1, arg2)
	ret0, _ := ret[0].(*lib.MemberList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopPercent indicates an expected call of GetTopPercent
func (mr *MockPodiumInterfaceMockRecorder) GetTopPercent(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopPercent", reflect.TypeOf((*MockPodiumInterface)(nil).GetTopPercent), arg0, arg1, arg2)
}

// Healthcheck mocks base method
func (m *MockPodiumInterface) Healthcheck(arg0 context.Context) (string, error) {
	ret := m.ctrl.Call(m, "Healthcheck", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Healthcheck indicates an expected call of Healthcheck
func (mr *MockPodiumInterfaceMockRecorder) Healthcheck(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthcheck", reflect.TypeOf((*MockPodiumInterface)(nil).Healthcheck), arg0)
}

// IncrementScore mocks base method
func (m *MockPodiumInterface) IncrementScore(arg0 context.Context, arg1, arg2 string, arg3, arg4 int) (*lib.MemberList, error) {
	ret := m.ctrl.Call(m, "IncrementScore", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*lib.MemberList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncrementScore indicates an expected call of IncrementScore
func (mr *MockPodiumInterfaceMockRecorder) IncrementScore(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementScore", reflect.TypeOf((*MockPodiumInterface)(nil).IncrementScore), arg0, arg1, arg2, arg3, arg4)
}

// RemoveMemberFromLeaderboard mocks base method
func (m *MockPodiumInterface) RemoveMemberFromLeaderboard(arg0 context.Context, arg1, arg2 string) (*lib.Response, error) {
	ret := m.ctrl.Call(m, "RemoveMemberFromLeaderboard", arg0, arg1, arg2)
	ret0, _ := ret[0].(*lib.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveMemberFromLeaderboard indicates an expected call of RemoveMemberFromLeaderboard
func (mr *MockPodiumInterfaceMockRecorder) RemoveMemberFromLeaderboard(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMemberFromLeaderboard", reflect.TypeOf((*MockPodiumInterface)(nil).RemoveMemberFromLeaderboard), arg0, arg1, arg2)
}

// UpdateScore mocks base method
func (m *MockPodiumInterface) UpdateScore(arg0 context.Context, arg1, arg2 string, arg3, arg4 int) (*lib.Member, error) {
	ret := m.ctrl.Call(m, "UpdateScore", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*lib.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScore indicates an expected call of UpdateScore
func (mr *MockPodiumInterfaceMockRecorder) UpdateScore(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScore", reflect.TypeOf((*MockPodiumInterface)(nil).UpdateScore), arg0, arg1, arg2, arg3, arg4)
}

// UpdateScores mocks base method
func (m *MockPodiumInterface) UpdateScores(arg0 context.Context, arg1 []string, arg2 string, arg3, arg4 int) (*lib.ScoreList, error) {
	ret := m.ctrl.Call(m, "UpdateScores", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*lib.ScoreList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScores indicates an expected call of UpdateScores
func (mr *MockPodiumInterfaceMockRecorder) UpdateScores(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScores", reflect.TypeOf((*MockPodiumInterface)(nil).UpdateScores), arg0, arg1, arg2, arg3, arg4)
}
