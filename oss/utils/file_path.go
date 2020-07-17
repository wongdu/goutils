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
		os.Mkdir(absDir, 0777)
	}

	fi, err := os.Stat(absDir)
	if !fi.IsDir() {
		log.Errorf("the %s is not a valid oss directory", dirName)
		return false
	}

	return true
}
