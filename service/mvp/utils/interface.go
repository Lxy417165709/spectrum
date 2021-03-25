package utils

type Chargeable interface {
	GetChargeableObjectName() string
	GetID() int64
}
