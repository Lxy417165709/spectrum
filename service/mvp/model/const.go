package model

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

//var ChargeableObjectMap = map[string]Chargeable{
//	reflect.TypeOf(&pb.Good{}).String(): &Good{},
//	reflect.TypeOf(&pb.Desk{}).String(): &Desk{},
//}
