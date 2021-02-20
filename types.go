package types

type LispType interface {
}

type List struct {
	Val []LispType
	Meta LispType
}

type Symbol struct {
	Val string
}