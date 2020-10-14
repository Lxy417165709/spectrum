package controller

import (
	"context"
	"log"
	"test/common/pb"
	"testing"
)

func TestMvpServer_AddGood(t *testing.T) {
	initMvpClient()

	if _, err := mvpClient.AddGood(context.Background(), &pb.AddGoodReq{
		GoodName: "超神水果茶",
		Price:    18.0,
	}); err != nil {
		t.Fatal(err)
		return
	}

	log.Println("Add successfully!")
}
