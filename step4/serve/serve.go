package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"

	"net/rpc/jsonrpc"
	"rpc-hello/step4"
)

type HelloService struct{}

func (i *HelloService) Hello(request *step4.HelloRequest, reply *step4.HelloResponse) error {
	reply.Content = fmt.Sprintf("[%s] Hello %s!", time.Now().Format("2006-01-02 15:04:05"), request)
	return nil
}

var _ step4.HelloServiceInterface = (*HelloService)(nil)

func main() {

	var err error
	err = rpc.RegisterName(step4.HelloServiceName, new(HelloService))
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	fmt.Println("rpc serve at: 127.0.0.1:1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// 使用默认的gob
		//go rpc.ServeConn(conn)

		// 使用json
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
