package minegate

import (
	log "github.com/jackyyf/golog"
	"runtime"
)

func Run() {
	PreLoadConfig()
	confInit()
	PostLoadConfig()
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Infof("SRO反向代理 %s 已启动.", version_full)
	ServerSocket()
}
