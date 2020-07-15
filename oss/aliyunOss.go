package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"oss/g"
	"oss/syncaliyunoss"
	"runtime/debug"
	"syscall"

	log "github.com/toolkits_/logrus"
)

func registerSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGPIPE)
	go func() {
		for {
			sig := <-sigs
			if sig != syscall.SIGPIPE {
				log.Print(sig, "==============EXIT==============")
				os.Exit(1)
			}
		}
	}()
}

func main() {
	registerSignal()
	defer func() {
		if e := recover(); e != nil {
			log.Print(e)
			debug.PrintStack()
			log.Print("==============EXIT==============")
		}
	}()
	cfg := flag.String("c", "oss/cfg.json", "configuration file")
	//cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()
	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	if g.Config().Debug {
		g.InitLog("debug")
	} else {
		g.InitLog("info")
	}

	if g.Config().Rpc.Enabled {
		go syncaliyunoss.StartRpcServer(g.Config().Rpc.Port)
		log.Println("start the rpc server...")
	}

	quit := make(chan bool)
	log.Println("synchronise aliyun oss server running...")
	<-quit
}
