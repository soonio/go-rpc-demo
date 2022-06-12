package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc-hello/step2"
	"rpc-hello/step4"
)

type HelloServiceClient struct {
	*rpc.Client
}

func (p *HelloServiceClient) Hello(request *step4.HelloRequest, reply *step4.HelloResponse) error {
	return p.Client.Call(step2.HelloServiceName+".Hello", request, reply)
}

var _ step4.HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (client *HelloServiceClient, err error) {
	client = &HelloServiceClient{}

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

	var reply step4.HelloResponse
	err = client.Hello(&step4.HelloRequest{Content: "Hari"}, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply.String())
}
