package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gsession"
	ossxx "ossxx-ext/app/service"
	"time"
)

func init() {
	ossxx.NewOssXXClient()
	server := g.Server()
	server.SetSessionMaxAge(time.Hour)
	server.SetSessionStorage(gsession.NewStorageMemory())
}
