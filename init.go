package provider

import (
	"fmt"
	"github.com/textthree/cvgokit/filekit"
	"github.com/textthree/cvgokit/syskit"
	"github.com/textthree/provider/clog"
	"github.com/textthree/provider/config"
	"github.com/textthree/provider/core"
	"github.com/textthree/provider/i18n"
	"github.com/textthree/provider/localcache"
	"github.com/textthree/provider/orm"
	"github.com/textthree/provider/redis"
)

var services *core.ServicesContainer
var log clog.Service

func init() {
	services = core.NewContainer()
	services.Bind(&config.ConfigProvider{})
	services.Bind(&clog.ClogProvider{})
	services.Bind(&orm.OrmProvider{})
	services.Bind(&redis.RedisProvider{})
	services.Bind(&i18n.I18nProvider{})
	services.Bind(&localcache.LocalCacheProvider{})

	log = services.NewSingle(clog.Name).(clog.Service)
	env := syskit.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	str := "[provider init] current path:" + filekit.Getwd() + " ENV:" + env
	fmt.Println("\033[37m"+str+"\033[0m", "\n")
}

func Clog() clog.Service {
	return log
}

func Svc() *core.ServicesContainer {
	return services
}
