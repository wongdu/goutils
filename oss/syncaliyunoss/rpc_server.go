package syncaliyunoss

import (
	"oss/g"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/toolkits_/httplib"
	log "github.com/toolkits_/logrus"
	"golang.org/x/net/context"
)

type RpcServer struct{}

type ossInfo struct {
	StatusCode      int64  `json:"StatusCode"`
	SecurityToken   string `json:"SecurityToken"`
	AccessKeyId     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	Expiration      string `json:"Expiration"`
}

func (s *RpcServer) SyncOssFile(ctx context.Context, request *AliyunOssRequest) (response *AliyunOssReply, err error) {
	log.Println("SyncOssFile get the request str is:", request.String())

	if !checkAliyunOssRequest(request) {
		return &AliyunOssReply{Message: "invalid parameters", ErrCode: -1}, nil
	}
	go aliyunOssDownload(request.Endpoint, request.BucketName, request.ObjectNamePrefix, request.FileName, request.Md5SumValue, request.Timestamp)
	return &AliyunOssReply{Message: "", ErrCode: 0}, nil
}

func aliyunOssDownload(endpoint, bucket_name, object_name_prefix, file_name, md5sum_value string, timestamp int64) {
	log.Println("the config uri is:", g.Config().OssUri)
	uri := g.Config().OssUri
	req := httplib.Get(uri).SetTimeout(5*time.Second, 30*time.Second)

	var oss_info ossInfo
	err := req.ToJson(&oss_info)
	if err != nil {
		log.Errorf("curl %s failed: %v", uri, err)
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
	err = bucket.GetObjectToFile(object_name_prefix+file_name, "F:/1081.7z")
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
