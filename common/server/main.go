package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
	"test/common/pb"
)

type Server struct {
	pb.UnimplementedServerServer
}

func (s *Server) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	var res pb.LoginRes
	if req.Email == "417165709@qq.com" && req.Password == "123456" {
		return &res, nil
	}
	return &res, errors.New("账号或密码错误")
}

func main() {
	lis, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatal(err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterServerServer(grpcServer, &Server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
