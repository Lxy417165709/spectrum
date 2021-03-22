package model

import "time"

const (
	NotCheckedOut = 1
	HasCheckedOut = 2
)

const (
	ChargeableObjectNameOfGood  = "good"
	ChargeableObjectNameOfOrder = "order"
	ChargeableObjectNameOfDesk  = "desk"
)

const INF = 1E18

var (
	NilTime = time.Unix(10086, 0)
)
