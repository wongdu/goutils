package syncaliyunoss

import (
	"oss/g"
	"oss/utils"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/toolkits_/httplib"
	log "github.com/toolkits_/logrus"
)

func SyncWhileStart() bool {
	const layout = "20060102"
	t := time.Now()
	strCurrDay := t.Format(layout)
	if !utils.CheckAndCreate(g.Config().OssDirectory + strCurrDay) {
		log.Errorf("create the %s directory failed", g.Config().OssDirectory+strCurrDay)
		return false
	}

	return syncOssFiles(g.Config().OssDirectory + strCurrDay)
}

func syncOssFiles(parentDir string) bool {
	var bRet bool = true
	sObjectNamePrefix := g.Config().ObjectNamePrefix
	for _, objNamePrefix := range sObjectNamePrefix {
		if !syncOssWithObjectNamePrefix(parentDir, objNamePrefix) {
			bRet = false
		}
	}

	return bRet
}

func syncOssWithObjectNamePrefix(parentDir, objNamePrefix string) bool {
	uri := g.Config().OssUri
	req := httplib.Get(uri).SetTimeout(5*time.Second, 30*time.Second)

	var oss_info g.OssInfo
	err := req.ToJson(&oss_info)
	if err != nil {
		log.Errorf("curl %s failed: %v", uri, err)
		return false
	}

	if !utils.CheckAndCreate(parentDir + "/" + objNamePrefix) {
		log.Errorf("create the %s directory failed", parentDir+"/"+objNamePrefix)
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
	for {
		lor, err := bucket.ListObjects(marker, prefix)
		if err != nil {
			log.Error("Error:", err)
			return false
		}

		for _, object := range lor.Objects {
			//fmt.Printf("%s %d %s\n", object.LastModified, object.Size, object.Key)
			strFileFullDir := object.Key
			strFileDir := strFileFullDir[strings.Index(strFileFullDir, "/"):strings.LastIndex(strFileFullDir, "/")]
			if strFileDir != "" {
				if !utils.CheckAndCreate(parentDir + "/" + objNamePrefix + strFileDir) {

					continue
				}
			}
			err = bucket.GetObjectToFile(strFileFullDir, parentDir+"/"+strFileFullDir)
			if err != nil {

			}

		}

		prefix = oss.Prefix(lor.Prefix)
		marker = oss.Marker(lor.NextMarker)
		if !lor.IsTruncated {
			break
		}
	}

	return true
}
