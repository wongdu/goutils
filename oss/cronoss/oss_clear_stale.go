package cronoss

import (
	"io/ioutil"
	"os"
	"oss/g"
	"time"

	log "github.com/toolkits_/logrus"
)

func ClearOldOssFiles() {
	spec := "20 * * * * *"
	cronPointer.AddFunc(spec, func() {
		cleanStale()
	})
	go cronPointer.Start()
	defer cronPointer.Stop()
	select {}
}

func cleanStale() {
	log.Println("start cleanStale ...")
	reserveRecent := g.Config().ReserveRecent
	if reserveRecent <= 0 {
		return
	}

	//listRecentDate := make([]string, reserveRecent)
	var listRecentDate []string
	const layout = "20060102"
	t := time.Now()
	for idx := 0; idx < reserveRecent; idx++ {
		strCurrDay := (t.Add(time.Duration(-1*idx*24) * time.Hour)).Format(layout)
		listRecentDate = append(listRecentDate, strCurrDay)
	}

	//clear the old oss files and directories
	fs, _ := ioutil.ReadDir(g.Config().OssDirectory)
	for _, file := range fs {
		if file.IsDir() {
			currDirName := file.Name()
			var bOssDir bool = false
			for _, strDay := range listRecentDate {
				if strDay == currDirName {
					bOssDir = true
					break
				}
			}
			if !bOssDir {
				os.RemoveAll(g.Config().OssDirectory + file.Name())
			}
		} else {
			os.Remove(g.Config().OssDirectory + file.Name())
		}
	}
}
