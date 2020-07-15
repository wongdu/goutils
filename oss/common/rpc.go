package common

import (
	"fmt"
)

// code == 0 => success
// code == 1 => bad request
type SimpleRpcResponse struct {
	Code int `json:"code"`
}

func (this *SimpleRpcResponse) String() string {
	return fmt.Sprintf("<Code: %d>", this.Code)
}

type NullRpcRequest struct {
}

type AliyunOssRequest struct {
	Endpoint         string `json:"endpoint"`
	BucketName       string `json:"bucket_name"`
	ObjectNamePrefix string `json:"object_name_prefix"`
	FileName         string `json:"file_name"`
	Md5sumValue      string `json:"md5sum_value"`
	Timestamp        int64  `json:"timestamp"`
}

func (this *AliyunOssRequest) String() string {
	return fmt.Sprintf(
		"<Endpoint:%s, BucketName:%s, ObjectNamePrefix:%s, FileName:%s, Md5sumValue:%s, Timestamp:%d>",
		this.Endpoint,
		this.BucketName,
		this.ObjectNamePrefix,
		this.FileName,
		this.Md5sumValue,
		this.Timestamp,
	)
}

type AliyunOssResponse struct {
	Message string `json:"message"`
	Invalid int    `json:"invalid"`
}

func (this *AliyunOssResponse) String() string {
	return fmt.Sprintf(
		"<Invalid:%v, Message:%s>",
		this.Invalid,
		this.Message,
	)
}
