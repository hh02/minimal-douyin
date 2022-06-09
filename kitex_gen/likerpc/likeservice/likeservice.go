// Code generated by Kitex v0.3.1. DO NOT EDIT.

package likeservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/hh02/minimal-douyin/kitex_gen/likerpc"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return likeServiceServiceInfo
}

var likeServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "LikeService"
	handlerType := (*likerpc.LikeService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateLike":        kitex.NewMethodInfo(createLikeHandler, newCreateLikeArgs, newCreateLikeResult, false),
		"DeleteLike":        kitex.NewMethodInfo(deleteLikeHandler, newDeleteLikeArgs, newDeleteLikeResult, false),
		"QueryLikeByUserId": kitex.NewMethodInfo(queryLikeByUserIdHandler, newQueryLikeByUserIdArgs, newQueryLikeByUserIdResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "like",
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

func createLikeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(likerpc.CreateLikeRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(likerpc.LikeService).CreateLike(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateLikeArgs:
		success, err := handler.(likerpc.LikeService).CreateLike(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateLikeResult)
		realResult.Success = success
	}
	return nil
}
func newCreateLikeArgs() interface{} {
	return &CreateLikeArgs{}
}

func newCreateLikeResult() interface{} {
	return &CreateLikeResult{}
}

type CreateLikeArgs struct {
	Req *likerpc.CreateLikeRequest
}

func (p *CreateLikeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateLikeArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateLikeArgs) Unmarshal(in []byte) error {
	msg := new(likerpc.CreateLikeRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateLikeArgs_Req_DEFAULT *likerpc.CreateLikeRequest

func (p *CreateLikeArgs) GetReq() *likerpc.CreateLikeRequest {
	if !p.IsSetReq() {
		return CreateLikeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateLikeArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateLikeResult struct {
	Success *likerpc.CreateLikeResponse
}

var CreateLikeResult_Success_DEFAULT *likerpc.CreateLikeResponse

func (p *CreateLikeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateLikeResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateLikeResult) Unmarshal(in []byte) error {
	msg := new(likerpc.CreateLikeResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateLikeResult) GetSuccess() *likerpc.CreateLikeResponse {
	if !p.IsSetSuccess() {
		return CreateLikeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateLikeResult) SetSuccess(x interface{}) {
	p.Success = x.(*likerpc.CreateLikeResponse)
}

func (p *CreateLikeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deleteLikeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(likerpc.DeleteLikeRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(likerpc.LikeService).DeleteLike(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteLikeArgs:
		success, err := handler.(likerpc.LikeService).DeleteLike(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteLikeResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteLikeArgs() interface{} {
	return &DeleteLikeArgs{}
}

func newDeleteLikeResult() interface{} {
	return &DeleteLikeResult{}
}

type DeleteLikeArgs struct {
	Req *likerpc.DeleteLikeRequest
}

func (p *DeleteLikeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteLikeArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteLikeArgs) Unmarshal(in []byte) error {
	msg := new(likerpc.DeleteLikeRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteLikeArgs_Req_DEFAULT *likerpc.DeleteLikeRequest

func (p *DeleteLikeArgs) GetReq() *likerpc.DeleteLikeRequest {
	if !p.IsSetReq() {
		return DeleteLikeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteLikeArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteLikeResult struct {
	Success *likerpc.DeleteLikeResponse
}

var DeleteLikeResult_Success_DEFAULT *likerpc.DeleteLikeResponse

func (p *DeleteLikeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteLikeResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteLikeResult) Unmarshal(in []byte) error {
	msg := new(likerpc.DeleteLikeResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteLikeResult) GetSuccess() *likerpc.DeleteLikeResponse {
	if !p.IsSetSuccess() {
		return DeleteLikeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteLikeResult) SetSuccess(x interface{}) {
	p.Success = x.(*likerpc.DeleteLikeResponse)
}

func (p *DeleteLikeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func queryLikeByUserIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(likerpc.QueryLikeByUserIdRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(likerpc.LikeService).QueryLikeByUserId(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryLikeByUserIdArgs:
		success, err := handler.(likerpc.LikeService).QueryLikeByUserId(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryLikeByUserIdResult)
		realResult.Success = success
	}
	return nil
}
func newQueryLikeByUserIdArgs() interface{} {
	return &QueryLikeByUserIdArgs{}
}

func newQueryLikeByUserIdResult() interface{} {
	return &QueryLikeByUserIdResult{}
}

type QueryLikeByUserIdArgs struct {
	Req *likerpc.QueryLikeByUserIdRequest
}

func (p *QueryLikeByUserIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryLikeByUserIdArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryLikeByUserIdArgs) Unmarshal(in []byte) error {
	msg := new(likerpc.QueryLikeByUserIdRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryLikeByUserIdArgs_Req_DEFAULT *likerpc.QueryLikeByUserIdRequest

func (p *QueryLikeByUserIdArgs) GetReq() *likerpc.QueryLikeByUserIdRequest {
	if !p.IsSetReq() {
		return QueryLikeByUserIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryLikeByUserIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type QueryLikeByUserIdResult struct {
	Success *likerpc.QueryLikeByUserIdResponse
}

var QueryLikeByUserIdResult_Success_DEFAULT *likerpc.QueryLikeByUserIdResponse

func (p *QueryLikeByUserIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryLikeByUserIdResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryLikeByUserIdResult) Unmarshal(in []byte) error {
	msg := new(likerpc.QueryLikeByUserIdResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryLikeByUserIdResult) GetSuccess() *likerpc.QueryLikeByUserIdResponse {
	if !p.IsSetSuccess() {
		return QueryLikeByUserIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryLikeByUserIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*likerpc.QueryLikeByUserIdResponse)
}

func (p *QueryLikeByUserIdResult) IsSetSuccess() bool {
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

func (p *kClient) CreateLike(ctx context.Context, Req *likerpc.CreateLikeRequest) (r *likerpc.CreateLikeResponse, err error) {
	var _args CreateLikeArgs
	_args.Req = Req
	var _result CreateLikeResult
	if err = p.c.Call(ctx, "CreateLike", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteLike(ctx context.Context, Req *likerpc.DeleteLikeRequest) (r *likerpc.DeleteLikeResponse, err error) {
	var _args DeleteLikeArgs
	_args.Req = Req
	var _result DeleteLikeResult
	if err = p.c.Call(ctx, "DeleteLike", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryLikeByUserId(ctx context.Context, Req *likerpc.QueryLikeByUserIdRequest) (r *likerpc.QueryLikeByUserIdResponse, err error) {
	var _args QueryLikeByUserIdArgs
	_args.Req = Req
	var _result QueryLikeByUserIdResult
	if err = p.c.Call(ctx, "QueryLikeByUserId", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
