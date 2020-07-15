package main

import (
	"flag"
	"fmt"
	"os"
	"oss/g"
	"oss/rpc"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/toolkits_/httplib"
	log "github.com/toolkits_/logrus"
)

type ossInfo struct {
	StatusCode      int64  `json:"StatusCode"`
	SecurityToken   string `json:"SecurityToken"`
	AccessKeyId     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	Expiration      string `json:"Expiration"`
}

func main() {
	cfg := flag.String("c", "oss/cfg.json", "configuration file")
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

	fmt.Println("the config uri is:", g.Config().OssUri)
	uri := g.Config().OssUri
	req := httplib.Get(uri).SetTimeout(5*time.Second, 30*time.Second)
	fmt.Println(req)

	var oss_info ossInfo
	err := req.ToJson(&oss_info)
	if err != nil {
		log.Errorf("curl %s fail: %v", uri, err)
		return
	}
	fmt.Println(oss_info)

	Endpoint := "oss-cn-shenzhen.aliyuncs.com"
	BucketName := "res-leimans-com-1"
	keyPrefix := "user_game_data/"

	// 创建OSSClient实例。
	client, err := oss.New(Endpoint, oss_info.AccessKeyId, oss_info.AccessKeySecret, oss.SecurityToken(oss_info.SecurityToken))
	if err != nil {
		fmt.Println("Error 1:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	//bucket, err := client.Bucket("<yourBucketName>")
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		fmt.Println("Error 2:", err)
		os.Exit(-1)
	}

	// 下载文件到本地文件。
	//err = bucket.GetObjectToFile("<yourObjectName>", "LocalFile")
	err = bucket.GetObjectToFile(keyPrefix+"10833_10111.7z", "F:/1081.7z")
	//err = bucket.GetObjectToFile(keyPrefix+"10833_10111_20200710_110541.7z", "F:/1083.7z")
	if err != nil {
		fmt.Println("Error 3:", err)
		os.Exit(-1)
	}

	go rpc.StartRpc()

	select {}
}
