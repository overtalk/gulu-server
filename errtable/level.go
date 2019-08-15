package errtable

type ErrLevel int

// 错误级别
// 默认1为没有错误
const (
	Business ErrLevel = 2
	System   ErrLevel = 3

	BusinessString string = "Business"
	SystemString   string = "System"
)
