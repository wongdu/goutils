package syncaliyunoss

import (
	"net"
	"time"

	log "github.com/toolkits_/logrus"
	grpc "google.golang.org/grpc"
)

// changelog:
// 1.0.0: code refactor
const (
	VERSION          = "1.0.0"
	COLLECT_INTERVAL = time.Second
)

func StartRpcServer(rpcPort string) {
	lis, err := net.Listen("tcp", ":"+rpcPort)
	if err != nil {

		log.Errorf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterSyncAliyunOssFileServer(s, &RpcServer{})
	if err := s.Serve(lis); err != nil {
		log.Errorf("Failed to serve: %v", err)
	}
}
