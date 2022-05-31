// Code generated by Kitex v0.3.1. DO NOT EDIT.

package userservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/hh02/minimal-douyin/kitex_gen/userrpc"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*userrpc.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateUser": kitex.NewMethodInfo(createUserHandler, newCreateUserArgs, newCreateUserResult, false),
		"GetUser":    kitex.NewMethodInfo(getUserHandler, newGetUserArgs, newGetUserResult, false),
		"MGetUser":   kitex.NewMethodInfo(mGetUserHandler, newMGetUserArgs, newMGetUserResult, false),
		"CheckUser":  kitex.NewMethodInfo(checkUserHandler, newCheckUserArgs, newCheckUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userrpc.CreateUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userrpc.UserService).CreateUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateUserArgs:
		success, err := handler.(userrpc.UserService).CreateUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateUserResult)
		realResult.Success = success
	}
	return nil
}
func newCreateUserArgs() interface{} {
	return &CreateUserArgs{}
}

func newCreateUserResult() interface{} {
	return &CreateUserResult{}
}

type CreateUserArgs struct {
	Req *userrpc.CreateUserRequest
}

func (p *CreateUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateUserArgs) Unmarshal(in []byte) error {
	msg := new(userrpc.CreateUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateUserArgs_Req_DEFAULT *userrpc.CreateUserRequest

func (p *CreateUserArgs) GetReq() *userrpc.CreateUserRequest {
	if !p.IsSetReq() {
		return CreateUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateUserResult struct {
	Success *userrpc.CreateUserResponse
}

var CreateUserResult_Success_DEFAULT *userrpc.CreateUserResponse

func (p *CreateUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateUserResult) Unmarshal(in []byte) error {
	msg := new(userrpc.CreateUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateUserResult) GetSuccess() *userrpc.CreateUserResponse {
	if !p.IsSetSuccess() {
		return CreateUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userrpc.CreateUserResponse)
}

func (p *CreateUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userrpc.GetUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userrpc.UserService).GetUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserArgs:
		success, err := handler.(userrpc.UserService).GetUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserArgs() interface{} {
	return &GetUserArgs{}
}

func newGetUserResult() interface{} {
	return &GetUserResult{}
}

type GetUserArgs struct {
	Req *userrpc.GetUserRequest
}

func (p *GetUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserArgs) Unmarshal(in []byte) error {
	msg := new(userrpc.GetUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserArgs_Req_DEFAULT *userrpc.GetUserRequest

func (p *GetUserArgs) GetReq() *userrpc.GetUserRequest {
	if !p.IsSetReq() {
		return GetUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserResult struct {
	Success *userrpc.GetUserResponse
}

var GetUserResult_Success_DEFAULT *userrpc.GetUserResponse

func (p *GetUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserResult) Unmarshal(in []byte) error {
	msg := new(userrpc.GetUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserResult) GetSuccess() *userrpc.GetUserResponse {
	if !p.IsSetSuccess() {
		return GetUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userrpc.GetUserResponse)
}

func (p *GetUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func mGetUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userrpc.MGetUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userrpc.UserService).MGetUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MGetUserArgs:
		success, err := handler.(userrpc.UserService).MGetUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MGetUserResult)
		realResult.Success = success
	}
	return nil
}
func newMGetUserArgs() interface{} {
	return &MGetUserArgs{}
}

func newMGetUserResult() interface{} {
	return &MGetUserResult{}
}

type MGetUserArgs struct {
	Req *userrpc.MGetUserRequest
}

func (p *MGetUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MGetUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *MGetUserArgs) Unmarshal(in []byte) error {
	msg := new(userrpc.MGetUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MGetUserArgs_Req_DEFAULT *userrpc.MGetUserRequest

func (p *MGetUserArgs) GetReq() *userrpc.MGetUserRequest {
	if !p.IsSetReq() {
		return MGetUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MGetUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type MGetUserResult struct {
	Success *userrpc.MGetUserResponse
}

var MGetUserResult_Success_DEFAULT *userrpc.MGetUserResponse

func (p *MGetUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MGetUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *MGetUserResult) Unmarshal(in []byte) error {
	msg := new(userrpc.MGetUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MGetUserResult) GetSuccess() *userrpc.MGetUserResponse {
	if !p.IsSetSuccess() {
		return MGetUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MGetUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userrpc.MGetUserResponse)
}

func (p *MGetUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userrpc.CheckUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userrpc.UserService).CheckUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckUserArgs:
		success, err := handler.(userrpc.UserService).CheckUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckUserResult)
		realResult.Success = success
	}
	return nil
}
func newCheckUserArgs() interface{} {
	return &CheckUserArgs{}
}

func newCheckUserResult() interface{} {
	return &CheckUserResult{}
}

type CheckUserArgs struct {
	Req *userrpc.CheckUserRequest
}

func (p *CheckUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CheckUserArgs) Unmarshal(in []byte) error {
	msg := new(userrpc.CheckUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckUserArgs_Req_DEFAULT *userrpc.CheckUserRequest

func (p *CheckUserArgs) GetReq() *userrpc.CheckUserRequest {
	if !p.IsSetReq() {
		return CheckUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckUserResult struct {
	Success *userrpc.CheckUserResponse
}

var CheckUserResult_Success_DEFAULT *userrpc.CheckUserResponse

func (p *CheckUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CheckUserResult) Unmarshal(in []byte) error {
	msg := new(userrpc.CheckUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckUserResult) GetSuccess() *userrpc.CheckUserResponse {
	if !p.IsSetSuccess() {
		return CheckUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userrpc.CheckUserResponse)
}

func (p *CheckUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateUser(ctx context.Context, Req *userrpc.CreateUserRequest) (r *userrpc.CreateUserResponse, err error) {
	var _args CreateUserArgs
	_args.Req = Req
	var _result CreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUser(ctx context.Context, Req *userrpc.GetUserRequest) (r *userrpc.GetUserResponse, err error) {
	var _args GetUserArgs
	_args.Req = Req
	var _result GetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetUser(ctx context.Context, Req *userrpc.MGetUserRequest) (r *userrpc.MGetUserResponse, err error) {
	var _args MGetUserArgs
	_args.Req = Req
	var _result MGetUserResult
	if err = p.c.Call(ctx, "MGetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUser(ctx context.Context, Req *userrpc.CheckUserRequest) (r *userrpc.CheckUserResponse, err error) {
	var _args CheckUserArgs
	_args.Req = Req
	var _result CheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
