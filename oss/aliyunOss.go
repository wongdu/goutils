package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"oss/g"
	"oss/syncaliyunoss"
	"path/filepath"
	"runtime/debug"
	"strings"
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

func checkOssDirectory(dirName string) bool {
	if dirName == "" {
		return false
	}
	absDir, err := filepath.Abs(dirName)
	_, err = os.Stat(absDir)
	if err != nil {
		log.Printf("the oss directory %s is not exits,just create it", absDir)
		os.Mkdir(absDir, 0777)
	}

	fi, err := os.Stat(absDir)
	if !fi.IsDir() {
		log.Errorf("the %s is not a valid oss directory", dirName)
		return false
	}

	return true
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

	g.Config().OssDirectory = filepath.Clean(g.Config().OssDirectory)
	if !checkOssDirectory(g.Config().OssDirectory) {
		log.Error("the oss directory specified in the config is invalid")
		os.Exit(2)
	}

	if !strings.HasSuffix(g.Config().OssDirectory, "/") {
		g.Config().OssDirectory = g.Config().OssDirectory + "/"
	}

	if g.Config().Rpc.Enabled {
		go syncaliyunoss.StartRpcServer(g.Config().Rpc.Port)
		log.Println("start the rpc server...")
	}

	quit := make(chan bool)
	log.Println("synchronise aliyun oss server running...")
	<-quit
}
