package controller

import (
	"context"
	"github.com/astaxie/beego/logs"
	"log"
	"spectrum/common/pb"
	"testing"
	"time"
)

func TestMvpServer_AddGood(t *testing.T) {
	initMvpClient()

	if _, err := mvpClient.AddGood(context.Background(), &pb.AddGoodReq{
		GoodName: "椰果",
		Price:    1,
	}); err != nil {
		t.Fatal(err)
		return
	}

	log.Println("Add successfully!")
}

func TestMvpServer_SellGood(t *testing.T) {
	initMvpClient()

	if _, err := mvpClient.SellGood(context.Background(), &pb.SellGoodReq{
		GoodName:  "超神水果茶",
		SellPrice: 18.0,
	}); err != nil {
		t.Fatal(err)
		return
	}

	log.Println("Sell successfully!")
}

func TestMvpServer_AddBilliardDesk(t *testing.T) {
	initMvpClient()

	if _, err := mvpClient.AddBilliardDesk(context.Background(), &pb.AddBilliardDeskReq{
		BilliardDeskName: "一号桌",
	}); err != nil {
		t.Fatal(err)
		return
	}

	log.Println("Add successfully!")
}

func TestMvpServer_BeginPlayBilliard(t *testing.T) {
	initMvpClient()

	if _, err := mvpClient.BeginPlayBilliard(context.Background(), &pb.BeginPlayBilliardReq{
		BilliardDeskName:   "一号桌",
		BeginPlayTimestamp: time.Unix(10001, 0).Unix(),
	}); err != nil {
		t.Fatal(err)
		return
	}

	log.Println("Play successfully!")
}

func TestMvpServer_StopPlayBilliard(t *testing.T) {
	initMvpClient()

	if _, err := mvpClient.StopPlayBilliard(context.Background(), &pb.StopPlayBilliardReq{
		BilliardDeskName:   "一号桌",
		BeginPlayTimestamp: time.Unix(10001, 0).Unix(),
		StopPlayTimestamp:  time.Unix(2000000, 0).Unix(),
	}); err != nil {
		t.Fatal(err)
		return
	}

	log.Println("Stop successfully!")
}

func TestMvpServer_Order(t *testing.T) {
	initMvpClient()

	res, err := mvpClient.Order(context.Background(), &pb.OrderReq{
		Goods: []*pb.Good{
			{
				Name: "波霸奶茶",
				AttachGoods: []*pb.AttachGood{
					{
						Name: "椰果",
					},
					{
						Name: "少冰",
					},
				},
			},
			{
				Name: "波霸奶茶",
				AttachGoods: []*pb.AttachGood{
					{
						Name: "椰果",
					},
				},
			},
			{
				Name: "超神水果茶",
			},
		},
	})
	logs.Info(res, err)
	if err != nil {
		t.Fatal(err)
		return
	}

	log.Println("Order successfully!")
}

func TestMvpServer_GetOrderGoods(t *testing.T) {
	initMvpClient()

	res, err := mvpClient.GetOrderGoods(context.Background(), &pb.GetOrderGoodsReq{
		OrderID: 0,
	})
	logs.Info(res, len(res.Goods[0].AttachGoods),err)
	if err != nil {
		t.Fatal(err)
		return
	}

	log.Println("Order successfully!")
}
