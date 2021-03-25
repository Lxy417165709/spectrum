package model

import (
	"spectrum/common/pb"
	"spectrum/service/mvp/utils"
	"time"
)

type QueryOrderParameter struct {
	CheckOutState       pb.CheckOutState
	CreatedTimeInterval TimeInterval
	OrderID             int64
}

func (q *QueryOrderParameter) GetWhereClauseAndValues() (string, []interface{}) {
	var whereClause string
	var whereValues []interface{}
	whereClause = " 1 = 1 "

	// 1. 处理结账状态
	switch q.CheckOutState {
	case pb.CheckOutState_All:
	case pb.CheckOutState_HadCheckOut:
		whereClause += " and check_out_at != ? "
		whereValues = append(whereValues, utils.NilTime)
	case pb.CheckOutState_NotCheckOut:
		whereClause += " and check_out_at = ? "
		whereValues = append(whereValues, utils.NilTime)
	}

	// 2. 处理创建时间区间
	if q.CreatedTimeInterval.Start != utils.NilTime {
		whereClause += " and created_at >= ? "
		whereValues = append(whereValues, q.CreatedTimeInterval.Start)
	}
	if q.CreatedTimeInterval.End != utils.NilTime {
		whereClause += " and created_at <= ? "
		whereValues = append(whereValues, q.CreatedTimeInterval.End)
	}

	// 3. 处理订单号
	if q.OrderID != 0 {
		whereClause += " and id = ? "
		whereValues = append(whereValues, q.OrderID)
	}

	// 4. 返回
	return whereClause, whereValues
}

type TimeInterval struct {
	Start time.Time
	End   time.Time
}
