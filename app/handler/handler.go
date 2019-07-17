package handler

import(
	"log"
	"context"
	pb "envoy-grpc/protos"
)


type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler)SayHello(ctx context.Context, in *pb.HelloRequest)(*pb.HelloReply, error){
	log.Println("request here ==============>")
	return &pb.HelloReply{
		Message : "hellow," + in.Name,
	}, nil
}