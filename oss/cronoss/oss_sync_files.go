package cronoss

import (
	"oss/g"
	"oss/utils"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/robfig/cron"
	"github.com/toolkits_/httplib"
	log "github.com/toolkits_/logrus"
)

var (
	cronPointer    *cron.Cron
	ossSyncObjects *SafeMap
	idTryOssFailed cron.EntryID
	lock           *sync.Mutex
	bRetry         bool
)

func init() {
	cronPointer = cron.New(cron.WithSeconds())
	ossSyncObjects = NewSafeMap()
	lock = new(sync.Mutex)
	bRetry = false
}
func StartSyncOssFiles() {
	// spec := *0 0 23 * * *"
	spec := "* * " + strconv.Itoa(g.Config().SyncShartTime) + " * * *"
	cronPointer.AddFunc(spec, func() {
		SyncOssFiles()
	})
	spec = "* * " + strconv.Itoa(g.Config().ClearShartTime) + " * * *"
	cronPointer.AddFunc(spec, func() {
		cleanStale()
	})
	go cronPointer.Start()
	defer cronPointer.Stop()

	select {}
}

func SyncOssFiles() bool {
	lock.Lock()
	defer lock.Unlock()
	log.Println("start SyncOssFiles ...")
	const layout = "20060102"
	t := time.Now()
	strCurrDay := t.Format(layout)
	ossSyncObjects.Clear()
	if bRetry {
		cronPointer.Remove(idTryOssFailed)
		bRetry = false
	}
	if !utils.CheckAndCreate(g.Config().OssDirectory + strCurrDay) {
		log.Errorf("create the %s directory failed", g.Config().OssDirectory+strCurrDay)
		return false
	}
	return syncOssFiles(g.Config().OssDirectory + strCurrDay)
}

func syncOssFiles(currDateDir string) bool {
	var bRet bool = true
	sObjectNamePrefix := g.Config().ObjectNamePrefix
	for _, objNamePrefix := range sObjectNamePrefix {
		if !syncOssWithObjectNamePrefix(currDateDir, objNamePrefix) {
			bRet = false
		}
	}

	return bRet
}

func syncOssWithObjectNamePrefix(currDateDir, objNamePrefix string) bool {
	uri := g.Config().OssUri
	req := httplib.Get(uri).SetTimeout(5*time.Second, 30*time.Second)

	var oss_info g.OssInfo
	err := req.ToJson(&oss_info)
	if err != nil {
		log.Errorf("curl %s failed: %v", uri, err)
		return false
	}

	if !utils.CheckAndCreate(currDateDir + "/" + objNamePrefix) {
		log.Errorf("create the %s directory failed", currDateDir+"/"+objNamePrefix)
		return false
	}

	client, err := oss.New(g.Config().Endpoint, oss_info.AccessKeyId, oss_info.AccessKeySecret, oss.SecurityToken(oss_info.SecurityToken))
	if err != nil {
		log.Errorf("create the oss client failed: %v", err)
		return false
	}

	bucket, err := client.Bucket(g.Config().BucketName)
	if err != nil {
		log.Errorf("get the oss bucket failed: %v", err)
		return false
	}

	if !strings.HasSuffix(objNamePrefix, "/") {
		objNamePrefix = objNamePrefix + "/"
	}
	marker := oss.Marker("")
	prefix := oss.Prefix(objNamePrefix)
	strCurrDay := currDateDir[strings.LastIndex(currDateDir, "/")+1:]
	for {
		lor, err := bucket.ListObjects(marker, prefix)
		if err != nil {
			log.Error("Error:", err)
			return false
		}

		for _, object := range lor.Objects {
			strFileFullDir := object.Key
			strFileDir := strFileFullDir[:strings.LastIndex(strFileFullDir, "/")]

			if !utils.CheckAndCreate(currDateDir + "/" + strFileDir) {
				ossSyncObjects.CheckWrite(strFileFullDir, strCurrDay)
				continue
			}

			bExists, _ := utils.FileExists(currDateDir + "/" + strFileFullDir)
			if bExists {
				continue
			}
			err = bucket.GetObjectToFile(strFileFullDir, currDateDir+"/"+strFileFullDir)
			if err != nil {
				ossSyncObjects.CheckWrite(strFileFullDir, strCurrDay)
			}
		}

		prefix = oss.Prefix(lor.Prefix)
		marker = oss.Marker(lor.NextMarker)
		if !lor.IsTruncated {
			break
		}
	}

	if ossSyncObjects.Count() > 0 && !bRetry {
		spec := "* */" + strconv.Itoa(g.Config().RetryInterval) + " * * * *"
		idTryOssFailed, _ = cronPointer.AddFunc(spec, func() {
			retrySyncOssFailedFile()
		})
		bRetry = true
	}

	return true
}

func retrySyncOssFailedFile() {
	log.Println("start retrySyncOssFailedFile ...")
	if ossSyncObjects.Count() == 0 {
		cronPointer.Remove(idTryOssFailed)
		log.Println("remove the retry crontab task ...")
		return
	}

	uri := g.Config().OssUri
	req := httplib.Get(uri).SetTimeout(5*time.Second, 30*time.Second)

	var oss_info g.OssInfo
	err := req.ToJson(&oss_info)
	if err != nil {
		log.Errorf("curl %s failed when try syncronise: %v", uri, err)
		return
	}

	client, err := oss.New(g.Config().Endpoint, oss_info.AccessKeyId, oss_info.AccessKeySecret, oss.SecurityToken(oss_info.SecurityToken))
	if err != nil {
		log.Errorf("create the oss client failed when try syncronise: %v", err)
		return
	}

	bucket, err := client.Bucket(g.Config().BucketName)
	if err != nil {
		log.Errorf("get the oss bucket failed when try syncronise: %v", err)
		return
	}

	for tempKey, tempValue := range ossSyncObjects.Items() {
		objInfo := tempKey.(string)
		strCurrDay := tempValue.(string)
		bExists, _ := utils.FileExists(g.Config().OssDirectory + strCurrDay + "/" + objInfo)
		if bExists {
			ossSyncObjects.Delete(objInfo)
			continue
		}
		err = bucket.GetObjectToFile(objInfo, g.Config().OssDirectory+strCurrDay+"/"+objInfo)
		if err == nil {
			ossSyncObjects.Delete(objInfo)
		}
	}

}
