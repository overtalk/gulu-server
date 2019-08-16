package errtable_test

import (
	"testing"

	"gitlab.com/SausageShoot/admin-server/errtable"
)

func TestErrCode(t *testing.T) {
	t.Log(errtable.GenCode(errtable.System, errtable.GM, 1))
}

func BenchmarkGenCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		errtable.GenCode(errtable.Business, errtable.GM, 1)
	}
}
