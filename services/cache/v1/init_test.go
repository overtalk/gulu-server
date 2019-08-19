package cache_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jiachuhuang/concurrentcache"
)

func TestAAA(t *testing.T) {
	cc, err := concurrentcache.NewConcurrentCache(256, 10240)
	if err != nil {
		panic(err)
	}
	ok, err := cc.Set("foo", "bar", 2*time.Second)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("set fail")
	}
	v, err := cc.Get("foo")
	if v != nil {
		fmt.Println(v.(string))
	}

	time.Sleep(3 * time.Second)
	v, err = cc.Get("foo")
	if v != nil {
		fmt.Println(v.(string))
	} else {
		fmt.Println("没有了")
	}

}
