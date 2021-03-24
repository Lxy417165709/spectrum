package model

import (
	"spectrum/common/ers"
	"strconv"
)

// -------------------------------------- Favor 接口 --------------------------------------
type Favor interface {
	GetExpense(nonFavorExpense float64) float64
	ParseParameters(parameters []string) (Favor, error)
	GetPriority() int // 获取优先级，优先级越小，越先计算
}

// -------------------------------------- None 类 --------------------------------------
type None struct{}

func (n *None) GetPriority() int {
	return 0
}

func (n *None) ParseParameters(parameters []string) (Favor, error) {
	return &None{}, nil
}

func (n *None) GetExpense(nonFavorExpense float64) float64 {
	return nonFavorExpense
}

// -------------------------------------- Rebate 类 --------------------------------------
type Rebate struct {
	Granularity float64 // 取值为 (0, 1]
}

func (r *Rebate) GetPriority() int {
	return 2
}

func (r *Rebate) ParseParameters(parameters []string) (Favor, error) {
	if len(parameters) == 0 {
		return nil, ers.New("折扣参数数目 不能为 0。")
	}
	granularity, err := strconv.ParseFloat(parameters[0], 64)
	if err != nil {
		return nil, ers.New("折扣粒度 不是合法小数。")
	}
	if granularity <= 0 {
		return nil, ers.New("折扣粒度 不能小于 0, 只能在 (0, 1] 区间内。")
	}
	if granularity > 1 {
		return nil, ers.New("折扣粒度 不能大于 1, 只能在 (0, 1] 区间内。")
	}
	return &Rebate{
		Granularity: granularity,
	}, nil
}

func (r *Rebate) GetExpense(nonFavorExpense float64) float64 {
	return nonFavorExpense * r.Granularity
}

// -------------------------------------- FullReduction 类 --------------------------------------
type FullReduction struct {
	Full      float64
	Reduction float64
}

func (f *FullReduction) GetPriority() int {
	return 1
}

func (f *FullReduction) ParseParameters(parameters []string) (Favor, error) {
	if len(parameters) <= 1 {
		return nil, ers.New("满减参数数目 不能小于 2。")
	}
	full, err := strconv.ParseFloat(parameters[0], 64)
	if err != nil {
		return nil, ers.New("满减参数 满花费 不是合法小数。")
	}
	if full <= 0 {
		return nil, ers.New("满减参数 满花费 必须大于 0。")
	}

	reduction, err := strconv.ParseFloat(parameters[1], 64)
	if err != nil {
		return nil, ers.New("满减参数 减花费 不是合法小数。")
	}
	if reduction <= 0 {
		return nil, ers.New("满减参数 减花费 必须大于 0。")
	}
	return &FullReduction{
		Full:      full,
		Reduction: reduction,
	}, nil
}

func (f *FullReduction) GetExpense(nonFavorExpense float64) float64 {
	return nonFavorExpense - float64(int(nonFavorExpense/f.Full))*f.Reduction
}

// -------------------------------------- Free 类 --------------------------------------
type Free struct{}

func (f *Free) GetPriority() int {
	return 0
}

func (f *Free) ParseParameters(parameters []string) (Favor, error) {
	return &Free{}, nil
}

func (f *Free) GetExpense(nonFavorExpense float64) float64 {
	return 0 * nonFavorExpense
}
