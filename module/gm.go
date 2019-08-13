package module

type GMSource interface {
	Fetch(fileName string) ([]byte, error)
}

type GM interface {
}
