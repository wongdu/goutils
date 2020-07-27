package utils

import (
	"os"
	"path/filepath"

	log "github.com/toolkits_/logrus"
)

func CheckAndCreate(dirName string) bool {
	if dirName == "" {
		return false
	}
	absDir, err := filepath.Abs(dirName)
	_, err = os.Stat(absDir)
	if err != nil {
		log.Printf("the oss directory %s is not exits,just create it", absDir)
		os.MkdirAll(absDir, 0777)
	}

	fi, err := os.Stat(absDir)
	if !fi.IsDir() {
		log.Errorf("the %s is not a valid oss directory", dirName)
		return false
	}

	return true
}

func FileExists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil //文件或者文件夹存在
	}
	if os.IsNotExist(err) {
		return false, nil //不存在
	}
	return false, err //不存在，这里的err可以查到具体的错误信息
}

//判断目录是否存在
func IsDir(dir string) bool {
	info, err := os.Stat(dir)
	if err == nil {
		return false
	}
	return info.IsDir()
}

//判断文件是否存在
func IsFile(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
