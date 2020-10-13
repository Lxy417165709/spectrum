package controller

import (
	"google.golang.org/grpc"
	"log"
	"test/common/pb"
)

var serverClient pb.ServerClient

func init() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}()
	serverClient = pb.NewServerClient(conn)
}
