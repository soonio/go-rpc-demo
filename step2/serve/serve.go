package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	"rpc-hello/step2"
)

type HelloService struct{}

func (i *HelloService) Hello(request string, reply *string) error {
	*reply = fmt.Sprintf("[%s] Hello %s!", time.Now().Format("2006-01-02 15:04:05"), request)
	return nil
}

var _ step2.HelloServiceInterface = (*HelloService)(nil)

func main() {

	var err error
	err = rpc.RegisterName(step2.HelloServiceName, new(HelloService))
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
