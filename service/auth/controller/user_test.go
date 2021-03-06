package controller

import (
	"context"
	"log"
	"spectrum/common/pb"
	"testing"
)

func TestServer_Login(t *testing.T) {
	initAuthClient()

	if _, err := authClient.Login(context.Background(), &pb.LoginReq{
		Email:    "957903559@qq.com",
		Password: "123456",
	}); err != nil {
		t.Fatal(err)
		return
	}

	log.Println("Login successfully!")
}

func TestServer_Register(t *testing.T) {
	initAuthClient()

	if _, err := authClient.Register(context.Background(), &pb.RegisterReq{
		Email:    "957903559@qq.com",
		Password: "123456",
	}); err != nil {
		t.Fatal(err)
		return
	}

	log.Println("Register successfully!")
}
