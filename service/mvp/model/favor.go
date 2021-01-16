package model

import (
	"spectrum/common/pb"
	"strconv"
)

const INF = 1E18

func GetFavor(favor *pb.Favor) Favor {
	// todo: 要确认优惠参数是合法的
	if favor == nil {
		return &None{}
	}
	switch favor.FavorType {
	case pb.FavorType_NONE:
		return (&None{}).ParseParameters(favor.Parameters)
	case pb.FavorType_REBATE:
		return (&Rebate{}).ParseParameters(favor.Parameters)
	case pb.FavorType_FULL_REDUCTION:
		return (&FullReduction{}).ParseParameters(favor.Parameters)
	case pb.FavorType_FREE:
		return (&Free{}).ParseParameters(favor.Parameters)
	}
	return &None{}
}

type Favor interface {
	GetExpense(nonFavorExpense float64) float64
	ParseParameters(parameters []string) Favor
	GetPriority() int // 获取优先级，优先级越小，越先计算
}

type None struct{}

func (n *None) GetPriority() int {
	return 0
}

func (n *None) ParseParameters(parameters []string) Favor {
	return &None{}
}

func (n *None) GetExpense(nonFavorExpense float64) float64 {
	return nonFavorExpense
}

type Rebate struct {
	Granularity float64 // 取值为 (0, 1]
}

func (r *Rebate) GetPriority() int {
	return 2
}

func (r *Rebate) ParseParameters(parameters []string) Favor {
	if len(parameters) <= 0 {
		return &Rebate{
			Granularity: 1,
		}
	}
	granularity, _ := strconv.ParseFloat(parameters[0], 64)
	if granularity <= 0 {
		return &Rebate{
			Granularity: 1,
		}
	}
	return &Rebate{
		Granularity: granularity,
	}
}

func (r *Rebate) GetExpense(nonFavorExpense float64) float64 {
	return nonFavorExpense * r.Granularity
}

type FullReduction struct {
	Full      float64
	Reduction float64
}

func (f *FullReduction) GetPriority() int {
	return 1
}

func (f *FullReduction) ParseParameters(parameters []string) Favor {
	if len(parameters) <= 1 {
		return &FullReduction{
			Full:      INF,
			Reduction: 0,
		}
	}
	full, _ := strconv.ParseFloat(parameters[0], 64)
	reduction, _ := strconv.ParseFloat(parameters[1], 64)
	if full <= 0 || reduction <= 0 {
		return &FullReduction{
			Full:      INF,
			Reduction: 0,
		}
	}
	return &FullReduction{
		Full:      full,
		Reduction: reduction,
	}
}

func (f *FullReduction) GetExpense(nonFavorExpense float64) float64 {
	return nonFavorExpense - float64(int(nonFavorExpense/f.Full))*f.Reduction
}

type Free struct{}

func (f *Free) GetPriority() int {
	return 0
}

func (f *Free) ParseParameters(parameters []string) Favor {
	return &Free{}
}

func (f *Free) GetExpense(nonFavorExpense float64) float64 {
	return 0 * nonFavorExpense
}
