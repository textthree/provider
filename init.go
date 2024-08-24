package provider

import (
	"fmt"
	"github.com/textthree/provider/clog"
	"github.com/textthree/provider/config"
	"github.com/textthree/provider/core"
	"github.com/textthree/provider/i18n"
	"github.com/textthree/provider/localcache"
	"github.com/textthree/provider/orm"
	"github.com/textthree/provider/redis"
)

var Services = core.NewContainer()
var log clog.Service

func init() {
	Services.Bind(&config.ConfigProvider{})
	Services.Bind(&clog.ClogProvider{})
	Services.Bind(&orm.OrmProvider{})
	Services.Bind(&redis.RedisProvider{})
	Services.Bind(&i18n.I18nProvider{})
	Services.Bind(&localcache.LocalCacheProvider{})

	log = Services.NewSingle(clog.Name).(clog.Service)
	fmt.Println("\033[37mprovider init\033[0m", "\n")
}

func Clog() clog.Service {
	return log
}
