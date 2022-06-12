package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

type HelloService struct{}

func (i *HelloService) Hello(request string, reply *string) error {
	*reply = fmt.Sprintf("[%s]Hello %s!", time.Now().Format(time.RFC3339), request)
	return nil
}

func main() {
	var err error
	err = rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	fmt.Println("rpc serve at: 127.0.0.1:1234")

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	rpc.ServeConn(conn)
}
