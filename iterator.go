package data_structures

type Iterator[T any] interface {
	HasNext() bool
	GetNext() *T
}
