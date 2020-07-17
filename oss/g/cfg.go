package g

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/toolkits_/file"
)

type RpcConfig struct {
	Enabled bool   `json:"enabled"`
	Port    string `json:"port"`
}

type GlobalConfig struct {
	Debug            bool       `json:"debug"`
	OssUri           string     `json:"oss_uri"`
	Endpoint         string     `json:"endpoint"`
	BucketName       string     `json:"bucket_name"`
	ObjectPrefix     string     `json:"object_prefix"`
	OssDirectory     string     `json:"oss_directory"`
	ObjectNamePrefix []string   `json:"objectNamePrefix"`
	Rpc              *RpcConfig `json:"rpc"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file: GetCurrentDirectory is", GetCurrentDirectory())
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
