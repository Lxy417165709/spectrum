package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"test/pb"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	serverClient := pb.NewServerClient(conn)
	_, err = serverClient.Login(context.Background(), &pb.LoginReq{
		Email:    "417165709@qq.com",
		Password: "123456",
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Login successfully!")
}
