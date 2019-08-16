package module

type GMSource interface {
	Fetch(fileName string) ([]byte, error)
}

type GameGM interface {
	GetWeaponList() []int64
	GetFashionList() []int64
}
