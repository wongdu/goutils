package syncaliyunoss

import (
	"os"
	"oss/cronoss"
	"oss/g"
	"regexp"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/toolkits_/httplib"
	log "github.com/toolkits_/logrus"
	"golang.org/x/net/context"
)

type RpcServer struct{}

func (s *RpcServer) SyncOssFile(ctx context.Context, request *AliyunOssRequest) (response *AliyunOssReply, err error) {
	log.Println("SyncOssFile get the request str is:", request.String())

	if !checkAliyunOssRequest(request) {
		return &AliyunOssReply{Message: "invalid parameters", ErrCode: -1}, nil
	}
	go aliyunOssDownload(request.Endpoint, request.BucketName, request.ObjectNamePrefix, request.FileName, request.Md5SumValue, request.Timestamp)
	return &AliyunOssReply{Message: "", ErrCode: 0}, nil
}

func (s *RpcServer) SyncOssNow(ctx context.Context, request *empty.Empty) (response *AliyunOssReply, err error) {
	go cronoss.SyncOssFiles()
	return &AliyunOssReply{Message: "receive the sync oss file now notification!", ErrCode: 0}, nil
}

func aliyunOssDownload(endpoint, bucket_name, object_name_prefix, file_name, md5sum_value string, timestamp int64) {
	log.Println("the config uri is:", g.Config().OssUri)
	uri := g.Config().OssUri
	req := httplib.Get(uri).SetTimeout(5*time.Second, 30*time.Second)

	var oss_info g.OssInfo
	err := req.ToJson(&oss_info)
	if err != nil {
		log.Errorf("curl %s failed: %v", uri, err)
		return
	}

	strTimestamp := unixTsFormat(timestamp)
	strOssFileDate := ossFileDate(strTimestamp)
	if !checkOssFileDate(strOssFileDate) {
		log.Errorf("the oss file date %s is invalid", strOssFileDate)
		return
	}

	if !checkOssDateDir(strOssFileDate) {
		log.Errorf("the oss file date directory %s is invalid", g.Config().OssDirectory+strOssFileDate)
		return
	}

	// 创建OSSClient实例。
	client, err := oss.New(endpoint, oss_info.AccessKeyId, oss_info.AccessKeySecret, oss.SecurityToken(oss_info.SecurityToken))
	if err != nil {
		log.Errorf("create the oss client failed: %v", err)
		return
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucket_name)
	if err != nil {
		log.Errorf("get the oss bucket failed: %v", err)
		return
	}

	// 下载文件到本地文件。
	err = bucket.GetObjectToFile(object_name_prefix+file_name, g.Config().OssDirectory+strOssFileDate+"/"+object_name_prefix+file_name)
	if err != nil {
		log.Errorf("get the oss file failed: %v", err)
		return
	}

}

func checkAliyunOssRequest(request *AliyunOssRequest) bool {
	log.Println("checkAliyunOssRequest get the request str is:", request.String())
	log.Printf(" get the request str is %v:", request)
	if request.Endpoint == "" ||
		request.BucketName == "" ||
		request.ObjectNamePrefix == "" ||
		request.FileName == "" ||
		request.Md5SumValue == "" ||
		request.Timestamp == 0 {
		return false
	}

	return true
}

func unixTsFormat(ts int64) string {
	return time.Unix(ts, 0).Format("20060102_150405")
}

func ossFileDate(strTimestamp string) string {
	if strTimestamp == "" {
		return ""
	}
	underIdx := strings.Index(strTimestamp, "_")
	if underIdx == -1 {
		return ""
	}

	return strTimestamp[0:underIdx]
}

func checkOssFileDate(strOssFileDate string) bool {
	match, err := regexp.Match(`^[1-9]\d{3}(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])$`, []byte(strOssFileDate))
	if err != nil {
		return false
	}

	return match
}

func checkOssDateDir(strOssFileDate string) bool {
	absDir := g.Config().OssDirectory + strOssFileDate
	_, err := os.Stat(absDir)
	if err != nil {
		log.Printf("the oss file date directory %s is not exits,just create it", absDir)
		os.Mkdir(absDir, 0777)
	}
	fi, err := os.Stat(absDir)
	if !fi.IsDir() {
		log.Errorf("the %s is not a valid oss file date directory", absDir)
		return false
	}
	return true
}
