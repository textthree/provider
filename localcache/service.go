package localcache

import (
	"github.com/textthree/provider/core"
)

type Service interface {
	FreeCache(string) *FreeCacheHolder
	BigCache(string) *BigCacheHolder
	FastCache(string) *FastCacheHolder
}

type CacheService struct {
	Service
	holder core.Container
}
