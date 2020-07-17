package syncaliyunoss

import (
	"net"

	log "github.com/toolkits_/logrus"
	grpc "google.golang.org/grpc"
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
