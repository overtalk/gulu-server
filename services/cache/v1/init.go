package cache

import (
	"github.com/jiachuhuang/concurrentcache"

	"gitlab.com/SausageShoot/admin-server/utils/log"
)

// menCache 内存缓存
type memCache struct {
	cc *concurrentcache.ConcurrentCache
}

func Cache() *memCache {
	cc, err := concurrentcache.NewConcurrentCache(256, 10240)
	if err != nil {
		log.Logger.Fatal("Create Mem Cache", log.ErrorField(err))
	}

	return &memCache{cc: cc}
}
