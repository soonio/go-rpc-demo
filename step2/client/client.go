package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc-hello/step2"
)

type HelloServiceClient struct {
	*rpc.Client
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(step2.HelloServiceName+".Hello", request, reply)
}

var _ step2.HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (client *HelloServiceClient, err error) {
	client = &HelloServiceClient{}

	// 使用gob
	//c, err := rpc.Dial(network, address)
	//if err == nil {
	//	client.Client = c
	//}

	// 使用json
	conn, err := net.Dial(network, address)
	if err == nil {
		client.Client = rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	}

	return
}

func main() {
	var err error
	client, err := DialHelloService("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	var reply string
	err = client.Hello("Hari", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
