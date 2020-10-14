package controller

import (
	"context"
	"log"
	"test/common/pb"
	"testing"
)

func TestServer_Login(t *testing.T) {
	initServerClient()

	if _, err := serverClient.Login(context.Background(), &pb.LoginReq{
		Email:    "417165709@qq.com",
		Password: "123456",
	}); err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Login successfully!")
}

func TestServer_Register(t *testing.T) {
	initServerClient()

	if _, err := serverClient.Register(context.Background(), &pb.RegisterReq{
		Email:    "417165709@qq.com",
		Password: "123456",
	}); err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Register successfully!")
}
