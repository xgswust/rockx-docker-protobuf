package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"rockx-docker/net/grpc/config"
	pb "rockx-docker/net/grpc/proto"
)

// Client ..
type Client struct {
	DialAddr string
}

// NewClient ..
func NewClient(ip string) *Client {
	log.Println("connect remote ip is", ip)
	ip = "127.0.0.1"
	rpcAddr := ip + ":" + config.Port

	return &Client{
		DialAddr: rpcAddr,
	}
}

func (c *Client) connect() (pb.GrpcClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(c.DialAddr, grpc.WithInsecure()) //连接gRPC服务器
	if err != nil {
		return nil, nil, err
	}
	client := pb.NewGrpcClient(conn) //建立客户端
	return client, conn, nil
}

func (c *Client) GetContainers(ctx context.Context) (int, error) {
	client, conn, err := c.connect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	req := new(pb.MessageRequest)

	req.Question = "Get"
	resp, err := client.GetContainerList(ctx, req) //调用方法
	if err != nil {
		return 0, err
	}
	return len(resp.Info), nil
}
