package common

type IModel interface {
	TableName() string
	GetID() int64
}
