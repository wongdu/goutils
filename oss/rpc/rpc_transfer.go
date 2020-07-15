package rpc

import (
	"fmt"
	"oss/common"
)

type AliyunOss int

func (this *AliyunOss) Ping(req common.NullRpcRequest, resp *common.SimpleRpcResponse) error {
	return nil
}

func (t *AliyunOss) SyncOssFile(req common.AliyunOssRequest, reply *common.AliyunOssResponse) error {
	fmt.Println(req)
	fmt.Println(&req)
	fmt.Println(reply)
	reply.Invalid = 100
	reply.Message = "test reply"
	return nil
}
