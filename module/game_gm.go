package module

// 游戏后台的配置

type GMSource interface {
	Fetch(fileName string) ([]byte, error)
}

type GameGM interface {
	GetWeaponList() []int64
	GetFashionList() []int64
}
