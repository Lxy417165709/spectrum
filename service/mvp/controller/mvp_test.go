package controller

import (
	"fmt"
	"spectrum/common/pb"
	"spectrum/service/mvp/model"
	"testing"
)

// 2021年01月17日01:26:56 Parse
func TestGetSelectSizeInfo(t *testing.T) {
	testSamples := [][]*pb.SizeInfo{
		{
			{
				Size:             "小",
				Price:            1,
				PictureStorePath: "",
				IsSelected:       false,
			},
			{
				Size:             "中",
				Price:            2,
				PictureStorePath: "",
				IsSelected:       false,
			},
			{
				Size:             "大",
				Price:            3,
				PictureStorePath: "",
				IsSelected:       true,
			},
		}, // 大
		{
			{
				Size:             "小",
				Price:            1,
				PictureStorePath: "",
				IsSelected:       false,
			},
			{
				Size:             "中",
				Price:            2,
				PictureStorePath: "",
				IsSelected:       false,
			},
			{
				Size:             "大",
				Price:            3,
				PictureStorePath: "",
				IsSelected:       false,
			},
		}, // nil
		{
			{
				Size:             "小",
				Price:            1,
				PictureStorePath: "",
				IsSelected:       true,
			},
			{
				Size:             "中",
				Price:            2,
				PictureStorePath: "",
				IsSelected:       false,
			},
			{
				Size:             "大",
				Price:            3,
				PictureStorePath: "",
				IsSelected:       false,
			},
		}, // 小
		{
			{
				Size:             "小",
				Price:            1,
				PictureStorePath: "",
				IsSelected:       false,
			},
			{
				Size:             "中",
				Price:            2,
				PictureStorePath: "",
				IsSelected:       true,
			},
			{
				Size:             "大",
				Price:            3,
				PictureStorePath: "",
				IsSelected:       false,
			},
		}, // 中
	}
	for index, sample := range testSamples {
		fmt.Printf("Sample %+v: %+v\n", index, model.GetSelectSizeInfo(sample))
	}
	// Sample 0: size:"大" price:3 isSelected:true
	// Sample 1: <nil>
	// Sample 2: size:"小" price:1 isSelected:true
	// Sample 3: size:"中" price:2 isSelected:true
}
