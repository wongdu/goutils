package main

import (
	"log"
	"math/rand"
	"os"
	pb "pb/helloworld"
	. "pb/pb"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

var (
	TransferClientsLock *sync.RWMutex                   = new(sync.RWMutex)
	TransferClients     map[string]*SingleConnRpcClient = map[string]*SingleConnRpcClient{}
	//Addrs    []string ={"127.0.0.1:7877","127.0.0.1:8433"}
	Addrs []string = []string{"127.0.0.1:7877"}
)

func main1() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func main() {
	//    var temp AliyunOssRequest{
	temp := AliyunOssRequest{
		Endpoint:         "",
		BucketName:       "",
		ObjectNamePrefix: "",
		FileName:         "",
		Md5sumValue:      "",
		Timestamp:        time.Now().Unix()}

	var resp AliyunOssResponse
	SendMetrics(temp, &resp)
	log.Println("<=", &resp)
}

func SendMetrics(metrics AliyunOssRequest, resp *AliyunOssResponse) {
	rand.Seed(time.Now().UnixNano())
	for _, i := range rand.Perm(len(Addrs)) {
		addr := Addrs[i]

		c := getTransferClient(addr)
		if c == nil {
			c = initTransferClient(addr)
		}

		if syncOssFile(c, metrics, resp) {
			break
		}
	}
}

func initTransferClient(addr string) *SingleConnRpcClient {
	var c *SingleConnRpcClient = &SingleConnRpcClient{
		RpcServer: addr,
		Timeout:   time.Duration(1000) * time.Millisecond,
	}
	TransferClientsLock.Lock()
	defer TransferClientsLock.Unlock()
	TransferClients[addr] = c

	return c
}

func syncOssFile(c *SingleConnRpcClient, metrics AliyunOssRequest, resp *AliyunOssResponse) bool {
	err := c.Call("AliyunOss.SyncOssFile", metrics, resp)
	if err != nil {
		log.Println("call AliyunOss.SyncOssFile fail:", c, err)
		return false
	}
	return true
}

func getTransferClient(addr string) *SingleConnRpcClient {
	TransferClientsLock.RLock()
	defer TransferClientsLock.RUnlock()

	if c, ok := TransferClients[addr]; ok {
		return c
	}
	return nil
}
