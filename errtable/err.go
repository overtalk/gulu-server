package errtable

const OkCode int = 1000000

func GenCode(l ErrLevel, m ErrModule, t int) int {
	return int(l)*1000000 + int(m)*10000 + t
}
