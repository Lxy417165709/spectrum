package model

type Chargeable interface {
	GetName() string
	GetID() int64
	SetID(id int64)
	SetCheckOutTimestamp(timestamp int64)
	GetCheckOutTimestamp() int64
}
