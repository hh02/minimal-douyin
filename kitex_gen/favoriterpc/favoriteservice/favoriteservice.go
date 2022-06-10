// Code generated by Kitex v0.3.1. DO NOT EDIT.

package favoriteservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/hh02/minimal-douyin/kitex_gen/favoriterpc"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favoriterpc.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateFavorite":        kitex.NewMethodInfo(createFavoriteHandler, newCreateFavoriteArgs, newCreateFavoriteResult, false),
		"DeleteFavorite":        kitex.NewMethodInfo(deleteFavoriteHandler, newDeleteFavoriteArgs, newDeleteFavoriteResult, false),
		"QueryFavoriteByUserId": kitex.NewMethodInfo(queryFavoriteByUserIdHandler, newQueryFavoriteByUserIdArgs, newQueryFavoriteByUserIdResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "favorite",
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

func createFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favoriterpc.CreateFavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favoriterpc.FavoriteService).CreateFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateFavoriteArgs:
		success, err := handler.(favoriterpc.FavoriteService).CreateFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newCreateFavoriteArgs() interface{} {
	return &CreateFavoriteArgs{}
}

func newCreateFavoriteResult() interface{} {
	return &CreateFavoriteResult{}
}

type CreateFavoriteArgs struct {
	Req *favoriterpc.CreateFavoriteRequest
}

func (p *CreateFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(favoriterpc.CreateFavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateFavoriteArgs_Req_DEFAULT *favoriterpc.CreateFavoriteRequest

func (p *CreateFavoriteArgs) GetReq() *favoriterpc.CreateFavoriteRequest {
	if !p.IsSetReq() {
		return CreateFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateFavoriteResult struct {
	Success *favoriterpc.CreateFavoriteResponse
}

var CreateFavoriteResult_Success_DEFAULT *favoriterpc.CreateFavoriteResponse

func (p *CreateFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateFavoriteResult) Unmarshal(in []byte) error {
	msg := new(favoriterpc.CreateFavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateFavoriteResult) GetSuccess() *favoriterpc.CreateFavoriteResponse {
	if !p.IsSetSuccess() {
		return CreateFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*favoriterpc.CreateFavoriteResponse)
}

func (p *CreateFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deleteFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favoriterpc.DeleteFavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favoriterpc.FavoriteService).DeleteFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteFavoriteArgs:
		success, err := handler.(favoriterpc.FavoriteService).DeleteFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteFavoriteArgs() interface{} {
	return &DeleteFavoriteArgs{}
}

func newDeleteFavoriteResult() interface{} {
	return &DeleteFavoriteResult{}
}

type DeleteFavoriteArgs struct {
	Req *favoriterpc.DeleteFavoriteRequest
}

func (p *DeleteFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(favoriterpc.DeleteFavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteFavoriteArgs_Req_DEFAULT *favoriterpc.DeleteFavoriteRequest

func (p *DeleteFavoriteArgs) GetReq() *favoriterpc.DeleteFavoriteRequest {
	if !p.IsSetReq() {
		return DeleteFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteFavoriteResult struct {
	Success *favoriterpc.DeleteFavoriteResponse
}

var DeleteFavoriteResult_Success_DEFAULT *favoriterpc.DeleteFavoriteResponse

func (p *DeleteFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteFavoriteResult) Unmarshal(in []byte) error {
	msg := new(favoriterpc.DeleteFavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteFavoriteResult) GetSuccess() *favoriterpc.DeleteFavoriteResponse {
	if !p.IsSetSuccess() {
		return DeleteFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*favoriterpc.DeleteFavoriteResponse)
}

func (p *DeleteFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func queryFavoriteByUserIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favoriterpc.QueryFavoriteByUserIdRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favoriterpc.FavoriteService).QueryFavoriteByUserId(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryFavoriteByUserIdArgs:
		success, err := handler.(favoriterpc.FavoriteService).QueryFavoriteByUserId(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryFavoriteByUserIdResult)
		realResult.Success = success
	}
	return nil
}
func newQueryFavoriteByUserIdArgs() interface{} {
	return &QueryFavoriteByUserIdArgs{}
}

func newQueryFavoriteByUserIdResult() interface{} {
	return &QueryFavoriteByUserIdResult{}
}

type QueryFavoriteByUserIdArgs struct {
	Req *favoriterpc.QueryFavoriteByUserIdRequest
}

func (p *QueryFavoriteByUserIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryFavoriteByUserIdArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryFavoriteByUserIdArgs) Unmarshal(in []byte) error {
	msg := new(favoriterpc.QueryFavoriteByUserIdRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryFavoriteByUserIdArgs_Req_DEFAULT *favoriterpc.QueryFavoriteByUserIdRequest

func (p *QueryFavoriteByUserIdArgs) GetReq() *favoriterpc.QueryFavoriteByUserIdRequest {
	if !p.IsSetReq() {
		return QueryFavoriteByUserIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryFavoriteByUserIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type QueryFavoriteByUserIdResult struct {
	Success *favoriterpc.QueryFavoriteByUserIdResponse
}

var QueryFavoriteByUserIdResult_Success_DEFAULT *favoriterpc.QueryFavoriteByUserIdResponse

func (p *QueryFavoriteByUserIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryFavoriteByUserIdResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryFavoriteByUserIdResult) Unmarshal(in []byte) error {
	msg := new(favoriterpc.QueryFavoriteByUserIdResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryFavoriteByUserIdResult) GetSuccess() *favoriterpc.QueryFavoriteByUserIdResponse {
	if !p.IsSetSuccess() {
		return QueryFavoriteByUserIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryFavoriteByUserIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*favoriterpc.QueryFavoriteByUserIdResponse)
}

func (p *QueryFavoriteByUserIdResult) IsSetSuccess() bool {
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

func (p *kClient) CreateFavorite(ctx context.Context, Req *favoriterpc.CreateFavoriteRequest) (r *favoriterpc.CreateFavoriteResponse, err error) {
	var _args CreateFavoriteArgs
	_args.Req = Req
	var _result CreateFavoriteResult
	if err = p.c.Call(ctx, "CreateFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteFavorite(ctx context.Context, Req *favoriterpc.DeleteFavoriteRequest) (r *favoriterpc.DeleteFavoriteResponse, err error) {
	var _args DeleteFavoriteArgs
	_args.Req = Req
	var _result DeleteFavoriteResult
	if err = p.c.Call(ctx, "DeleteFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryFavoriteByUserId(ctx context.Context, Req *favoriterpc.QueryFavoriteByUserIdRequest) (r *favoriterpc.QueryFavoriteByUserIdResponse, err error) {
	var _args QueryFavoriteByUserIdArgs
	_args.Req = Req
	var _result QueryFavoriteByUserIdResult
	if err = p.c.Call(ctx, "QueryFavoriteByUserId", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
