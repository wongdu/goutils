package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"oss/cronoss"
	"oss/g"
	"oss/syncaliyunoss"
	"oss/utils"
	"path/filepath"
	"runtime/debug"
	"strings"
	"syscall"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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
	//cfg := flag.String("c", "cfg.json", "configuration file") //for debug
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

	log.Println("OSS Go SDK Version: ", oss.Version)
	g.Config().OssDirectory = filepath.Clean(g.Config().OssDirectory)
	if !utils.CheckAndCreate(g.Config().OssDirectory) {
		log.Error("the oss directory specified in the config is invalid")
		os.Exit(2)
	}

	if !strings.HasSuffix(g.Config().OssDirectory, "/") {
		g.Config().OssDirectory = g.Config().OssDirectory + "/"
	}

	go cronoss.SyncOssFilesCron()
	go cronoss.SyncOssFiles()

	if g.Config().Rpc.Enabled {
		go syncaliyunoss.StartRpcServer(g.Config().Rpc.Port)
		log.Println("start the rpc server...")
	}

	quit := make(chan bool)
	log.Println("synchronise aliyun oss server running...")
	<-quit
}
