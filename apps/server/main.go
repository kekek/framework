package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"

	"wps.ktkt.com/"
)

var (
	version string = "1.0.0"
	appName string = "user.event"

	showVersion = flag.Bool("version", false, "显示版本信息")
)

func main() {

	flag.Parse()

	if *showVersion {
		printVersion()

		os.Exit(0)
	}


	ctx, cancel := context.WithCancel(context.Background())


	// 开启http服务
	go server.StartHttpServer(cfg.Web, version)

	InitSignal(cancel)
}

// InitSignal register signals handler.
func InitSignal(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			cancel()
			return
		case syscall.SIGHUP:

			glog.Info("signal reload config")
			//err := configs.Init()
			//if err != nil {
			//	glog.Error("reload config failed.")
			//}

		default:
			return
		}
	}
}

func printVersion() {
	fmt.Printf("build info :  %s - %s \n", appName, version)
}
