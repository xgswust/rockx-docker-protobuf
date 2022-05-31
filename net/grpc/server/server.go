package server

import (
	"context"
	"github.com/xgswust/rockx-docker-protobuf/net/grpc/config"
	"github.com/xgswust/rockx-docker-protobuf/net/grpc/ip"
	pb "github.com/xgswust/rockx-docker-protobuf/net/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Service 定义我们的服务
type Service struct {
}

func (s *Service) StartNewValidator(ctx context.Context, info *pb.ValidatorInfo) (*pb.MessageResponse, error) {
	//TODO implement me
	return nil, nil
}

func (s *Service) StopValidator(ctx context.Context, info *pb.ValidatorInfo) (*pb.MessageResponse, error) {
	//TODO implement me
	return nil, nil
}

func (s *Service) GetContainerList(ctx context.Context, req *pb.MessageRequest) (*pb.ContainerInfos, error) {
	result := &pb.ContainerInfos{}
	return result, nil
}

// Run ..
func Run() {
	ip, err := ip.GetExternal()
	if err != nil {
		log.Println("getip error")
		return
	}

	log.Println("grpc server external ip", ip)

	ip = "127.0.0.1"
	rpcAddr := ip + ":" + config.Port

	listener, err := net.Listen("tcp", rpcAddr) // 监听本地端口
	if err != nil {
		log.Println(err)
	}
	log.Println("grpc server Listing on", rpcAddr)

	grpcServer := grpc.NewServer() // 新建gRPC服务器实例
	// 在gRPC服务器注册我们的服务
	server := &Service{}
	pb.RegisterGrpcServer(grpcServer, server)

	if err = grpcServer.Serve(listener); err != nil { //用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
		log.Println(err)
	}
}
