package model

type Chargeable interface {
	GetName() string
	GetID() int64
}
