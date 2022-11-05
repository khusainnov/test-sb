package sbercloud

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/khusainnov/sbercloud/internal/entity"
	"github.com/khusainnov/sbercloud/pb"
	"google.golang.org/protobuf/types/known/anypb"
)

type ConfigResp struct {
	body map[string]*anypb.Any `json:"body"`
}

func (s *sb.Server) UploadConfig(ctx context.Context, req *pb.Config) (*pb.ConfigBodyResponse, error) {
	var d entity.UploadData
	var rsp pb.ConfigBodyResponse

	fmt.Printf("Any body")

	json.Unmarshal(req.GetService(), &d)

	//d.Data = req.GetData()

	fmt.Printf("%+v\n", d)

	rsp.Message = "ready"

	return &rsp, nil
}

func (s *Server) GetConfig(ctx context.Context, reqName *pb.ConfigName) (*pb.ConfigBody, error) {
	//fmt.Printf("%+v\n", reqName.String())
	////var rsp ConfigResp
	//var strv interface{}
	//strv = "Hello body"
	//msg, _ := anypb.New(strv.(protoreflect.ProtoMessage))
	///*rsp = map[string]*anypb.Any{
	//	"key1": msg,
	//}*/
	//
	//data := map[string]*anypb.Any{
	//	"key1": msg,
	//}
	//
	//respMsg := &pb.ConfigBody{Data: data}
	//
	//fmt.Printf("%+v\n", respMsg.Data)

	return nil, nil
}
func (s *Server) UpdateConfig(ctx context.Context, req *pb.Config) (*pb.ConfigBodyResponse, error) {
	return nil, nil
}
func (s *Server) DeleteConfig(ctx context.Context, reqName *pb.ConfigName) (*pb.ConfigBodyResponse, error) {

	return nil, nil
}
