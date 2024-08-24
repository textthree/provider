package redis

import (
	"github.com/redis/go-redis/v9" // https://redis.uptrace.dev/zh/guide
	"github.com/spf13/cast"
	"github.com/textthree/provider/config"
	"github.com/textthree/provider/core"
	"github.com/textthree/provider/core/types"
	"sync"
)

type Service interface {
	init()
	GetConnPool(connName ...string) *redis.Client
}

type RedisService struct {
	Service
	c           core.Container
	clients     map[string]*redis.Client
	lock        sync.Mutex
	cfgSvc      config.Service
	redisConfig map[string]types.RedisConfig
}

// 如果能获取到配置文件则进行连接，
// 上锁防止调用端在协程中并发初始化时出错
func (self *RedisService) init() {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.redisConfig = self.cfgSvc.GetRedis()
	for connName, configItem := range self.redisConfig {
		rdb := redis.NewClient(&redis.Options{
			Addr:     configItem.Host + ":" + cast.ToString(configItem.Port),
			Password: configItem.Auth,
			DB:       configItem.Db,
		})
		// 挂到 map 中
		if self.clients == nil {
			self.clients = make(map[string]*redis.Client)
		}
		self.clients[connName] = rdb
	}
}

func (self *RedisService) GetConnPool(connName ...string) *redis.Client {
	if self.clients == nil {
		self.init()
	}
	var key string
	if len(connName) > 0 {
		key = connName[0]
	} else {
		for _, v := range self.redisConfig {
			key = v.DefaultConn
			break
		}
	}
	return self.clients[key]
}
