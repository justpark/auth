// Automatically generated by MockGen. DO NOT EDIT!
// Source: ./vendor/github.com/justpark/auth/keto/sdk/go/keto/sdk_warden.go

package proxy

import (
	gomock "github.com/golang/mock/gomock"
	swagger "github.com/justpark/auth/keto/sdk/go/keto/swagger"
)

// Mock of WardenSDK interface
type MockWardenSDK struct {
	ctrl     *gomock.Controller
	recorder *_MockWardenSDKRecorder
}

// Recorder for MockWardenSDK (not exported)
type _MockWardenSDKRecorder struct {
	mock *MockWardenSDK
}

func NewMockWardenSDK(ctrl *gomock.Controller) *MockWardenSDK {
	mock := &MockWardenSDK{ctrl: ctrl}
	mock.recorder = &_MockWardenSDKRecorder{mock}
	return mock
}

func (_m *MockWardenSDK) EXPECT() *_MockWardenSDKRecorder {
	return _m.recorder
}

func (_m *MockWardenSDK) IsSubjectAuthorized(body swagger.WardenSubjectAuthorizationRequest) (*swagger.WardenSubjectAuthorizationResponse, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "IsSubjectAuthorized", body)
	ret0, _ := ret[0].(*swagger.WardenSubjectAuthorizationResponse)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockWardenSDKRecorder) IsSubjectAuthorized(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsSubjectAuthorized", arg0)
}

func (_m *MockWardenSDK) IsOAuth2AccessTokenAuthorized(body swagger.WardenOAuth2AccessTokenAuthorizationRequest) (*swagger.WardenOAuth2AccessTokenAuthorizationResponse, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "IsOAuth2AccessTokenAuthorized", body)
	ret0, _ := ret[0].(*swagger.WardenOAuth2AccessTokenAuthorizationResponse)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockWardenSDKRecorder) IsOAuth2AccessTokenAuthorized(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsOAuth2AccessTokenAuthorized", arg0)
}

func (_m *MockWardenSDK) IsOAuth2ClientAuthorized(body swagger.WardenOAuth2ClientAuthorizationRequest) (*swagger.WardenOAuth2ClientAuthorizationResponse, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "IsOAuth2ClientAuthorized", body)
	ret0, _ := ret[0].(*swagger.WardenOAuth2ClientAuthorizationResponse)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockWardenSDKRecorder) IsOAuth2ClientAuthorized(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsOAuth2ClientAuthorized", arg0)
}
