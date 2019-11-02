// +build !windows

package minegate

import (
	log "github.com/jackyyf/golog"
	"os"
	"os/signal"
	"syscall"
	"runtime"
)

func Run() {
	PreLoadConfig()
	confInit()
	PostLoadConfig()
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Infof("SRO反向代理 %s 已启动.", version_full)
	go ServerSocket()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGUSR1)
	for {
		cur := <-sig
		switch cur {
		case syscall.SIGHUP:
			log.Warn("收到指令, 重载配置文件中...")
			PreLoadConfig()
			ConfReload()
			PostLoadConfig()
		case syscall.SIGUSR1:
			log.Warn("收到指令, 正在更新log文件...")
			log.Rotate()
		default:
			log.Errorf("未知指令: %s", cur.String())
			continue
		}
	}
}
